package vedirect

func (vd *Vedirect) write(b []byte) (n int, err error) {
	if vd.cfg.DebugLogger != nil {
		vd.debugPrintf("write(b=%s) begin len(b)=%d", b, len(b))
		defer func() {
			vd.debugPrintf("write end n=%d err=%s", n, err)
		}()
	}

	n, err = vd.cfg.IOPort.Write(b)
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
		vd.debugPrintf("recvUntil(needle=%c) begin", needle)
		defer func() {
			vd.debugPrintf("recvUntil end data=%s err=%s", data, err)
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

	if err := vd.cfg.IOPort.Flush(); err != nil {
		vd.debugPrintf("err=%v", err)
	}
	vd.reader.Reset(vd.cfg.IOPort)
}
