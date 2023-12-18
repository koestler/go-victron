package vedirect_test

import (
	"bytes"
	"github.com/koestler/go-victron/vedirect"
	"log"
	"testing"
)

func TestVedirect(t *testing.T) {
	io := NewLookupIOPort(t, map[string]string{
		":154\n": ":51641F9\n",
	})
	defer io.CheckEverythingHeard()
	defer io.CheckClosed()

	vd, err := vedirect.NewVedirect(&vedirect.Config{
		io,
		log.Default(),
		log.Default(),
	})

	if err != nil {
		t.Fatalf("cannot create vedirect: %v", err)
	}

	if err := vd.Ping(); err != nil {
		t.Fatalf("cannot ping: %v", err)
	}

	if err := io.Close(); err != nil {
		t.Fatalf("cannot close io: %v", err)
	}
}

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
		l.t.Logf("write to buffer: %s", rx)
		l.buff.Write(rx)
	}

	return len(b), nil
}

func (l *LookupIOPort) Read(b []byte) (n int, err error) {
	n, err = l.buff.Read(b)
	l.t.Logf("read from buffer: b=%s, n=%d, err=%s", b[:n], n, err)
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
