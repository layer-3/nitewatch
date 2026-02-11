package store

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	nw "github.com/layer-3/nitewatch"
)

// WithdrawalModel is the GORM model for persisting withdrawal events.
type WithdrawalModel struct {
	gorm.Model
	WithdrawalID string `gorm:"uniqueIndex;type:varchar(66)"` // 0x + 64 hex chars
	User         string `gorm:"index;type:varchar(42)"`       // 0x + 40 hex chars
	Token        string `gorm:"index;type:varchar(42)"`       // 0x + 40 hex chars
	Amount       string `gorm:"type:text"`                    // big.Int as string
	BlockNumber  uint64
	TxHash       string    `gorm:"type:varchar(66)"`
	Timestamp    time.Time `gorm:"index"`
}

// Adapter implements the WithdrawalStore interface using GORM and SQLite.
type Adapter struct {
	db *gorm.DB
}

// NewAdapter initializes a new GORM adapter with SQLite.
func NewAdapter(dbPath string) (*Adapter, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&WithdrawalModel{}); err != nil {
		return nil, err
	}

	return &Adapter{db: db}, nil
}

// Save persists a withdrawal event to the database.
func (a *Adapter) Save(w *nw.Withdrawal) error {
	model := &WithdrawalModel{
		WithdrawalID: common.Hash(w.WithdrawalID).Hex(),
		User:         w.User.Hex(),
		Token:        w.Token.Hex(),
		Amount:       w.Amount.String(),
		BlockNumber:  w.BlockNumber,
		TxHash:       w.TxHash.Hex(),
		Timestamp:    w.Timestamp,
	}
	return a.db.Create(model).Error
}

// GetTotalWithdrawn calculates the total amount withdrawn for a token since a given time.
func (a *Adapter) GetTotalWithdrawn(token common.Address, since time.Time) (*big.Int, error) {
	var withdrawals []WithdrawalModel
	
	// Query all withdrawals for the token since the timestamp
	// Note: We perform summation in Go to ensure big.Int precision as SQLite stores amounts as strings.
	if err := a.db.Where("token = ? AND timestamp >= ?", token.Hex(), since).Find(&withdrawals).Error; err != nil {
		return nil, err
	}

	total := new(big.Int)
	for _, w := range withdrawals {
		amount, ok := new(big.Int).SetString(w.Amount, 10)
		if ok {
			total.Add(total, amount)
		}
	}

	return total, nil
}
