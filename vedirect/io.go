package vedirect

import "log"

func (vd *Vedirect) write(b []byte) (n int, err error) {
	vd.debugPrintf("vedirect: write b=%s len=%v", b, len(b))
	n, err = vd.cfg.ioPort.Write(b)
	if err != nil {
		log.Printf("vedirect: write error: %v\n", err)
		return 0, err
	}
	if vd.cfg.ioLogger != nil {
		vd.lastWritten = b
	}

	return
}

func (vd *Vedirect) recvUntil(needle byte) (data []byte, err error) {
	vd.debugPrintf("vedirect: recvUntil needle=%c", needle)
	data, err = vd.reader.ReadBytes(needle)
	if err == nil {
		data = data[:len(data)-1] // exclude delimiter
		vd.printIO(data)
	}
	return
}

// FlushReceiver flushes the underlying receiver buffer. This after some inactivity since some devices
// like the BMV will start sending asynchronous messages after a while.
func (vd *Vedirect) flushReceiver() {
	vd.debugPrintf("vedirect: FlushReceiver begin")

	if err := vd.cfg.ioPort.Flush(); err != nil {
		vd.debugPrintf("vedirect: FlushReceiver err=%v", err)
	}
	vd.reader.Reset(vd.cfg.ioPort)

	vd.debugPrintf("vedirect: FlushReceiver end")
}
