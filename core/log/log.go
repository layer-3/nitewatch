// Package log provides a global logger for zerolog.
package log

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog"
)

// LogLevel defines the severity of log messages
type LogLevel string

const (
	LevelDebug LogLevel = "debug"
	LevelInfo  LogLevel = "info"
	LevelWarn  LogLevel = "warn"
	LevelError LogLevel = "error"
)

// logger is the global logger.
var logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

// SetLogger sets the global logger.
func SetLogger(l zerolog.Logger) {
	logger = l
}

// SetLevel sets the global log level.
func SetLevel(level LogLevel) error {
	l, err := zerolog.ParseLevel(string(level))
	if err != nil {
		return err
	}
	zerolog.SetGlobalLevel(l)
	return nil
}

// GetLogger returns the global logger.
func GetLogger() zerolog.Logger {
	return logger
}

// Output duplicates the global logger and sets w as its output.
func Output(w io.Writer) zerolog.Logger {
	return logger.Output(w)
}

// EnableConsoleLogger switches the global logger to output formatted logs to stderr.
func EnableConsoleLogger() {
	logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

// With creates a child logger with the field added to its context.
func With() zerolog.Context {
	return logger.With()
}

// Level creates a child logger with the minimum accepted level set to level.
func Level(level zerolog.Level) zerolog.Logger {
	return logger.Level(level)
}

// Sample returns a logger with the s sampler.
func Sample(s zerolog.Sampler) zerolog.Logger {
	return logger.Sample(s)
}

// Hook returns a logger with the h Hook.
func Hook(h zerolog.Hook) zerolog.Logger {
	return logger.Hook(h)
}

// Err starts a new message with error level with err as a field if not nil or
// with info level if err is nil.
//
// You must call Msg on the returned event in order to send the event.
func Err(err error) *zerolog.Event {
	return logger.Err(err)
}

// Trace starts a new message with trace level.
//
// You must call Msg on the returned event in order to send the event.
func Trace() *zerolog.Event {
	return logger.Trace()
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func Debug() *zerolog.Event {
	return logger.Debug()
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func Info() *zerolog.Event {
	return logger.Info()
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func Warn() *zerolog.Event {
	return logger.Warn()
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func Error() *zerolog.Event {
	return logger.Error()
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func Fatal() *zerolog.Event {
	return logger.Fatal()
}

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func Panic() *zerolog.Event {
	return logger.Panic()
}

// WithLevel starts a new message with level.
//
// You must call Msg on the returned event in order to send the event.
func WithLevel(level zerolog.Level) *zerolog.Event {
	return logger.WithLevel(level)
}

// Log starts a new message with no level. Setting zerolog.GlobalLevel to
// zerolog.Disabled will still disable events produced by this method.
//
// You must call Msg on the returned event in order to send the event.
func Log() *zerolog.Event {
	return logger.Log()
}

// Print sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...any) {
	logger.Debug().CallerSkipFrame(1).Msg(fmt.Sprint(v...))
}

// Println is an alias of Print
func Println(v ...any) {
	Print(v...)
}

// Printf sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...any) {
	logger.Debug().CallerSkipFrame(1).Msgf(format, v...)
}

// Infof sends a log event using info level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func Infof(format string, v ...any) {
	logger.Info().CallerSkipFrame(1).Msgf(format, v...)
}

// Warnf sends a log event using warn level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func Warnf(format string, v ...any) {
	logger.Warn().CallerSkipFrame(1).Msgf(format, v...)
}

// Errorf sends a log event using error level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, v ...any) {
	logger.Error().CallerSkipFrame(1).Msgf(format, v...)
}

// Fatalf sends a log event using fatal level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func Fatalf(format string, v ...any) {
	logger.Fatal().CallerSkipFrame(1).Msgf(format, v...)
}

// Ctx returns the Logger associated with the ctx. If no logger
// is associated, a disabled logger is returned.
func Ctx(ctx context.Context) *zerolog.Logger {
	return zerolog.Ctx(ctx)
}
