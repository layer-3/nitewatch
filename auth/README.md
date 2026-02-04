# Authentication System Documentation

## Overview

nitewatch implements a secure authentication system using email-based one-time PIN codes with the Resend API as the email provider. The system also supports JWT-based session management with refresh tokens.

## Authentication Flow

### Email One-Time PIN (OTP) Flow

1. **User Initiates Login**
   - User enters their email address
   - Frontend sends request to `/auth/login`

2. **PIN Generation and Delivery**
   - Backend generates a random PIN code (6 digits)
   - PIN is hashed and stored in the `users` table with a timestamp.
   - PIN is sent to user's email via Resend API
   - PIN expires after 10 minutes

3. **PIN Verification**
   - User enters PIN received in email
   - Frontend sends PIN to `/auth/verify`
   - Backend validates PIN against stored hash and timestamp
   - If valid, a new device is created and the backend generates tokens

4. **Token Generation**
   - Upon successful verification, two tokens are generated:
     - **Refresh Token**: Long-lived JWT (7 days), stored as a secure, HTTP-only cookie
     - **Access Token**: Short-lived JWT (5 minutes), returned in response

5. **Session Management**
   - Access token is used for API authentication via Authorization header
   - When access token expires, client uses refresh token to get a new refresh/access token via `/auth/refresh`
   - When user refresh his tokens, the refresh token jti is kept but expiration is extended for another 7 days
   - On logout (`/auth/logout`), provided refresh token is invalidated
   - On logout all devices (`/auth/logout`), refresh token is invalidated for all active devices

## JWT Structure

The `jti` (JWT ID) claim for both access and refresh tokens is a random UUID, ensuring each token is unique.
The `sub` (Subject) claim contains the device token UUID, linking the JWT to a specific device.

### Access Token

```json
{
  "iss": "nitewatch",
  "sub": "<device_uuid>",
  "email": "<user_email>",
  "exp": "<expiration_timestamp>",
  "iat": "<issued_at_timestamp>",
  "jti": "<random_uuid>",
  "type": "access"
}
```

### Refresh Token

```json
{
  "iss": "nitewatch",
  "sub": "<device_uuid>",
  "email": "<user_email>",
  "exp": "<expiration_timestamp>",
  "iat": "<issued_at_timestamp>",
  "jti": "<random_uuid>",
  "type": "refresh"
}
```

## API Endpoints

### Login Endpoints

#### Email Login

```
POST /auth/login
```

**Request:**

```json
{
  "email": "user@example.com"
}
```

**Response:**

```json
{
  "message": "PIN sent to email",
  "expires_in": 600
}
```

#### Verify Email PIN

```
POST /auth/verify
```

**Request:**

```json
{
  "email": "user@example.com",
  "pin": "123456"
}
```

**Response:**

```json
{
  "refresh_token": "<jwt_token>",
  "access_token": "<jwt_token>"
}
```

*Note: Refresh token is set as an HTTP-only cookie*

### Session Management

#### Refresh Token

```
POST /auth/refresh
```

**Request:**

```json
{
  "refresh_token": "<jwt_token>",
  "access_token": "<jwt_token>"
}
```

**Response:**

```json
{
  "refresh_token": "<jwt_token>",
  "access_token": "<new_jwt_token>"
}
```

#### Logout

```
POST /auth/logout
```

**Request:**

```json
{
  revoke: <token_id>,
  "refresh_token": "<jwt_token>",
  "access_token": "<jwt_token>"
}
```

Note: If `revoke: "all"` then all active device tokens will be revoked

**Response:**

```json
{
  "message": "Logged out successfully"
}
```

*Note: This endpoint invalidates the refresh token and clears the cookie*

## Database Schema

The authentication system requires the following database tables:

### Users Table

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    user_pin BIGINT,
    pin_sent_at TIMESTAMP,
    last_seen_at TIMESTAMP,
    notified_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
```

### Devices Table

```sql
CREATE TABLE devices (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL REFERENCES users(id),
  user_agent TEXT NOT NULL, -- User Agent
  token_id VARCHAR(36) NOT NULL, -- Stores the UUID (jti) of the token
  expires_at TIMESTAMP NOT NULL,
  revoked_at TIMESTAMP DEFAULT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
```

## Implementation Details

### Email Service Integration

The system uses Resend API for sending emails. Example integration:

```go
func sendEmailWithPin(email string, pin string) error {
 client := resend.NewClient("your_resend_api_key")

 params := &resend.SendEmailRequest{
  From:    "auth@nitewatch.app",
  To:      []string{email},
  Subject: "Your nitewatch Login PIN",
  Html:    fmt.Sprintf("<p>Your login PIN is: <strong>%s</strong></p><p>This PIN will expire in 10 minutes.</p>", pin),
 }

 _, err := client.Emails.Send(params)
 return err
}
```

## Security Considerations

1. **PIN Security**
   - PINs are stored using secure 64 bit hashing (hash/crc64)
   - Rate limiting is enforced to prevent brute force attacks
   - PINs expire after 10 minutes

2. **Token Security**
   - Refresh tokens are stored in HTTP-only, secure cookies
   - Access tokens are short-lived (5 minutes)
   - Both tokens contain unique identifiers (jti)

3. **Database Security**
   - Sensitive information is encrypted at rest
   - Tokens can be revoked server-side

## Future Enhancements: Passkey Support

In future versions, the system will support WebAuthn/Passkey for passwordless authentication:

1. **Registration**
   - After email verification, users can register a passkey
   - Passkey credentials will be stored securely

2. **Authentication**
   - Users can authenticate using their passkey instead of email OTP
   - Upon successful verification, tokens will be issued as in the email flow

3. **Security Benefits**
   - Phishing-resistant authentication
   - No shared secrets between client and server
   - Device-bound credentials
