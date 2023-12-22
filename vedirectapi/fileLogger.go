package vedirectapi

import (
	"bufio"
	"fmt"
	"os"
)

// FileLogger implements the vedirect.Logger interface and simply writes all log lines to a file.
// Usage:
//
//	logger, err := NewFileLogger("io.log")
//	if err != nil {
//	  panic(err)
//	}
//	defer logger.Close()
type FileLogger struct {
	f *os.File
	w *bufio.Writer
}

// NewFileLogger creates a new FileLogger instance.
func NewFileLogger(path string) (*FileLogger, error) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("cannot open io log file: %w", err)
	}
	w := bufio.NewWriter(f)
	return &FileLogger{f: f, w: w}, nil
}

// Println writes the given values to the log file.
func (l *FileLogger) Println(v ...any) {
	_, err := fmt.Fprintln(l.w, v...)
	if err != nil {
		fmt.Printf("error writing to io log: %s\n", err)
	}
}

// Close flushes the log file and closes it.
func (l *FileLogger) Close() error {
	if err := l.w.Flush(); err != nil {
		return nil
	}
	return l.f.Close()
}
