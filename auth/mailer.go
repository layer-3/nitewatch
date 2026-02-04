package auth

import "context"

type Mailer interface {
	Send(ctx context.Context, params map[string]any) error
}
