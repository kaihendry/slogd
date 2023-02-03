package gotrace

import (
	"time"

	"golang.org/x/exp/slog"
)

type traceEntry struct {
	Record    slog.Record
	StartedAt time.Time
}

func Trace(msg string, args ...any) traceEntry {
	var v traceEntry
	slog.Info(msg, args...)
	v.Record.Message = msg
	// it doesn't seem trivial to pass args into Record?
	// https://cs.opensource.google/go/x/exp/+/54bba9f4:slog/record.go;drc=f062dba9d201f5ec084d25785efec05637818c00;l=158
	v.StartedAt = time.Now()
	return v
}

func (v traceEntry) Stop(err error) {
	if err == nil {
		// I'm using slog.Duration, though I can't make sense of the output. Prefer just milliseconds.
		// sidenote: maybe i should override to level = trace?
		slog.Info(v.Record.Message, slog.Duration("duration", time.Since(v.StartedAt)))
	} else {
		slog.Error(v.Record.Message, err, slog.Duration("duration", time.Since(v.StartedAt)))
	}
}
