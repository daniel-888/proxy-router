// log package helpers to enforce consistent log entries
// uses logrus
// greatly inspired by the infamous Scott Millner
package log

import (
	"io"
	"os"
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

type Fields map[string]interface{}

type Logger struct {
	logger *logrus.Logger
}

var defaultOutput = os.Stdout

// New returns a new logger.
// Default output is os.Stdout.
func New() *Logger {
	l := logrus.New()
	l.SetOutput(defaultOutput)

	return &Logger{
		l,
	}
}

// SetOutput sets the output for the logger.
func (l *Logger) SetOutput(output io.Writer) *Logger {
	l.logger.SetOutput(output)

	return l
}

// SetFormatJSON sets the logger to JSON format.
func (l *Logger) SetFormatJSON() *Logger {
	l.logger.SetFormatter(&logrus.JSONFormatter{})

	return l
}

// Tracef logs a trace event.
func (l *Logger) Tracef(format string, args ...interface{}) {
	l.logger.Tracef(format, args...)
}

// TraceJSON logs a trace event with fields, in JSON format.
func (l *Logger) TraceJSON(fields Fields, args ...interface{}) {
	l.withFields(fields).Trace(args...)
}

// Debugf logs a debug event.
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

// DebugJSON logs a debug event with fields, in JSON format.
func (l *Logger) DebugJSON(fields Fields, args ...interface{}) {
	l.withFields(fields).Debug(args...)
}

// Infof logs an info event.
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

// InfoJSON logs an info event with fields, in JSON format.
func (l *Logger) InfoJSON(fields Fields, args ...interface{}) {
	l.withFields(fields).Info(args...)
}

// Warnf logs a warning event.
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

// WarnJSON logs a warning event with fields, in JSON format.
func (l *Logger) WarnJSON(fields Fields, args ...interface{}) {
	l.withFields(fields).Warn(args...)
}

// Errorf logs an error event.
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

// ErrorJSON logs an error event with fields, in JSON format.
func (l *Logger) ErrorJSON(fields Fields, args ...interface{}) {
	l.withFields(fields).Error(args...)
}

// Fatalf logs a fatal event, calls os.Exit(1) aftering logging.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

// FatalJSON logs a fatal event with fields, in JSON format.
// Calls os.Exit(1) aftering logging.
func (l *Logger) FatalJSON(fields Fields, args ...interface{}) {
	l.withFields(fields).Fatal(args...)
}

// Panicf logs a panic event, calls panic() after logging.
func (l *Logger) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}

// PanicJSON logs a panic event with fields, in JSON format.
// Calls panic() after logging.
func (l *Logger) PanicJSON(fields Fields, args ...interface{}) {
	l.withFields(fields).Panic(args...)
}

// Close will close the logger's writer.
func (l *Logger) Close() error {
	return l.logger.Writer().Close()
}

func (l *Logger) withFields(fields Fields) *logrus.Entry {
	fields["stacktrace"] = string(debug.Stack())

	return l.logger.WithFields(logrus.Fields(fields))
}
