package log

import "log"

// Logger is the interface for a logger. It is implemented by e.g. log.Logger.
type Logger interface {
	Printf(format string, a ...any)
}

type NoOppLogger struct{}

func (NoOppLogger) Printf(string, ...any) {}

type DefaultLogger struct {
	prefix string
}

func (l DefaultLogger) Printf(format string, a ...any) {
	log.Printf(l.prefix+format, a...)
}
