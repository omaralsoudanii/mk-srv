package zlog

import (
	"encoding/json"
	"fmt"
	"os"
)

func defaultLog(level, msg string, fields ...Field) {
	log := struct {
		Level   string  `json:"level"`
		Message string  `json:"msg"`
		Fields  []Field `json:"fields,omitempty"`
	}{
		level,
		msg,
		fields,
	}

	if b, err := json.Marshal(log); err != nil {
		fmt.Fprintf(os.Stderr, `{"level":"error","msg":"failed to encode log message"}%s`, "\n")
	} else {
		fmt.Fprintf(os.Stderr, "%s\n", b)
	}
}

func defaultDebugLog(msg string, fields ...Field) {
	defaultLog("debug", msg, fields...)
}

func defaultInfoLog(msg string, fields ...Field) {
	defaultLog("info", msg, fields...)
}

func defaultWarnLog(msg string, fields ...Field) {
	defaultLog("warn", msg, fields...)
}

func defaultErrorLog(msg string, fields ...Field) {
	defaultLog("error", msg, fields...)
}

func defaultCriticalLog(msg string, fields ...Field) {
	defaultLog("error", msg, fields...)
}
