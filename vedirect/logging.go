package vedirect

import (
	"fmt"
	"strings"
)

func (vd *Vedirect) debugPrintf(format string, v ...interface{}) {
	// check if debug output is enabled
	if vd.cfg.DebugLogger == nil {
		return
	}

	intro := strings.Split(format, "=")[0]

	if vd.logDebugIndent > 0 && strings.Contains(intro, " end") {
		vd.logDebugIndent -= 1
	}

	s := fmt.Sprintf(format, v...)
	s = strings.Replace(s, "\n", "\\n", -1)

	vd.cfg.DebugLogger.Println(strings.Repeat("  ", vd.logDebugIndent), s)

	if vd.logDebugIndent < 64 && strings.Contains(intro, " begin") {
		vd.logDebugIndent += 1
	}
}

func (vd *Vedirect) ioLoggerLineEnd(commentFormat string, commentParams ...any) {
	if vd.cfg.IoLogger == nil {
		return
	}
	vd.cfg.IoLogger.Println(fmt.Sprintf("%q: %q, // %s",
		vd.ioLogTxBuff,
		vd.ioLogRxBuff,
		fmt.Sprintf(commentFormat, commentParams...),
	))
	vd.clearIoLogBuffers()
}
