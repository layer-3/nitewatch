// Copied from github.com/layer-3/pathfinder/pkg/ethlistener (commit hash 686dc94b80985eba798fdec499b9a802dbf80471).
// Adapted: replaced slog with ipfs/go-log, removed FetchHistoricalLogs,
// which depends on pathfinder's ethclient package.
package custody

import (
	"context"
	"errors"
	"fmt"
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

var ethLogger = logging.Logger("ethlistener")

const (
	maxBackOffCount = 5
)

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
	var historicalCh, currentCh chan types.Log
	var eventSubscription event.Subscription

	ethLogger.Debugw("starting listening events", "subID", subID, "contractAddress", contractAddress.String())
	for {
		if err := ctx.Err(); err != nil {
			ethLogger.Infow("context cancelled, stopping listener", "subID", subID)
			if eventSubscription != nil {
				eventSubscription.Unsubscribe()
			}
			return
		}

		if eventSubscription == nil {
			if !waitForBackOffTimeout(ctx, int(backOffCount.Load()), "event subscription") {
				return
			}

			historicalCh = make(chan types.Log, 1)
			currentCh = make(chan types.Log, 100)

			if lastBlock == 0 {
				ethLogger.Infow("skipping historical logs fetching", "subID", subID, "contractAddress", contractAddress.String())
			} else {
				var header *types.Header
				var err error
				headerCtx, cancel := context.WithTimeout(ctx, 1*time.Minute)
				err = debounce.Debounce(headerCtx, ethLogger, func(ctx context.Context) error {
					header, err = client.HeaderByNumber(ctx, nil)
					return err
				})
				cancel()
				if err != nil {
					if ctx.Err() != nil {
						return
					}
					ethLogger.Errorw("failed to get latest block", "error", err, "subID", subID, "contractAddress", contractAddress.String())
					backOffCount.Add(1)
					continue
				}

				go reconcileBlockRange(
					ctx,
					client,
					subID,
					contractAddress,
					networkID,
					header.Number.Uint64(),
					lastBlock,
					lastIndex,
					topics,
					historicalCh,
				)
			}

			watchFQ := ethereum.FilterQuery{
				Addresses: []common.Address{contractAddress},
			}
			eventSub, err := client.SubscribeFilterLogs(ctx, watchFQ, currentCh)
			if err != nil {
				if ctx.Err() != nil {
					return
				}
				ethLogger.Errorw("failed to subscribe on events", "error", err, "subID", subID, "contractAddress", contractAddress.String())
				backOffCount.Add(1)
				continue
			}

			eventSubscription = eventSub
			ethLogger.Infow("watching events", "subID", subID, "contractAddress", contractAddress.String())
			backOffCount.Store(0)
		}

		select {
		case <-ctx.Done():
			ethLogger.Infow("context cancelled, stopping listener", "subID", subID)
			eventSubscription.Unsubscribe()
			return
		case eventLog := <-historicalCh:
			ethLogger.Debugw("received historical event", "subID", subID, "blockNumber", eventLog.BlockNumber, "logIndex", eventLog.Index)
			handler(eventLog)
		case eventLog := <-currentCh:
			lastBlock = eventLog.BlockNumber
			ethLogger.Debugw("received new event", "subID", subID, "blockNumber", lastBlock, "logIndex", eventLog.Index)
			handler(eventLog)
		case err := <-eventSubscription.Err():
			if err != nil {
				ethLogger.Errorw("event subscription error", "error", err, "subID", subID, "contractAddress", contractAddress.String())
				eventSubscription.Unsubscribe()
			} else {
				ethLogger.Debugw("subscription closed, resubscribing", "subID", subID, "contractAddress", contractAddress.String())
			}

			eventSubscription = nil
		}
	}
}

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
	historicalCh chan types.Log,
) {
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
		err = debounce.Debounce(logsCtx, ethLogger, func(ctx context.Context) error {
			logs, err = client.FilterLogs(ctx, fetchFQ)
			return err
		})
		cancel()
		if err != nil {
			if strings.Contains(err.Error(), "Exceeded max range limit for eth_getLogs:") {
				newEndBlock := endBlock - (endBlock-startBlock)/2
				ethLogger.Infow("eth_getLogs exceeded max range limit, reducing block range", "subID", subID, "startBlock", startBlock, "oldEndBlock", endBlock, "newEndBlock", newEndBlock)
				endBlock = newEndBlock
				continue
			}

			newStartBlock, newEndBlock, extractErr := extractAdvisedBlockRange(err.Error())
			if extractErr != nil {
				ethLogger.Errorw("failed to filter logs", "error", err, "extractErr", extractErr, "subID", subID, "startBlock", startBlock, "endBlock", endBlock)
				backOffCount.Add(1)
				continue
			}
			startBlock, endBlock = newStartBlock, newEndBlock
			ethLogger.Infow("retrying with advised block range", "subID", subID, "startBlock", startBlock, "endBlock", endBlock)
			continue
		}
		ethLogger.Infow("fetched historical logs", "subID", subID, "count", len(logs), "startBlock", startBlock, "endBlock", endBlock)

		for _, ethLog := range logs {
			if ethLog.BlockNumber == lastBlock && ethLog.Index <= uint(lastIndex) {
				ethLogger.Infow("skipping previously known event", "subID", subID, "blockNumber", ethLog.BlockNumber, "logIndex", ethLog.Index)
				continue
			}

			historicalCh <- ethLog
		}

		startBlock = endBlock + 1
		endBlock += blockStep
	}
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
		ethLogger.Errorw("back off limit reached, exiting", "originator", originator, "backOffCount", backOffCount)
		return true
	}

	if backOffCount > 0 {
		ethLogger.Infow("backing off", "originator", originator, "backOffCount", backOffCount)
		select {
		case <-time.After(time.Duration(2^backOffCount-1) * time.Second):
		case <-ctx.Done():
			return false
		}
	}
	return true
}
