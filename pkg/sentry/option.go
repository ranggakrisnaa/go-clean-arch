package sentry

import "time"

type Option func(*Sentry)

func DSN(dsn string) Option {
	return func(s *Sentry) {
		s.DSN = dsn
	}
}

func Environment(env string) Option {
	return func(s *Sentry) {
		s.Environment = env
	}
}

func SampleRate(rate float64) Option {
	return func(s *Sentry) {
		s.SampleRate = rate
	}
}

func FlushTime(duration time.Duration) Option {
	return func(s *Sentry) {
		s.FlushTime = duration
	}
}
