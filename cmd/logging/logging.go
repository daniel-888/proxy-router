// Uses logrus framework and serves as a wrapper around the standard log package to enforce consistent log entries
package logging

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"

	"github.com/sirupsen/logrus"

	debugger "runtime/debug"
)

type Config struct {
	Level int `json:"level"`
	FilePath string `json:"filePath"`
	TestFilePath string `json:"testFilePath"`
}

func loadConfiguration(file string) (Config, error) {
	var config Config;
	configFile, err := os.Open(file)
	if err != nil {
		logrus.WithFields(logrus.Fields{stacktraceField: getStacktrace(), id: trace.id}).Errorf(errorM.message, "Could not load logger config file")
		return config, err
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, err
}

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*logrus.Logger
}

// Create single instance of logger with once.Do to use across app
var once sync.Once
var standardLogger *StandardLogger
var logFile *os.File

// Initialize and configure standard logger
// TODO: set test mode using environment variable and remove parameter
func Init(isTestMode bool) {
	once.Do(func() {
		// Change working directory to logging directory to be consistent with tests
		os.Chdir("./logging")
		// Register Cleanup() to fire when fatal level message is logged
		logrus.RegisterExitHandler(Cleanup)
		var baseLogger = logrus.New()
		standardLogger = &StandardLogger{baseLogger}
		standardLogger.Formatter = &logrus.JSONFormatter{}
		// Load config, set log level and file path
		var logLevel int;
		var filePath string;
		config, err := loadConfiguration("config.json")
		if err != nil {
			logLevel = int(logrus.InfoLevel);
			filePath = "/var/log/lumerin.log"
		} else {
			filePath = config.FilePath
			logLevel = config.Level
		}
		if isTestMode {
			filePath = config.TestFilePath
		}
		standardLogger.SetLevel(logrus.Level(logLevel))
		// Fargate container automatically pipes stdout to Cloudwatch log group
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
		logFile = file
		if err == nil {
			// Write to stdout and file
			mw := io.MultiWriter(os.Stdout, file)
			standardLogger.SetOutput(mw)
		} else {
			standardLogger.SetOutput(os.Stdout)
		}
	})
}

// Public getter for logger instance
func GetLogger() (*StandardLogger, error) {
	if standardLogger == nil {
		return standardLogger, errors.New("StandardLogger not initialized")
	}
	return standardLogger, nil
}

// Cleanup resources: called when main routine shutting down or fatal error
func Cleanup() {
	closeLogFile()
}

func closeLogFile() {
	logFile.Close()
}

// Logging Example:
// standardLogger.PanicEvent("Adding Default Dest Failed:", errors.New("kaboom"))
// produces the log entry below
// {"level":"panic", "msg":"Adding Default Dest Failed: - [kaboom]", "stacktrace": ..., "time": ...}

// Stores messages to log
type LogMessage struct {
	id      int
	message string
}

// LogEvent enum: add events as needed
type LogEvent int
const (
	Trace LogEvent = 0
	Debug LogEvent = 1
	Info LogEvent = 2
	Warn LogEvent = 3
	Error LogEvent = 4
	Fatal LogEvent = 5
	Panic LogEvent = 6
)

// Default format only handles two arguments: add new formats as needed
const defaultLogFormat = "%s - %s"
// Declare log messages to standardize log entries: add new messages as needed
var (
	trace = LogMessage{int(Trace), defaultLogFormat}
	debug = LogMessage{int(Debug), defaultLogFormat}
	info = LogMessage{int(Info), defaultLogFormat}
	warn = LogMessage{int(Warn), defaultLogFormat}
	// name errorM to avoid clash with error
	errorM = LogMessage{int(Error), defaultLogFormat}
	fatal = LogMessage{int(Fatal), defaultLogFormat}
	panic = LogMessage{int(Panic), defaultLogFormat}
  )

// Additional log entry fields
const (
	id = "eventId"
	stacktraceField = "stacktrace"
)

func getStacktrace() string {
	return string(debugger.Stack())
}

func (l *StandardLogger) getLoggerEntryWithFields() *logrus.Entry {
	return l.WithFields(logrus.Fields{stacktraceField: getStacktrace(), id: trace.id})
}

// Receiver functions correspond to a LogMessage declared above
func (l *StandardLogger) TraceEvent(msg string, args ...interface{}) {
	l.getLoggerEntryWithFields().Tracef(trace.message, msg, args)
}

func (l*StandardLogger) DebugEvent(msg string, args ...interface{}) {
	l.getLoggerEntryWithFields().Debugf(debug.message, msg, args)
}

func (l*StandardLogger) InfoEvent(msg string, args ...interface{}) {
	l.getLoggerEntryWithFields().Infof(info.message, msg, args)
}

func (l*StandardLogger) WarnEvent(msg string, args ...interface{}) {
	l.getLoggerEntryWithFields().Warnf(warn.message, msg, args)
}

func (l*StandardLogger) ErrorEvent(msg string, args ...interface{}) {
	l.getLoggerEntryWithFields().Errorf(errorM.message, msg, args)
}

func (l *StandardLogger) FatalEvent(msg string, args ...interface{}) {
	l.getLoggerEntryWithFields().Errorf(fatal.message, msg, args)
	l.Fatal()
}

func (l *StandardLogger) PanicEvent(msg string, args ...interface{}) {
	l.getLoggerEntryWithFields().Panicf(panic.message, msg, args)
	l.Panic()
}