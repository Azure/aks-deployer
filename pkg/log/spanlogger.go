package log

import (
	"context"
	"time"
)

type SpanTrace interface {
	Debugf(ctx context.Context, format string, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})
	Warningf(ctx context.Context, format string, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
	Fatalf(ctx context.Context, format string, args ...interface{})

	Debug(ctx context.Context, msg string)
	Info(ctx context.Context, msg string)
	Warning(ctx context.Context, msg string)
	Error(ctx context.Context, msg string)
	ErrorWithStack(ctx context.Context, err error)
	Fatal(ctx context.Context, err error)

	Infow(msg string, keysAndValues ...interface{})
	Warningw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})

	InfofWithLatency(ctx context.Context, latency time.Duration, fmt string, args ...interface{})
}

func (l *Logger) Fatal(ctx context.Context, msg string) {
	l.withSpanInfo(ctx).traceFatal(msg)
}

func (l *Logger) Fatalf(ctx context.Context, format string, args ...interface{}) {
	l.withSpanInfo(ctx).traceFatalf(format, args...)
}

func (l *Logger) ErrorWithStack(ctx context.Context, err error) {
	l.withSpanInfo(ctx).traceErrorWithStack(err)
}

func (l *Logger) Error(ctx context.Context, msg string) {
	l.withSpanInfo(ctx).traceError(msg)
}

func (l *Logger) Errorf(ctx context.Context, format string, args ...interface{}) {
	l.withSpanInfo(ctx).traceErrorf(format, args...)
}

func (l *Logger) Debug(ctx context.Context, msg string) {
	l.withSpanInfo(ctx).traceDebug(msg)
}

func (l *Logger) Debugf(ctx context.Context, format string, args ...interface{}) {
	l.withSpanInfo(ctx).traceDebugf(format, args...)
}

func (l *Logger) Warning(ctx context.Context, msg string) {
	l.withSpanInfo(ctx).traceWarning(msg)
}

func (l *Logger) Warningf(ctx context.Context, format string, args ...interface{}) {
	l.withSpanInfo(ctx).traceWarningf(format, args...)
}

func (l *Logger) Info(ctx context.Context, msg string) {
	l.withSpanInfo(ctx).traceInfo(msg)
}

func (l *Logger) Infof(ctx context.Context, msg string, args ...interface{}) {
	l.withSpanInfo(ctx).traceInfof(msg, args...)
}

func (l *Logger) Infow(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.withSpanInfo(ctx).traceInfow(msg, keysAndValues...)
}

func (l *Logger) Errorw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.withSpanInfo(ctx).traceErrorw(msg, keysAndValues...)
}

func (l *Logger) Warningw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.withSpanInfo(ctx).traceWarningw(msg, keysAndValues...)
}

func (l *Logger) Panicw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.withSpanInfo(ctx).tracePanicw(msg, keysAndValues...)
}

func (l *Logger) InfofWithLatency(ctx context.Context, latency time.Duration, fmt string, args ...interface{}) {
	l.withSpanInfo(ctx).traceInfofWithLatency(latency, fmt, args...)
}
