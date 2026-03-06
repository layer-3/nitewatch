package custody

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	logging "github.com/ipfs/go-log/v2"
	"github.com/layer-3/clearsync/pkg/debounce"
)

var listenerLogger = logging.Logger("custody-listener")

const (
	maxBackOffCount = 5
)

// Listener handles monitoring the blockchain for events from the custody contract.
type Listener struct {
	client           bind.ContractBackend
	contractAddr     common.Address
	withdrawFilterer *IWithdrawFilterer
	depositFilterer  *IDepositFilterer
}

// NewListener creates a new Listener instance.
// client: an Ethereum client supporting log subscriptions (e.g. *ethclient.Client via WebSocket)
// contractAddr: address of the custody contract
// withdraw: bound IWithdraw contract instance
// deposit: bound IDeposit contract instance (can be nil if deposit events are not needed)
func NewListener(client bind.ContractBackend, contractAddr common.Address, withdraw *IWithdraw, deposit *IDeposit) *Listener {
	l := &Listener{
		client:       client,
		contractAddr: contractAddr,
	}
	if withdraw != nil {
		l.withdrawFilterer = &withdraw.IWithdrawFilterer
	}
	if deposit != nil {
		l.depositFilterer = &deposit.IDepositFilterer
	}
	return l
}

// Compile-time check that Listener implements EventListener.
var _ EventListener = (*Listener)(nil)

// WatchWithdrawStarted subscribes to WithdrawStarted events and sends them to the sink channel.
// This function blocks forever; run it in a goroutine. The sink channel is closed when it returns.
func (l *Listener) WatchWithdrawStarted(ctx context.Context, sink chan<- *WithdrawStartedEvent, fromBlock uint64, fromLogIndex uint32) {
	defer close(sink)

	parsedABI, err := IWithdrawMetaData.GetAbi()
	if err != nil {
		return
	}
	topic := parsedABI.Events["WithdrawStarted"].ID

	listenEvents(ctx, l.client, "withdraw-started", l.contractAddr, 0, fromBlock, fromLogIndex,
		[][]common.Hash{{topic}},
		func(log types.Log) {

			ev, err := l.withdrawFilterer.ParseWithdrawStarted(log)
			if err != nil {
				return
			}
			sink <- &WithdrawStartedEvent{
				WithdrawalID: ev.WithdrawalId,
				User:         ev.User,
				Token:        ev.Token,
				Amount:       ev.Amount,
				Nonce:        ev.Nonce,
				BlockNumber:  ev.Raw.BlockNumber,
				TxHash:       ev.Raw.TxHash,
				LogIndex:     ev.Raw.Index,
			}
		},
	)
}

// WatchWithdrawFinalized subscribes to WithdrawFinalized events and sends them to the sink channel.
// This function blocks forever; run it in a goroutine. The sink channel is closed when it returns.
func (l *Listener) WatchWithdrawFinalized(ctx context.Context, sink chan<- *WithdrawFinalizedEvent, fromBlock uint64, fromLogIndex uint32) {
	defer close(sink)

	parsedABI, err := IWithdrawMetaData.GetAbi()
	if err != nil {
		return
	}
	topic := parsedABI.Events["WithdrawFinalized"].ID

	listenEvents(ctx, l.client, "withdraw-finalized", l.contractAddr, 0, fromBlock, fromLogIndex,
		[][]common.Hash{{topic}},
		func(log types.Log) {

			ev, err := l.withdrawFilterer.ParseWithdrawFinalized(log)
			if err != nil {
				return
			}
			sink <- &WithdrawFinalizedEvent{
				WithdrawalID: ev.WithdrawalId,
				Success:      ev.Success,
				BlockNumber:  ev.Raw.BlockNumber,
				TxHash:       ev.Raw.TxHash,
				LogIndex:     ev.Raw.Index,
			}
		},
	)
}

// WatchDeposited subscribes to Deposited events and sends them to the sink channel.
// This function blocks forever; run it in a goroutine. The sink channel is closed when it returns.
func (l *Listener) WatchDeposited(ctx context.Context, sink chan<- *DepositedEvent, fromBlock uint64, fromLogIndex uint32) {
	defer close(sink)

	if l.depositFilterer == nil {
		return
	}

	parsedABI, err := IDepositMetaData.GetAbi()
	if err != nil {
		return
	}
	topic := parsedABI.Events["Deposited"].ID

	listenEvents(ctx, l.client, "deposited", l.contractAddr, 0, fromBlock, fromLogIndex,
		[][]common.Hash{{topic}},
		func(log types.Log) {

			ev, err := l.depositFilterer.ParseDeposited(log)
			if err != nil {
				return
			}
			sink <- &DepositedEvent{
				User:        ev.User,
				Token:       ev.Token,
				Amount:      ev.Amount,
				BlockNumber: ev.Raw.BlockNumber,
				TxHash:      ev.Raw.TxHash,
				LogIndex:    ev.Raw.Index,
			}
		},
	)
}

// listenEvents subscribes to on-chain logs matching the given topics and feeds them to handler.
// It handles reconnection with backoff, historical log reconciliation, and live subscription.
type logHandler func(log types.Log)

func listenEvents(
	ctx context.Context,
	client bind.ContractBackend,
	subID string,
	contractAddress common.Address,
	networkID uint32,
	lastBlock uint64,
	lastIndex uint32,
	topics [][]common.Hash,
	handler logHandler,
) {
	var backOffCount atomic.Uint64
	var eventSubscription event.Subscription
	var headCh chan *types.Header

	// Interface for SubscribeNewHead, as it's not in bind.ContractBackend
	type headerSubscriber interface {
		SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)
	}

	listenerLogger.Debugw("starting listening events", "subID", subID, "contractAddress", contractAddress.String())
	for {
		if err := ctx.Err(); err != nil {
			listenerLogger.Infow("context cancelled, stopping listener", "subID", subID)
			if eventSubscription != nil {
				eventSubscription.Unsubscribe()
			}
			return
		}

		if eventSubscription == nil {
			if !waitForBackOffTimeout(ctx, int(backOffCount.Load()), "event subscription") {
				return
			}

			// Channel for new block headers
			headCh = make(chan *types.Header, 100)

			subscriber, ok := client.(headerSubscriber)
			if !ok {
				listenerLogger.Errorw("client does not support SubscribeNewHead", "subID", subID)
				return
			}

			sub, subscribeErr := subscriber.SubscribeNewHead(ctx, headCh)
			if subscribeErr != nil {
				if ctx.Err() != nil {
					return
				}
				// Return instead of retrying: the underlying connection is
				// likely dead, so resubscribing on the same client will keep
				// failing. Returning causes the Service-level loop to dial a
				// fresh ethclient and restart the listener.
				listenerLogger.Errorw("failed to subscribe to new heads, exiting listener", "error", subscribeErr, "subID", subID)
				return
			}
			eventSubscription = sub
			listenerLogger.Infow("subscribed to new heads", "subID", subID)
			backOffCount.Store(0)

			// Initial sync up to current head
			var header *types.Header
			var err error
			headerCtx, cancel := context.WithTimeout(ctx, 1*time.Minute)
			err = debounce.Debounce(headerCtx, listenerLogger, func(ctx context.Context) error {
				header, err = client.HeaderByNumber(ctx, nil)
				return err
			})
			cancel()
			if err != nil {
				listenerLogger.Errorw("failed to get latest block header", "error", err, "subID", subID)
				eventSubscription.Unsubscribe()
				eventSubscription = nil
				backOffCount.Add(1)
				continue
			}

			currentHead := header.Number.Uint64()
			targetBlock := currentHead

			if targetBlock > lastBlock {
				lastBlock, lastIndex = reconcileBlockRange(
					ctx,
					client,
					subID,
					contractAddress,
					networkID,
					targetBlock,
					lastBlock,
					lastIndex,
					topics,
					handler,
				)
			}
		}

		select {
		case <-ctx.Done():
			listenerLogger.Infow("context cancelled, stopping listener", "subID", subID)
			eventSubscription.Unsubscribe()
			return
		case header := <-headCh:
			currentHead := header.Number.Uint64()
			targetBlock := currentHead

			if targetBlock > lastBlock {
				lastBlock, lastIndex = reconcileBlockRange(
					ctx,
					client,
					subID,
					contractAddress,
					networkID,
					targetBlock,
					lastBlock,
					lastIndex,
					topics,
					handler,
				)
			}
		case err := <-eventSubscription.Err():
			if err != nil {
				listenerLogger.Errorw("event subscription error", "error", err, "subID", subID, "contractAddress", contractAddress.String())
				eventSubscription.Unsubscribe()
			} else {
				listenerLogger.Debugw("subscription closed, resubscribing", "subID", subID, "contractAddress", contractAddress.String())
			}

			eventSubscription = nil
		}
	}
}

// reconcileBlockRange fetches historical logs from lastBlock up to
// currentBlock and feeds them to handler.  It returns the cursor
// (block, logIndex) representing the most recent log that was fully
// processed.  When an entire batch is consumed without any matching
// logs, the cursor advances to (endBlock, math.MaxUint32) so that the
// caller never re-queries logs from blocks that were already scanned.
//
// The returned logIndex is math.MaxUint32 when the indicated block was
// fully consumed (i.e. either it contained no matching events or all
// its events were dispatched).  Callers should propagate these values
// into lastBlock / lastIndex so the skip-guard at the top of the
// processing loop works correctly on the next invocation.
func reconcileBlockRange(
	ctx context.Context,
	client bind.ContractBackend,
	subID string,
	contractAddress common.Address,
	networkID uint32,
	currentBlock uint64,
	lastBlock uint64,
	lastIndex uint32,
	topics [][]common.Hash,
	handler logHandler,
) (finishedBlock uint64, finishedIndex uint32) {
	finishedBlock = lastBlock
	finishedIndex = lastIndex

	var backOffCount atomic.Uint64
	const blockStep = 10000
	startBlock := lastBlock
	endBlock := startBlock + blockStep

	for currentBlock > startBlock {
		if ctx.Err() != nil {
			return
		}
		if !waitForBackOffTimeout(ctx, int(backOffCount.Load()), "reconcile block range") {
			return
		}

		if endBlock > currentBlock {
			endBlock = currentBlock
		}

		fetchFQ := ethereum.FilterQuery{
			Addresses: []common.Address{contractAddress},
			FromBlock: new(big.Int).SetUint64(startBlock),
			ToBlock:   new(big.Int).SetUint64(endBlock),
			Topics:    topics,
		}

		var logs []types.Log
		var err error
		logsCtx, cancel := context.WithTimeout(ctx, 1*time.Minute)
		err = debounce.Debounce(logsCtx, listenerLogger, func(ctx context.Context) error {
			logs, err = client.FilterLogs(ctx, fetchFQ)
			return err
		})
		cancel()
		if err != nil {
			if strings.Contains(err.Error(), "Exceeded max range limit for eth_getLogs:") {
				newEndBlock := endBlock - (endBlock-startBlock)/2
				listenerLogger.Infow("eth_getLogs exceeded max range limit, reducing block range", "subID", subID, "startBlock", startBlock, "oldEndBlock", endBlock, "newEndBlock", newEndBlock)
				endBlock = newEndBlock
				continue
			}

			newStartBlock, newEndBlock, extractErr := extractAdvisedBlockRange(err.Error())
			if extractErr != nil {
				listenerLogger.Errorw("failed to filter logs", "error", err, "extractErr", extractErr, "subID", subID, "startBlock", startBlock, "endBlock", endBlock)
				backOffCount.Add(1)
				continue
			}
			startBlock, endBlock = newStartBlock, newEndBlock
			listenerLogger.Infow("retrying with advised block range", "subID", subID, "startBlock", startBlock, "endBlock", endBlock)
			continue
		}
		listenerLogger.Infow("fetched historical logs", "subID", subID, "count", len(logs), "startBlock", startBlock, "endBlock", endBlock)

		for _, ethLog := range logs {
			if ethLog.BlockNumber == lastBlock && ethLog.Index <= uint(lastIndex) {
				listenerLogger.Infow("skipping previously known event", "subID", subID, "blockNumber", ethLog.BlockNumber, "logIndex", ethLog.Index)
				continue
			}

			handler(ethLog)
			finishedBlock = ethLog.BlockNumber
			finishedIndex = uint32(ethLog.Index)
		}

		// Even if no matching logs were found in this batch, we have
		// fully scanned up to endBlock — advance the cursor so the
		// next call does not re-query the same range.
		if finishedBlock < endBlock {
			finishedBlock = endBlock
			finishedIndex = math.MaxUint32
		}

		startBlock = endBlock + 1
		endBlock += blockStep
	}
	return
}

func extractAdvisedBlockRange(msg string) (startBlock, endBlock uint64, err error) {
	if !strings.Contains(msg, "query returned more than 10000 results") {
		err = errors.New("error message doesn't contain advised block range")
		return
	}

	re := regexp.MustCompile(`\[0x([0-9a-fA-F]+), 0x([0-9a-fA-F]+)\]`)
	match := re.FindStringSubmatch(msg)
	if len(match) != 3 {
		err = errors.New("failed to extract block range from error message")
		return
	}

	startBlock, err = strconv.ParseUint(match[1], 16, 64)
	if err != nil {
		err = fmt.Errorf("failed to parse block range from error message: %w", err)
		return
	}
	endBlock, err = strconv.ParseUint(match[2], 16, 64)
	if err != nil {
		err = fmt.Errorf("failed to parse block range from error message: %w", err)
		return
	}
	return
}

func waitForBackOffTimeout(ctx context.Context, backOffCount int, originator string) bool {
	if backOffCount > maxBackOffCount {
		listenerLogger.Errorw("back off limit reached, exiting", "originator", originator, "backOffCount", backOffCount)
		return false
	}

	if backOffCount > 0 {
		listenerLogger.Infow("backing off", "originator", originator, "backOffCount", backOffCount)
		select {
		case <-time.After(time.Duration(1<<backOffCount) * time.Second):
		case <-ctx.Done():
			return false
		}
	}
	return true
}
