package vedirect

func (vd *Vedirect) write(b []byte) (n int, err error) {
	if vd.cfg.DebugLogger != nil {
		vd.debugPrintf("write(b=%q) begin len(b)=%d", b, len(b))
		defer func() {
			vd.debugPrintf("write end n=%d err=%v", n, err)
		}()
	}

	n, err = vd.ioPort.Write(b)
	if err != nil {
		return
	}

	if vd.cfg.IoLogger != nil {
		vd.logIoTxBuff = append(vd.logIoTxBuff, b...)
	}
	return
}

func (vd *Vedirect) recvUntil(needle byte) (data []byte, err error) {
	if vd.cfg.DebugLogger != nil {
		vd.debugPrintf("recvUntil(needle=%q) begin", needle)
		defer func() {
			vd.debugPrintf("recvUntil end data=%q err=%v", data, err)
		}()
	}

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
	if vd.cfg.DebugLogger != nil {
		vd.debugPrintf("flushReceiver() begin")
		defer vd.debugPrintf("flushReceiver end")
	}

	if err := vd.ioPort.Flush(); err != nil {
		vd.debugPrintf("err=%v", err)
	}
	vd.reader.Reset(vd.ioPort)
}
