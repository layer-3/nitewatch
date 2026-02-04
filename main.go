package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/layer-3/nitewatch/config"
	"github.com/layer-3/nitewatch/core/log"
	"github.com/layer-3/nitewatch/handlers"
)

func main() {
	// 1. Load Configuration
	cfg, err := config.LoadAppConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load configuration: %v", err))
	}

	// 2. Initialize Database
	db, err := config.InitDBWithConfig(&cfg.Database, cfg.GetLevel())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize database")
	}

	// 3. Initialize Handlers
	withdrawalHandler := handlers.NewWithdrawalHandler(db, cfg)

	// 4. Setup Router
	gin.SetMode(cfg.GetGinMode())
	r := gin.New()
	r.Use(gin.Recovery())
	// Add custom logger middleware if needed, for now standard Recovery is fine.

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// API Routes
	v1 := r.Group("/v1")
	{
		v1.POST("/withdrawals", withdrawalHandler.InitiateWithdrawal)
		
		// Status Endpoint
		v1.GET("/withdrawals/:id", func(c *gin.Context) {
			// Basic status check implementation inline or move to handler
			// For brevity, using inline lookup
			id := c.Param("id")
			var result struct {
				ID        string
				Status    string
				CreatedAt string
			}
			// Just raw query for speed/simplicity in this main file example
			// In real app, move to handler method
			if err := db.Table("withdrawals").Select("id, status, created_at").Where("id = ?", id).Scan(&result).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Withdrawal not found"})
				return
			}
			c.JSON(http.StatusOK, result)
		})
	}

	// 5. Run Server
	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Info().Msgf("Starting server on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal().Err(err).Msg("Server failed to start")
	}
}
