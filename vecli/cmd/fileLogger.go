package cmd

import (
	"bufio"
	"fmt"
	"os"
)

type fileLogger struct {
	f *os.File
	w *bufio.Writer
}

func newFileLogger(path string) (*fileLogger, error) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("cannot open io log file: %w", err)
	}
	w := bufio.NewWriter(f)
	return &fileLogger{f: f, w: w}, nil
}

func (l *fileLogger) Println(v ...any) {
	_, err := fmt.Fprintln(l.w, v...)
	if err != nil {
		fmt.Printf("error writing to io log: %s\n", err)
	}
}

func (l *fileLogger) Close() error {
	if err := l.w.Flush(); err != nil {
		return nil
	}
	return l.f.Close()
}
