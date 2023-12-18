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

	vd.cfg.DebugLogger.Printf("%s%s", strings.Repeat("  ", vd.logDebugIndent), s)

	if vd.logDebugIndent < 64 && strings.Contains(intro, " begin") {
		vd.logDebugIndent += 1
	}
}

func (vd *Vedirect) ioLogger(prefix string, d []byte) {
	if vd.cfg.IoLogger == nil {
		return
	}
	vd.cfg.IoLogger.Printf("%s: %s\n", prefix, strings.ReplaceAll(string(d), "\n", "\\n"))
}
