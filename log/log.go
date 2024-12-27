package log

// Logger is the interface for a logger. It is implemented by e.g. log.Logger.
type Logger interface {
	Printf(format string, a ...any)
}

type NoOppLogger struct{}

func (NoOppLogger) Printf(string, ...any) {}
