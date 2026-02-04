package log_test

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/layer-3/nitewatch/core/log"
	"github.com/rs/zerolog"
)

func TestGlobalLogFunctions(t *testing.T) {
	// Backup and restore the global logger
	originalLogger := log.GetLogger()
	defer func() { log.SetLogger(originalLogger) }()

	// Use a buffer to capture logs
	var buf bytes.Buffer
	log.SetLogger(zerolog.New(&buf))

	tests := []struct {
		name     string
		logFunc  func()
		expected string
		panic    bool
	}{
		{
			name: "Trace",
			logFunc: func() {
				log.Trace().Msg("trace msg")
			},
			expected: `"level":"trace","message":"trace msg"`,
		},
		{
			name: "Debug",
			logFunc: func() {
				log.Debug().Msg("debug msg")
			},
			expected: `"level":"debug","message":"debug msg"`,
		},
		{
			name: "Info",
			logFunc: func() {
				log.Info().Msg("info msg")
			},
			expected: `"level":"info","message":"info msg"`,
		},
		{
			name: "Warn",
			logFunc: func() {
				log.Warn().Msg("warn msg")
			},
			expected: `"level":"warn","message":"warn msg"`,
		},
		{
			name: "Error",
			logFunc: func() {
				log.Error().Msg("error msg")
			},
			expected: `"level":"error","message":"error msg"`,
		},
		{
			name: "Err",
			logFunc: func() {
				log.Err(errors.New("my error")).Msg("oops")
			},
			expected: `"error":"my error","message":"oops"`,
		},
		{
			name: "WithLevel",
			logFunc: func() {
				log.WithLevel(zerolog.InfoLevel).Msg("custom level")
			},
			expected: `"level":"info","message":"custom level"`,
		},
		{
			name: "Log",
			logFunc: func() {
				log.Log().Msg("no level")
			},
			expected: `"message":"no level"`,
		},
		{
			name: "Print",
			logFunc: func() {
				log.Print("hello print")
			},
			expected: `"level":"debug","message":"hello print"`,
		},
		{
			name: "Printf",
			logFunc: func() {
				log.Printf("hello %s", "printf")
			},
			expected: `"level":"debug","message":"hello printf"`,
		},
		{
			name: "Warnf",
			logFunc: func() {
				log.Warnf("hello %s", "warnf")
			},
			expected: `"level":"warn","message":"hello warnf"`,
		},
		{
			name: "Errorf",
			logFunc: func() {
				log.Errorf("hello %s", "errorf")
			},
			expected: `"level":"error","message":"hello errorf"`,
		},
		{
			name: "Panic",
			logFunc: func() {
				defer func() { _ = recover() }()
				log.Panic().Msg("panic msg")
			},
			expected: `"level":"panic","message":"panic msg"`,
			panic:    true,
		},
	}

	// We need to set the global level to Trace to ensure all levels are logged
	// Since log.Logger is a zerolog.Logger, it respects zerolog.GlobalLevel()
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	defer zerolog.SetGlobalLevel(zerolog.InfoLevel) // Restore default

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.Reset()
			tt.logFunc()
			if !strings.Contains(buf.String(), tt.expected) {
				t.Errorf("expected log to contain %s, got %s", tt.expected, buf.String())
			}
		})
	}
}

func TestHelpers(t *testing.T) {
	// Backup and restore the global logger
	originalLogger := log.GetLogger()
	defer func() { log.SetLogger(originalLogger) }()

	var buf bytes.Buffer
	log.SetLogger(zerolog.New(&buf))

	t.Run("Output", func(t *testing.T) {
		var newBuf bytes.Buffer
		// Output returns a new logger, verify it writes to the new writer
		l := log.Output(&newBuf)
		l.Info().Msg("test output")

		if !strings.Contains(newBuf.String(), "test output") {
			t.Errorf("expected new output to receive log")
		}
		if buf.String() != "" {
			t.Errorf("expected original output to be empty")
		}
	})

	t.Run("With", func(t *testing.T) {
		// With returns a context from the global logger
		// We add a field and create a logger
		l := log.With().Str("foo", "bar").Logger()

		// To verify, we need to capture output.
		// Since l is derived from log.Logger (which writes to buf), l should write to buf too.
		l.Info().Msg("test with")

		if !strings.Contains(buf.String(), `"foo":"bar"`) {
			t.Errorf("expected log to contain context field")
		}
	})

	t.Run("Level", func(t *testing.T) {
		buf.Reset()
		// Level returns a logger with minimum level
		l := log.Level(zerolog.WarnLevel)

		l.Info().Msg("should be ignored")
		if buf.String() != "" {
			t.Errorf("expected Info to be ignored")
		}

		l.Warn().Msg("should be logged")
		if !strings.Contains(buf.String(), "should be logged") {
			t.Errorf("expected Warn to be logged")
		}
	})

	t.Run("Sample", func(t *testing.T) {
		buf.Reset()
		// Sample returns a logger with sampler
		l := log.Sample(&zerolog.BasicSampler{N: 1})
		l.Info().Msg("sampled")
		if !strings.Contains(buf.String(), "sampled") {
			t.Errorf("expected sampled message")
		}
	})

	t.Run("Hook", func(t *testing.T) {
		called := false
		hook := zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, message string) {
			called = true
		})

		l := log.Hook(hook)
		l.Info().Msg("trigger hook")

		if !called {
			t.Errorf("expected hook to be called")
		}
	})
}

func TestCtx(t *testing.T) {
	// Test retrieving logger from context
	var buf bytes.Buffer
	l := zerolog.New(&buf).With().Str("ctx_field", "exists").Logger()

	ctx := l.WithContext(context.Background())

	// Use core/log.Ctx to retrieve it
	retrievedLogger := log.Ctx(ctx)
	retrievedLogger.Info().Msg("from context")

	if !strings.Contains(buf.String(), `"ctx_field":"exists"`) {
		t.Errorf("expected logger from context to preserve fields")
	}
	if !strings.Contains(buf.String(), "from context") {
		t.Errorf("expected message to be logged to the logger from context")
	}
}
