package vedirect

import "log"

func (vd *Vedirect) write(b []byte) (n int, err error) {
	vd.debugPrintf("write b=%s len=%v", b, len(b))
	n, err = vd.cfg.IOPort.Write(b)
	if err != nil {
		log.Printf("write error: %v\n", err)
		return 0, err
	}
	if vd.cfg.IoLogger != nil {
		vd.logIoTxBuff = append(vd.logIoTxBuff, b...)
	}
	return
}

func (vd *Vedirect) recvUntil(needle byte) (data []byte, err error) {
	vd.debugPrintf("recvUntil needle=%c", needle)
	data, err = vd.reader.ReadBytes(needle)
	if err == nil {
		if vd.cfg.IoLogger != nil {
			vd.logIoRxBuff = append(vd.logIoRxBuff, data...)
		}
		data = data[:len(data)-1] // exclude delimiter
	}
	return
}

// FlushReceiver flushes the underlying receiver buffer. This after some inactivity since some devices
// like the BMV will start sending asynchronous messages after a while.
func (vd *Vedirect) flushReceiver() {
	vd.debugPrintf("flushReceiver begin")

	if err := vd.cfg.IOPort.Flush(); err != nil {
		vd.debugPrintf("flushReceiver err=%v", err)
	}
	vd.reader.Reset(vd.cfg.IOPort)

	vd.debugPrintf("flushReceiver end")
}
