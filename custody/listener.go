package custody

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
