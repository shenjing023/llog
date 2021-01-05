package log

import (
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

// Level from logrus Level
type Level uint32

const (
	PanicLevel Level = iota

	FatalLevel

	ErrorLevel

	WarnLevel

	InfoLevel

	DebugLevel

	TraceLevel
)

var (
	// std is the name of the standard logger in stdlib `log`
	std    = logrus.New()
	rotate *rotatelogs.RotateLogs
	op     = options{
		formatter: Formatter{},
	}
)

type options struct {
	formatter Formatter
	ops       []rotatelogs.Option
}

type Option func(o *options)

func init() {
	std.SetFormatter(&Formatter{})
}

// SetConsoleLogger set console option
func SetConsoleLogger(opts ...Option) {
	for _, opt := range opts {
		opt(&op)
	}
	std.SetFormatter(&op.formatter)
}

// SetFileLogger set file log option
func SetFileLogger(format string, opts ...Option) (err error) {
	for _, opt := range opts {
		opt(&op)
	}
	op.formatter.DisableColors = true

	rotate, err = rotatelogs.New(format, op.ops...)
	if err != nil {
		return err
	}
	writers := []io.Writer{rotate, os.Stdout}
	twoWriters := io.MultiWriter(writers...)
	std.SetFormatter(&op.formatter)
	std.SetOutput(twoWriters)
	return err
}

// WithLevel set log level
func WithLevel(l Level) Option {
	return func(o *options) {
		std.SetLevel(logrus.Level(l))
	}
}

// WithCaller set logrus.SetReportCaller
func WithCaller(include bool) Option {
	return func(o *options) {
		std.SetReportCaller(include)
	}
}

// WithColor set formatter output color
func WithColor(b bool) Option {
	return func(o *options) {
		o.formatter.DisableColors = b
	}
}

// WithJSON set formatter output color
func WithJSON(b bool) Option {
	return func(o *options) {
		o.formatter.JSONFormat = b
	}
}

// WithHTMLEscape set formatter output color
func WithHTMLEscape(b bool) Option {
	return func(o *options) {
		o.formatter.DisableHTMLEscape = b
	}
}

// WithPrettyPrint set formatter output color
func WithPrettyPrint(b bool) Option {
	return func(o *options) {
		o.formatter.PrettyPrint = b
	}
}

// WithMaxAge 设置文件清理前的最长保存时间
func WithMaxAge(d time.Duration) Option {
	return func(o *options) {
		o.ops = append(o.ops, rotatelogs.WithMaxAge(d))
	}
}

// WithRotationTime 设置日志分割的时间，隔多久分割一次
func WithRotationTime(d time.Duration) Option {
	return func(o *options) {
		o.ops = append(o.ops, rotatelogs.WithRotationTime(d))
	}
}

// SetLevel sets the standard logger level.
// func SetLevel(level logrus.Level) {
// 	std.SetLevel(level)
// }

// SetReportCaller sets whether the standard logger will include the calling
// method as a field.
// func SetReportCaller(include bool) {
// 	std.SetReportCaller(include)
// }

// Trace logs a message at level Trace on the standard logger.
func Trace(args ...interface{}) {
	std.Trace(args...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	std.Debug(args...)
}

// Print logs a message at level Info on the standard logger.
func Print(args ...interface{}) {
	std.Print(args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	std.Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	std.Warn(args...)
}

// Warning logs a message at level Warn on the standard logger.
func Warning(args ...interface{}) {
	std.Warning(args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	std.Error(args...)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	std.Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

// Tracef logs a message at level Trace on the standard logger.
func Tracef(format string, args ...interface{}) {
	std.Tracef(format, args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}

// Printf logs a message at level Info on the standard logger.
func Printf(format string, args ...interface{}) {
	std.Printf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}

// Warningf logs a message at level Warn on the standard logger.
func Warningf(format string, args ...interface{}) {
	std.Warningf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(format string, args ...interface{}) {
	std.Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}

// Traceln logs a message at level Trace on the standard logger.
func Traceln(args ...interface{}) {
	std.Traceln(args...)
}

// Debugln logs a message at level Debug on the standard logger.
func Debugln(args ...interface{}) {
	std.Debugln(args...)
}

// Println logs a message at level Info on the standard logger.
func Println(args ...interface{}) {
	std.Println(args...)
}

// Infoln logs a message at level Info on the standard logger.
func Infoln(args ...interface{}) {
	std.Infoln(args...)
}

// Warnln logs a message at level Warn on the standard logger.
func Warnln(args ...interface{}) {
	std.Warnln(args...)
}

// Warningln logs a message at level Warn on the standard logger.
func Warningln(args ...interface{}) {
	std.Warningln(args...)
}

// Errorln logs a message at level Error on the standard logger.
func Errorln(args ...interface{}) {
	std.Errorln(args...)
}

// Panicln logs a message at level Panic on the standard logger.
func Panicln(args ...interface{}) {
	std.Panicln(args...)
}

// Fatalln logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalln(args ...interface{}) {
	std.Fatalln(args...)
}

// WithField creates an entry from the standard logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithField(key string, value interface{}) *logrus.Entry {
	return std.WithField(key, value)
}

// WithFields creates an entry from the standard logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithFields(fields Fields) *logrus.Entry {
	return std.WithFields(logrus.Fields(fields))
}
