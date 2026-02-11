package nitewatch

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// LimitChecker defines the business logic for enforcing withdrawal limits.
type LimitChecker interface {
	// Check verifies if a withdrawal amount is within limits for the given token.
	Check(token common.Address, amount *big.Int) error

	// Record persists the withdrawal event.
	Record(w *Withdrawal) error
}

// Custody defines the write operations for the ICustody smart contract.
type Custody interface {
	Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error)
	StartWithdraw(opts *bind.TransactOpts, user common.Address, token common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error)
	FinalizeWithdraw(opts *bind.TransactOpts, withdrawalId [32]byte) (*types.Transaction, error)
}

// Withdrawal represents a completed withdrawal event.
type Withdrawal struct {
	WithdrawalID [32]byte
	User         common.Address
	Token        common.Address
	Amount       *big.Int
	BlockNumber  uint64
	TxHash       common.Hash
	Timestamp    time.Time
}

// WithdrawalStore defines the storage operations for tracking withdrawals.
type WithdrawalStore interface {
	Save(w *Withdrawal) error
	GetTotalWithdrawn(token common.Address, since time.Time) (*big.Int, error)
}
