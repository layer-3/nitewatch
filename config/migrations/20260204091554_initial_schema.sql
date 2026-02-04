-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    pin_sent_at TIMESTAMP WITH TIME ZONE,
    last_seen_at TIMESTAMP WITH TIME ZONE,
    notified_at TIMESTAMP WITH TIME ZONE,
    email TEXT NOT NULL UNIQUE,
    user_pin BIGINT,
    metadata JSONB
);

CREATE INDEX idx_users_email ON users(email);

CREATE TABLE devices (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    revoked_at TIMESTAMP WITH TIME ZONE,
    user_agent TEXT DEFAULT '',
    token UUID NOT NULL
);

CREATE INDEX idx_devices_token ON devices(token);
CREATE INDEX idx_devices_user_id ON devices(user_id);

CREATE TABLE withdrawals (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Payload Data
    user_address TEXT NOT NULL,
    token_address TEXT NOT NULL,
    amount TEXT NOT NULL,
    chain_id BIGINT NOT NULL,
    nonce BIGINT NOT NULL,
    email TEXT NOT NULL,

    -- Signatures
    user_signature TEXT NOT NULL,
    broker_signature TEXT NOT NULL,
    nitewatch_signature TEXT,

    -- State
    status TEXT NOT NULL DEFAULT 'created' CHECK (status IN ('created', 'authorized', 'approved', 'rejected', 'failed')),
    error_code INTEGER,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE INDEX idx_withdrawals_email ON withdrawals(email);
CREATE INDEX idx_withdrawals_status ON withdrawals(status);
CREATE INDEX idx_withdrawals_error_code ON withdrawals(error_code);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS withdrawals;
DROP TABLE IF EXISTS devices;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd