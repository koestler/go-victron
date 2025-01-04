package vebleapi

type Listener struct {
	packets     chan Packet
	unsubscribe func()
}

func NewListener() *Listener {
	return &Listener{
		packets: make(chan Packet),
	}
}

func (l *Listener) Drain() <-chan Packet {
	return l.packets
}

func (l *Listener) End() {
	if l.unsubscribe != nil {
		l.unsubscribe()
	}
	l.close()
}

func (l *Listener) close() {
	close(l.packets)
	l.packets = nil
}
