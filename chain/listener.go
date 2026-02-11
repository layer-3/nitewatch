package chain

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Listener handles monitoring the blockchain for events from the ICustody contract.
type Listener struct {
	client        *ethclient.Client
	custody       *ICustody
	confirmations uint64
}

// NewListener creates a new Listener instance.
// client: connected ethclient (must support subscriptions, e.g. via WebSocket)
// contractAddress: address of the ICustody contract
// confirmations: number of block confirmations required before an event is considered final (e.g. 12)
func NewListener(client *ethclient.Client, contractAddress common.Address, confirmations uint64) (*Listener, error) {
	custody, err := NewICustody(contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind contract: %w", err)
	}

	return &Listener{
		client:        client,
		custody:       custody,
		confirmations: confirmations,
	}, nil
}

// WatchDeposited subscribes to Deposited events and sends confirmed events to the sink channel.
// The sink channel will be closed when the context is cancelled or an error occurs.
func (l *Listener) WatchDeposited(ctx context.Context, sink chan<- *ICustodyDeposited) error {
	defer close(sink)

	// Channel for raw events from the subscription
	rawSink := make(chan *ICustodyDeposited)

	// Subscribe to the contract event
	sub, err := l.custody.WatchDeposited(&bind.WatchOpts{Context: ctx}, rawSink, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to watch Deposited events: %w", err)
	}
	defer sub.Unsubscribe()

	// Subscribe to new block headers to track confirmations
	headers := make(chan *types.Header)
	headSub, err := l.client.SubscribeNewHead(ctx, headers)
	if err != nil {
		return fmt.Errorf("failed to subscribe to new heads: %w", err)
	}
	defer headSub.Unsubscribe()

	// Pending events waiting for confirmation
	type pendingEvent struct {
		event       *ICustodyDeposited
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
			if ev.Raw.Removed {
				// Handle reorg: remove the event from pending if it exists
				n := 0
				for _, p := range pending {
					if p.event.Raw.TxHash != ev.Raw.TxHash || p.event.Raw.Index != ev.Raw.Index {
						pending[n] = p
						n++
					}
				}
				pending = pending[:n]
			} else {
				// New event, add to pending
				// If confirmations is 0, send immediately
				if l.confirmations == 0 {
					select {
					case sink <- ev:
					case <-ctx.Done():
						return ctx.Err()
					}
				} else {
					pending = append(pending, pendingEvent{
						event:       ev,
						blockNumber: ev.Raw.BlockNumber,
					})
				}
			}
		case head := <-headers:
			// If confirmations == 0, we handled it above.
			if l.confirmations == 0 {
				continue
			}

			currentBlock := head.Number.Uint64()
			// Check confirmations
			n := 0
			for _, p := range pending {
				// Check depth. If confirmations=1, we need currentBlock >= blockNumber.
				// If confirmations=12, we need currentBlock >= blockNumber + 11.
				// Formula: depth = currentBlock - blockNumber + 1
				// depth >= confirmations  =>  currentBlock - blockNumber + 1 >= confirmations
				// => currentBlock >= blockNumber + confirmations - 1
				if currentBlock+1 >= p.blockNumber+l.confirmations {
					// Confirmed, send to sink
					select {
					case sink <- p.event:
					case <-ctx.Done():
						return ctx.Err()
					}
				} else {
					// Not yet confirmed, keep it
					pending[n] = p
					n++
				}
			}
			pending = pending[:n]
		}
	}
}

// WatchWithdrawStarted subscribes to WithdrawStarted events and sends confirmed events to the sink channel.
func (l *Listener) WatchWithdrawStarted(ctx context.Context, sink chan<- *ICustodyWithdrawStarted) error {
	defer close(sink)

	rawSink := make(chan *ICustodyWithdrawStarted)

	sub, err := l.custody.WatchWithdrawStarted(&bind.WatchOpts{Context: ctx}, rawSink, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to watch WithdrawStarted events: %w", err)
	}
	defer sub.Unsubscribe()

	headers := make(chan *types.Header)
	headSub, err := l.client.SubscribeNewHead(ctx, headers)
	if err != nil {
		return fmt.Errorf("failed to subscribe to new heads: %w", err)
	}
	defer headSub.Unsubscribe()

	type pendingEvent struct {
		event       *ICustodyWithdrawStarted
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
			if ev.Raw.Removed {
				n := 0
				for _, p := range pending {
					if p.event.Raw.TxHash != ev.Raw.TxHash || p.event.Raw.Index != ev.Raw.Index {
						pending[n] = p
						n++
					}
				}
				pending = pending[:n]
			} else {
				if l.confirmations == 0 {
					select {
					case sink <- ev:
					case <-ctx.Done():
						return ctx.Err()
					}
				} else {
					pending = append(pending, pendingEvent{
						event:       ev,
						blockNumber: ev.Raw.BlockNumber,
					})
				}
			}
		case head := <-headers:
			if l.confirmations == 0 {
				continue
			}
			currentBlock := head.Number.Uint64()
			n := 0
			for _, p := range pending {
				if currentBlock+1 >= p.blockNumber+l.confirmations {
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

// WatchWithdrawFinalized subscribes to WithdrawFinalized events and sends confirmed events to the sink channel.
func (l *Listener) WatchWithdrawFinalized(ctx context.Context, sink chan<- *ICustodyWithdrawFinalized) error {
	defer close(sink)

	rawSink := make(chan *ICustodyWithdrawFinalized)

	sub, err := l.custody.WatchWithdrawFinalized(&bind.WatchOpts{Context: ctx}, rawSink, nil)
	if err != nil {
		return fmt.Errorf("failed to watch WithdrawFinalized events: %w", err)
	}
	defer sub.Unsubscribe()

	headers := make(chan *types.Header)
	headSub, err := l.client.SubscribeNewHead(ctx, headers)
	if err != nil {
		return fmt.Errorf("failed to subscribe to new heads: %w", err)
	}
	defer headSub.Unsubscribe()

	type pendingEvent struct {
		event       *ICustodyWithdrawFinalized
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
			if ev.Raw.Removed {
				n := 0
				for _, p := range pending {
					if p.event.Raw.TxHash != ev.Raw.TxHash || p.event.Raw.Index != ev.Raw.Index {
						pending[n] = p
						n++
					}
				}
				pending = pending[:n]
			} else {
				if l.confirmations == 0 {
					select {
					case sink <- ev:
					case <-ctx.Done():
						return ctx.Err()
					}
				} else {
					pending = append(pending, pendingEvent{
						event:       ev,
						blockNumber: ev.Raw.BlockNumber,
					})
				}
			}
		case head := <-headers:
			if l.confirmations == 0 {
				continue
			}
			currentBlock := head.Number.Uint64()
			n := 0
			for _, p := range pending {
				if currentBlock+1 >= p.blockNumber+l.confirmations {
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
