package logo

import "log"

// LogLevel represents log levels.
type LogLevel uint8

const (
	// DEBUG represents debug messages.
	DEBUG LogLevel = 0
	// INFO represents info messages.
	INFO LogLevel = 1
	// WARNING represents warning messages.
	WARNING LogLevel = 2
	// ERROR represents error messages.
	ERROR LogLevel = 3
)

// DefaultLogMessageMap maps LogLevel to logging strings in English.
var DefaultLogMessageMap = map[LogLevel]string{DEBUG: "[DEBUG] ", INFO: "[INFO] ", WARNING: "[WARNING] ", ERROR: "[ERROR] "}

// Logo represents the Logo logger.
type Logo struct {
	// Logger is the *log.Logger used for the Logo. The default is log.Printf.
	Logger *log.Logger
	// Level determines the lowest log level to display.
	Level LogLevel
	// LogMessageMap maps LogLevel to logging strings.
	LogMessageMap map[LogLevel]string
}

// New initializes a Logo.
func New(level LogLevel, logger *log.Logger) *Logo {
	return &Logo{
		Logger:        logger,
		Level:         level,
		LogMessageMap: DefaultLogMessageMap,
	}
}

// D logs a debug message.
func (l *Logo) D(format string, v ...interface{}) {
	l.logfL(DEBUG, format, v...)
}

// I logs an info message.
func (l *Logo) I(format string, v ...interface{}) {
	l.logfL(INFO, format, v...)
}

// W logs a warning message.
func (l *Logo) W(format string, v ...interface{}) {
	l.logfL(WARNING, format, v...)
}

// E logs an error message.
func (l *Logo) E(format string, v ...interface{}) {
	l.logfL(ERROR, format, v...)
}

func (l *Logo) logfL(lv LogLevel, format string, v ...interface{}) {
	if lv >= l.Level {
		format = l.LogMessageMap[lv] + format
		l.logf(format, v...)
	}
}

func (l *Logo) logf(format string, v ...interface{}) {
	if l.Logger == nil {
		log.Printf(format, v...)
		return
	}
	l.Logger.Printf(format, v...)
}
