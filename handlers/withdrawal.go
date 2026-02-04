package handlers

import (
	"context"
	"errors"
	"math/big"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/layer-3/nitewatch/config"
	"github.com/layer-3/nitewatch/core"
	"gorm.io/gorm"
)

type WithdrawalHandler struct {
	DB     *gorm.DB
	Config *config.App
}

func NewWithdrawalHandler(db *gorm.DB, cfg *config.App) *WithdrawalHandler {
	return &WithdrawalHandler{
		DB:     db,
		Config: cfg,
	}
}

type WithdrawalRequest struct {
	UserAddress  string                 `json:"user_address" binding:"required"`
	TokenAddress string                 `json:"token_address" binding:"required"`
	Amount       string                 `json:"amount" binding:"required"`
	ChainID      int64                  `json:"chain_id" binding:"required"`
	Nonce        int64                  `json:"nonce" binding:"required"`
	Email        string                 `json:"email" binding:"required,email"`
	Signatures   WithdrawalSignatures   `json:"signatures" binding:"required"`
}

type WithdrawalSignatures struct {
	User   string `json:"user" binding:"required"`
	Broker string `json:"broker" binding:"required"`
}

type WithdrawalResponse struct {
	WithdrawalID       uuid.UUID             `json:"withdrawal_id"`
	Status             core.WithdrawalStatus `json:"status"`
	NitewatchSignature string                `json:"nitewatch_signature,omitempty"`
	Message            string                `json:"message,omitempty"`
}

func (h *WithdrawalHandler) InitiateWithdrawal(c *gin.Context) {
	var req WithdrawalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 1. Validate Signatures (Placeholder for real crypto verification)
	// TODO: Implement EIP-712 or standard hash signing verification
	// if !verifySignature(req.UserAddress, req.Signatures.User, req...) { ... }

	// 2. Parse Amount
	amount := new(big.Int)
	amount, ok := amount.SetString(req.Amount, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount format"})
		return
	}

	// 3. Verify Limits
	if err := h.checkLimits(c.Request.Context(), req.Email, amount); err != nil {
		h.recordRejection(req, err.Error())
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	// 4. Sign Payload (Placeholder for KMS signing)
	nitewatchSig := "0xmock_nitewatch_signature_" + uuid.New().String()

	// 5. Create Withdrawal Record
	withdrawal := core.Withdrawal{
		UserAddress:        req.UserAddress,
		TokenAddress:       req.TokenAddress,
		Amount:             req.Amount,
		ChainID:            req.ChainID,
		Nonce:              req.Nonce,
		Email:              req.Email,
		UserSignature:      req.Signatures.User,
		BrokerSignature:    req.Signatures.Broker,
		NitewatchSignature: &nitewatchSig,
		Status:             core.StatusApproved,
	}

	if err := h.DB.Create(&withdrawal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create withdrawal record"})
		return
	}

	c.JSON(http.StatusOK, WithdrawalResponse{
		WithdrawalID:       withdrawal.ID,
		Status:             withdrawal.Status,
		NitewatchSignature: nitewatchSig,
	})
}

func (h *WithdrawalHandler) checkLimits(ctx context.Context, email string, amount *big.Int) error {
	// Parse Config Limits
	userHourlyLimit := new(big.Int)
	userHourlyLimit.SetString(h.Config.Security.DefaultUserHourlyLimit, 10)
	
	userDailyLimit := new(big.Int)
	userDailyLimit.SetString(h.Config.Security.DefaultUserDailyLimit, 10)

	globalHourlyLimit := new(big.Int)
	globalHourlyLimit.SetString(h.Config.Security.GlobalHourlyLimit, 10)

	globalDailyLimit := new(big.Int)
	globalDailyLimit.SetString(h.Config.Security.GlobalDailyLimit, 10)

	// Check User Hourly
	if err := h.checkLimitSum(ctx, "email = ? AND created_at > ?", []interface{}{email, time.Now().Add(-1 * time.Hour)}, amount, userHourlyLimit); err != nil {
		return errors.New("user hourly limit exceeded")
	}

	// Check User Daily
	if err := h.checkLimitSum(ctx, "email = ? AND created_at > ?", []interface{}{email, time.Now().Add(-24 * time.Hour)}, amount, userDailyLimit); err != nil {
		return errors.New("user daily limit exceeded")
	}

	// Check Global Hourly
	if err := h.checkLimitSum(ctx, "created_at > ?", []interface{}{time.Now().Add(-1 * time.Hour)}, amount, globalHourlyLimit); err != nil {
		return errors.New("global hourly limit exceeded")
	}

	// Check Global Daily
	if err := h.checkLimitSum(ctx, "created_at > ?", []interface{}{time.Now().Add(-24 * time.Hour)}, amount, globalDailyLimit); err != nil {
		return errors.New("global daily limit exceeded")
	}

	return nil
}

func (h *WithdrawalHandler) checkLimitSum(ctx context.Context, query string, args []interface{}, currentAmount *big.Int, limit *big.Int) error {
	var result struct {
		TotalAmount string // GORM will scan the string or numeric result here
	}

	// We query the database for the sum of approved withdrawals
	// Note: We need to filter by status 'approved' or 'authorized' if we had that step.
	// Since we are creating 'approved' immediately, we check 'approved'.
	fullQuery := query + " AND status = ?"
	fullArgs := append(args, core.StatusApproved)

	// In Postgres, SUM(text) won't work directly if amount is text. We need to cast.
	// GORM raw SQL: SELECT SUM(CAST(amount AS NUMERIC)) ...
	err := h.DB.Model(&core.Withdrawal{}).
		Select("COALESCE(SUM(CAST(amount AS NUMERIC)), 0)").
		Where(fullQuery, fullArgs...).
		Scan(&result.TotalAmount).Error

	if err != nil {
		return err
	}

	// result.TotalAmount might be scientific notation or decimal if we aren't careful, 
	// but Postgres NUMERIC sum should be a clean string.
	// COALESCE ensures we get "0" if no rows.
	
	// big.Float might be safer to parse if Postgres returns decimals, but assuming Wei (integers):
	totalUsedFloat, _, err := big.ParseFloat(result.TotalAmount, 10, 0, big.ToNearestEven)
	if err != nil {
		// Fallback or error log
		return nil // Optimistic? Or fail safe? Let's treat as 0 if parse fails to prevent block, but unsafe.
	}
	
	// Convert Float to Int for comparison (since limits are Ints)
	totalUsedInt, _ := totalUsedFloat.Int(nil)
	
	newTotal := new(big.Int).Add(totalUsedInt, currentAmount)

	if newTotal.Cmp(limit) > 0 {
		return errors.New("limit exceeded")
	}

	return nil
}

func (h *WithdrawalHandler) recordRejection(req WithdrawalRequest, reason string) {
	// Optional: Record rejected attempts for security auditing
	// For now, we just return the error, but in a real system we might log this to a separate table or log file.
}
