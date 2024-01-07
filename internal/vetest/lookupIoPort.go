package vetest

import (
	"bytes"
	"testing"
)

type LookupIOPort struct {
	t           *testing.T
	lookupTable map[string]string
	lookupCount map[string]uint
	closed      bool
	buff        bytes.Buffer
}

func NewLookupIOPort(t *testing.T, lookupTable map[string]string) *LookupIOPort {
	return &LookupIOPort{
		t:           t,
		lookupTable: lookupTable,
		lookupCount: make(map[string]uint),
	}
}

func (l *LookupIOPort) lookup(tx []byte) (rx []byte, ok bool) {
	v, ok := l.lookupTable[string(tx)]
	if ok {
		l.lookupCount[string(tx)]++
	}
	return []byte(v), ok
}

func (l *LookupIOPort) Write(b []byte) (n int, err error) {
	if rx, ok := l.lookup(b); !ok {
		l.t.Errorf("LookupIOPort: Write: lookup failed: len(b)=%d, b=%s", len(b), b)
	} else {
		l.buff.Write(rx)
	}

	return len(b), nil
}

func (l *LookupIOPort) Read(b []byte) (n int, err error) {
	n, err = l.buff.Read(b)
	return
}

func (l *LookupIOPort) Close() error {
	l.closed = true
	return nil
}

func (l *LookupIOPort) CheckClosed() {
	if !l.closed {
		l.t.Errorf("LookupIOPort: not closed")
	}
}

func (l *LookupIOPort) CheckEverythingHeard() {
	for k := range l.lookupTable {
		if _, ok := l.lookupCount[k]; !ok {
			l.t.Errorf("LookupIOPort: not heard: tx=%s", k)
		}
	}
}

func (l *LookupIOPort) Flush() error {
	l.buff.Truncate(0)
	return nil
}
