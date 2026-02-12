package chain

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	nw "github.com/layer-3/nitewatch"
)

// HeadSubscriber abstracts the ability to subscribe to new block headers.
// *ethclient.Client satisfies this interface.
type HeadSubscriber interface {
	SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)
}

// Listener handles monitoring the blockchain for events from the ICustody contract.
type Listener struct {
	headSub       HeadSubscriber
	custody       *ICustody
	confirmations uint64
}

// NewListener creates a new Listener instance.
// headSub: a client supporting header subscriptions (e.g. *ethclient.Client via WebSocket)
// custody: bound ICustody contract instance
// confirmations: number of block confirmations required before an event is considered final
func NewListener(headSub HeadSubscriber, custody *ICustody, confirmations uint64) *Listener {
	return &Listener{
		headSub:       headSub,
		custody:       custody,
		confirmations: confirmations,
	}
}

// Compile-time check that Listener implements nw.EventListener.
var _ nw.EventListener = (*Listener)(nil)

// WatchDeposited subscribes to Deposited events and sends confirmed domain events to the sink channel.
// The sink channel will be closed when the context is cancelled or an error occurs.
func (l *Listener) WatchDeposited(ctx context.Context, sink chan<- *nw.DepositEvent) error {
	raw := make(chan *ICustodyDeposited)

	errCh := make(chan error, 1)
	go func() {
		errCh <- watchWithConfirmations(ctx, raw, l.confirmations,
			func(rawSink chan<- *ICustodyDeposited) (ethereum.Subscription, error) {
				return l.custody.WatchDeposited(&bind.WatchOpts{Context: ctx}, rawSink, nil, nil)
			},
			func(ch chan<- *types.Header) (ethereum.Subscription, error) {
				return l.headSub.SubscribeNewHead(ctx, ch)
			},
			func(e *ICustodyDeposited) types.Log { return e.Raw },
		)
	}()

	defer close(sink)
	for ev := range raw {
		select {
		case sink <- &nw.DepositEvent{
			User:        ev.User,
			Token:       ev.Token,
			Amount:      ev.Amount,
			BlockNumber: ev.Raw.BlockNumber,
			TxHash:      ev.Raw.TxHash,
		}:
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return <-errCh
}

// WatchWithdrawStarted subscribes to WithdrawStarted events and sends confirmed domain events to the sink channel.
func (l *Listener) WatchWithdrawStarted(ctx context.Context, sink chan<- *nw.WithdrawStartedEvent) error {
	raw := make(chan *ICustodyWithdrawStarted)

	errCh := make(chan error, 1)
	go func() {
		errCh <- watchWithConfirmations(ctx, raw, l.confirmations,
			func(rawSink chan<- *ICustodyWithdrawStarted) (ethereum.Subscription, error) {
				return l.custody.WatchWithdrawStarted(&bind.WatchOpts{Context: ctx}, rawSink, nil, nil, nil)
			},
			func(ch chan<- *types.Header) (ethereum.Subscription, error) {
				return l.headSub.SubscribeNewHead(ctx, ch)
			},
			func(e *ICustodyWithdrawStarted) types.Log { return e.Raw },
		)
	}()

	defer close(sink)
	for ev := range raw {
		select {
		case sink <- &nw.WithdrawStartedEvent{
			WithdrawalID: ev.WithdrawalId,
			User:         ev.User,
			Token:        ev.Token,
			Amount:       ev.Amount,
			Nonce:        ev.Nonce,
			BlockNumber:  ev.Raw.BlockNumber,
			TxHash:       ev.Raw.TxHash,
		}:
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return <-errCh
}

// WatchWithdrawFinalized subscribes to WithdrawFinalized events and sends confirmed domain events to the sink channel.
func (l *Listener) WatchWithdrawFinalized(ctx context.Context, sink chan<- *nw.WithdrawFinalizedEvent) error {
	raw := make(chan *ICustodyWithdrawFinalized)

	errCh := make(chan error, 1)
	go func() {
		errCh <- watchWithConfirmations(ctx, raw, l.confirmations,
			func(rawSink chan<- *ICustodyWithdrawFinalized) (ethereum.Subscription, error) {
				return l.custody.WatchWithdrawFinalized(&bind.WatchOpts{Context: ctx}, rawSink, nil)
			},
			func(ch chan<- *types.Header) (ethereum.Subscription, error) {
				return l.headSub.SubscribeNewHead(ctx, ch)
			},
			func(e *ICustodyWithdrawFinalized) types.Log { return e.Raw },
		)
	}()

	defer close(sink)
	for ev := range raw {
		select {
		case sink <- &nw.WithdrawFinalizedEvent{
			WithdrawalID: ev.WithdrawalId,
			Success:      ev.Success,
			BlockNumber:  ev.Raw.BlockNumber,
			TxHash:       ev.Raw.TxHash,
		}:
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return <-errCh
}

// watchWithConfirmations is the generic confirmation-tracking event watcher.
// It buffers events until they reach the required block depth before emitting them to sink.
// If confirmations is 0, events are emitted immediately.
func watchWithConfirmations[E any](
	ctx context.Context,
	sink chan<- E,
	confirmations uint64,
	subscribe func(chan<- E) (ethereum.Subscription, error),
	subscribeHead func(chan<- *types.Header) (ethereum.Subscription, error),
	getRaw func(E) types.Log,
) error {
	defer close(sink)

	rawSink := make(chan E)
	sub, err := subscribe(rawSink)
	if err != nil {
		return fmt.Errorf("failed to subscribe to events: %w", err)
	}
	defer sub.Unsubscribe()

	headers := make(chan *types.Header)
	headSub, err := subscribeHead(headers)
	if err != nil {
		return fmt.Errorf("failed to subscribe to new heads: %w", err)
	}
	defer headSub.Unsubscribe()

	type pendingEvent struct {
		event       E
		blockNumber uint64
	}
	var pending []pendingEvent

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-sub.Err():
			return fmt.Errorf("event subscription error: %w", err)
		case err := <-headSub.Err():
			return fmt.Errorf("header subscription error: %w", err)
		case ev := <-rawSink:
			raw := getRaw(ev)
			if raw.Removed {
				n := 0
				for _, p := range pending {
					pRaw := getRaw(p.event)
					if pRaw.TxHash != raw.TxHash || pRaw.Index != raw.Index {
						pending[n] = p
						n++
					}
				}
				pending = pending[:n]
			} else if confirmations == 0 {
				select {
				case sink <- ev:
				case <-ctx.Done():
					return ctx.Err()
				}
			} else {
				pending = append(pending, pendingEvent{
					event:       ev,
					blockNumber: raw.BlockNumber,
				})
			}
		case head := <-headers:
			if confirmations == 0 {
				continue
			}
			currentBlock := head.Number.Uint64()
			n := 0
			for _, p := range pending {
				if currentBlock+1 >= p.blockNumber+confirmations {
					select {
					case sink <- p.event:
					case <-ctx.Done():
						return ctx.Err()
					}
				} else {
					pending[n] = p
					n++
				}
			}
			pending = pending[:n]
		}
	}
}
