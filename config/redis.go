package config

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strings"

	"github.com/layer-3/nitewatch/core/log"

	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
)

// Redis holds Redis configuration.
type Redis struct {
	URL    string `env:"REDIS_URL"`
	CACert string `env:"REDIS_CACERT"`
}

// GetTLSConfig creates a tls.Config for a Redis connection, handling custom CA certificates.
// It returns nil if the connection is not TLS or no custom CA is provided.
func (c *Redis) GetTLSConfig() (*tls.Config, error) {
	if !strings.HasPrefix(c.URL, "rediss://") {
		return nil, nil
	}

	// Default to insecure for rediss:// if no CA is provided.
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	if c.CACert != "" {
		log.Println("DEBUG: Found REDIS_CACERT, attempting to configure TLS with it.")
		certPool := x509.NewCertPool()
		if ok := certPool.AppendCertsFromPEM([]byte(c.CACert)); !ok {
			return nil, fmt.Errorf("failed to append redis CA cert")
		}
		tlsConfig.RootCAs = certPool
		tlsConfig.InsecureSkipVerify = false // We have a CA, so verify.

		// Extract ServerName from the certificate for SNI
		block, _ := pem.Decode([]byte(c.CACert))
		if block == nil {
			return nil, fmt.Errorf("failed to decode PEM block from redis CA cert")
		}
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse redis CA cert: %w", err)
		}

		if len(cert.IPAddresses) > 0 {
			serverName := cert.IPAddresses[0].String()
			tlsConfig.ServerName = serverName
			log.Printf("DEBUG: Extracted ServerName from certificate SAN: %s", serverName)
		} else if cert.Subject.CommonName != "" {
			serverName := cert.Subject.CommonName
			tlsConfig.ServerName = serverName
			log.Printf("DEBUG: Extracted ServerName from certificate Subject CN: %s", serverName)
		} else {
			log.Println("WARN: Could not determine ServerName from Redis certificate. TLS verification might fail.")
		}
	} else {
		log.Println("WARN: Using rediss:// without a CA certificate (REDIS_CACERT). TLS connection will be insecure.")
	}

	return tlsConfig, nil
}

// GetAsynqRedisConnOpt creates an asynq.RedisConnOpt with support for custom CA certificates.
func (c *Redis) GetAsynqRedisConnOpt() (asynq.RedisConnOpt, error) {
	if c.URL == "" {
		return nil, fmt.Errorf("redis URL is not set")
	}

	opt, err := redis.ParseURL(c.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse redis URL: %w", err)
	}

	tlsConfig, err := c.GetTLSConfig()
	if err != nil {
		return nil, err
	}

	return asynq.RedisClientOpt{
		Addr:      opt.Addr,
		Username:  opt.Username,
		Password:  opt.Password,
		DB:        opt.DB,
		TLSConfig: tlsConfig,
	}, nil
}

// NewRedisClient creates a new Redis client from the configuration.
// It returns nil if no URL is provided.
func (c *Redis) NewRedisClient(ctx context.Context) (*redis.Client, error) {
	if c.URL == "" {
		return nil, nil
	}

	opt, err := redis.ParseURL(c.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse redis URL: %w", err)
	}

	tlsConfig, err := c.GetTLSConfig()
	if err != nil {
		return nil, err
	}
	opt.TLSConfig = tlsConfig

	client := redis.NewClient(opt)
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return client, nil
}
