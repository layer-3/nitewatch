package auth

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"html/template"
	text_template "text/template"

	"github.com/resend/resend-go/v2"
)

// resendMailer handles email operations
type resendMailer struct {
	client    *resend.Client
	fromEmail string
	appName   string
}

// newResendMailer creates a new resend mailer instance
func newResendMailer(apiKey string, fromEmail string) *resendMailer {
	if fromEmail == "" {
		fromEmail = "auth@nitewatch.app"
	}

	return &resendMailer{
		client:    resend.NewClient(apiKey),
		fromEmail: fromEmail,
		appName:   "nitewatch",
	}
}

//go:embed login.html.tmpl
var loginHTMLTemplate string

// Send implements the Mailer interface
func (r *resendMailer) Send(ctx context.Context, params map[string]any) error {
	email, ok := params["email"].(string)
	if !ok {
		return fmt.Errorf("email parameter is required")
	}

	pin, ok := params["pin"].(string)
	if !ok {
		return fmt.Errorf("pin parameter is required")
	}

	return r.sendPINEmail(email, pin)
}

// sendPINEmail sends a PIN code to the user's email
func (r *resendMailer) sendPINEmail(email string, pin string) error {
	// Data for templates
	data := struct {
		AppName string
		PIN     string
	}{
		AppName: r.appName,
		PIN:     pin,
	}

	// Generate HTML content
	htmlTmpl, err := template.New("html").Parse(loginHTMLTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse HTML template: %w", err)
	}

	var htmlBuffer bytes.Buffer
	if err := htmlTmpl.Execute(&htmlBuffer, data); err != nil {
		return fmt.Errorf("failed to execute HTML template: %w", err)
	}

	// Generate text content
	textTemplate := `Your {{.AppName}} Login PIN

You requested to log in to {{.AppName}}. Use the PIN below to complete your login:

{{.PIN}}

This PIN will expire in 10 minutes.

If you didn't request this login, please ignore this email.

This is an automated message from {{.AppName}}. Please do not reply to this email.`

	textTmpl, err := text_template.New("text").Parse(textTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse text template: %w", err)
	}

	var textBuffer bytes.Buffer
	if err := textTmpl.Execute(&textBuffer, data); err != nil {
		return fmt.Errorf("failed to execute text template: %w", err)
	}

	params := &resend.SendEmailRequest{
		From:    r.fromEmail,
		To:      []string{email},
		Subject: fmt.Sprintf("Your %s Login PIN", r.appName),
		Html:    htmlBuffer.String(),
		Text:    textBuffer.String(),
	}

	_, err = r.client.Emails.Send(params)
	return err
}
