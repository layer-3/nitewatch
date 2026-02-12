package nitewatch

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// --- Domain Events ---

// DepositEvent represents a confirmed Deposited event from the custody contract.
type DepositEvent struct {
	User        common.Address
	Token       common.Address
	Amount      *big.Int
	BlockNumber uint64
	TxHash      common.Hash
}

// WithdrawStartedEvent represents a confirmed WithdrawStarted event from the custody contract.
type WithdrawStartedEvent struct {
	WithdrawalID [32]byte
	User         common.Address
	Token        common.Address
	Amount       *big.Int
	Nonce        *big.Int
	BlockNumber  uint64
	TxHash       common.Hash
}

// WithdrawFinalizedEvent represents a confirmed WithdrawFinalized event from the custody contract.
type WithdrawFinalizedEvent struct {
	WithdrawalID [32]byte
	Success      bool
	BlockNumber  uint64
	TxHash       common.Hash
}

// --- Domain Model ---

// Withdrawal represents a recorded withdrawal for limit tracking.
type Withdrawal struct {
	WithdrawalID [32]byte
	User         common.Address
	Token        common.Address
	Amount       *big.Int
	BlockNumber  uint64
	TxHash       common.Hash
	Timestamp    time.Time
}

// --- Interfaces ---

// Custody defines the write operations for the ICustody smart contract.
// Cage uses StartWithdraw; Nitewatch uses FinalizeWithdraw.
type Custody interface {
	Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error)
	StartWithdraw(opts *bind.TransactOpts, user common.Address, token common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error)
	FinalizeWithdraw(opts *bind.TransactOpts, withdrawalId [32]byte) (*types.Transaction, error)
	RejectWithdraw(opts *bind.TransactOpts, withdrawalId [32]byte) (*types.Transaction, error)
}

// EventListener defines the ability to subscribe to ICustody contract events.
// The sink channel is closed when the context is cancelled or an error occurs.
type EventListener interface {
	WatchDeposited(ctx context.Context, sink chan<- *DepositEvent) error
	WatchWithdrawStarted(ctx context.Context, sink chan<- *WithdrawStartedEvent) error
	WatchWithdrawFinalized(ctx context.Context, sink chan<- *WithdrawFinalizedEvent) error
}

// WithdrawalStore defines the storage operations for tracking withdrawals.
type WithdrawalStore interface {
	Save(w *Withdrawal) error
	GetTotalWithdrawn(token common.Address, since time.Time) (*big.Int, error)
}
