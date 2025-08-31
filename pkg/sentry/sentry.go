package sentry

import (
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
	sentryfiber "github.com/getsentry/sentry-go/fiber"
	"github.com/gofiber/fiber/v2"
)

const (
	_defaultDSN         = ""
	_defaultEnvironment = "development"
	_defaultSampleRate  = 1.0
	_defaultFlushTime   = 2 * time.Second
)

// Sentry -.
type Sentry struct {
	DSN         string
	Environment string
	SampleRate  float64
	FlushTime   time.Duration
}

// New -.
func New(opts ...Option) (*Sentry, error) {
	s := &Sentry{
		DSN:         _defaultDSN,
		Environment: _defaultEnvironment,
		SampleRate:  _defaultSampleRate,
		FlushTime:   _defaultFlushTime,
	}
	for _, opt := range opts {
		opt(s)
	}

	defer sentry.Flush(s.FlushTime)

	err := sentry.Init(sentry.ClientOptions{
		Dsn:         s.DSN,
		Environment: s.Environment,
		SampleRate:  s.SampleRate,
	})
	if err != nil {
		return nil, fmt.Errorf("sentry - New - sentry.Init: %w", err)
	}

	return s, nil
}

// Middleware -.
func (s *Sentry) Middleware() fiber.Handler {
	return sentryfiber.New(sentryfiber.Options{
		Repanic: true,
	})
}

// Flush -.
func (s *Sentry) Flush() {
	sentry.Flush(s.FlushTime)
}
