package gotrace

import (
	"time"

	"golang.org/x/exp/slog"
)

type traceEntry struct {
	r     slog.Record
	start time.Time
}

func Trace(msg string, kvs ...any) (v traceEntry) {
	slog.Info(msg, kvs...)
	v.r.Message = msg
	v.start = time.Now()
	return v
}

func (v traceEntry) Stop(err *error) {
	if err == nil || *err == nil {
		slog.Info(v.r.Message, "duration", time.Since(v.start).Milliseconds())
	} else {
		slog.Error(v.r.Message, *err, "duration", time.Since(v.start).Milliseconds())
	}
}
