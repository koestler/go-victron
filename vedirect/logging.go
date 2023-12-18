package vedirect

import (
	"fmt"
	"strings"
)

func (vd *Vedirect) debugPrintf(format string, v ...interface{}) {
	// check if debug output is enabled
	if vd.cfg.debugLogger == nil {
		return
	}

	intro := strings.Split(format, "=")[0]

	if vd.logDebugIndent > 0 && strings.Contains(intro, " end") {
		vd.logDebugIndent -= 1
	}

	s := fmt.Sprintf(format, v...)
	s = strings.Replace(s, "\n", "\\n", -1)

	_, _ = fmt.Fprint(vd.cfg.debugLogger, strings.Repeat("  ", vd.logDebugIndent)+s+"\n")

	if vd.logDebugIndent < 64 && strings.Contains(intro, " begin") {
		vd.logDebugIndent += 1
	}
}

func (vd *Vedirect) printIO(oup []byte) {
	if vd.cfg.ioLogger == nil {
		return
	}

	_, _ = fmt.Fprintf(vd.cfg.ioLogger, "\"%x\": \"%x\"\n", vd.lastWritten, oup)
}
