package slogd

import (
	"time"

	"golang.org/x/exp/slog"
)

type slogd struct {
	message   string
	startedAt time.Time
	args      []any
}

func New(msg string, args ...any) slogd {
	var v slogd
	v.message = msg
	v.startedAt = time.Now()
	v.args = args
	return v
}

func (v slogd) Stop(err *error) {
	v.args = append(v.args, slog.Duration("duration", time.Duration(time.Since(v.startedAt))))
	if err == nil || *err == nil {
		slog.Info(v.message, v.args...)
	} else {
		slog.Error(v.message, *err, v.args...)
	}
}
