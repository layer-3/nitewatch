package auth

import (
	"context"
	"fmt"

	"github.com/layer-3/nitewatch/core/log"
)

// consoleMailer handles email operations by logging to console for local development
type consoleMailer struct {
	appName string
}

// newConsoleMailer creates a new console mailer instance
func newConsoleMailer() *consoleMailer {
	return &consoleMailer{
		appName: "nitewatch",
	}
}

// Send implements the Mailer interface for console output
func (c *consoleMailer) Send(ctx context.Context, params map[string]any) error {
	email, ok := params["email"].(string)
	if !ok {
		return fmt.Errorf("email parameter is required")
	}

	pin, ok := params["pin"].(string)
	if !ok {
		return fmt.Errorf("pin parameter is required")
	}

	// Log the PIN to console for local development
	log.Printf("\n"+
		"========================================\n"+
		"AUTH PIN CODE (LOCAL DEVELOPMENT)\n"+
		"========================================\n"+
		"To: %s\n"+
		"PIN: %s\n"+
		"Expires in: 10 minutes\n"+
		"========================================\n",
		email, pin)

	return nil
}
