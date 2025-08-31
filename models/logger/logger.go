package logger

import (
	"context"
	"log/slog"
	"os"
	"sync"
)

type Logger struct {
	ctx context.Context
	l   *slog.Logger
}

func (l Logger) WithContext(ctx context.Context) Logger {
	l1 := l
	l1.ctx = ctx
	return l1
}

func (l Logger) Info(msg string, args ...slog.Attr) {
	printLog(l.l.Info, msg, args...)
}

func (l Logger) Error(msg string, args ...slog.Attr) {
	printLog(l.l.Error, msg, args...)
}

func printLog(f func(msg string, attrs ...any), msg string, args ...slog.Attr) {
	attrs := make([]any, len(args))
	for i := range args {
		attrs[i] = args[i]
	}

	f(msg, attrs...)
}

var (
	o sync.Once

	logger Logger
)

func New() Logger {
	o.Do(func() {
		logger = Logger{
			ctx: context.Background(),
			l:   slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		}
	})

	return logger
}
