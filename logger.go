package log

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// Logger holds all logger configurations
type Logger struct {
	debug      bool
	trace      bool
	mu         mutex
	formatter  Formatter
	timeformat string
	localtime  bool
	staticKVs  []staticKV
	timeKey    string
	messageKey string
	levelKey   string
	writer     io.Writer
}

// New initializes a new logger using options to configure the logger
func New(formatter Formatter, options ...Option) *Logger {
	if formatter == nil {
		panic("missing log formatter")
	}
	l := &Logger{
		mu:         &sync.Mutex{},
		timeformat: time.RFC3339,
		timeKey:    "TIME",
		messageKey: "MESSAGE",
		levelKey:   "LEVEL",
		writer:     os.Stdout,
		formatter:  formatter,
	}
	for _, option := range options {
		if option == nil {
			panic("nil option in logger initialization found")
		}
		option(l)
	}
	l.validateOptions()
	return l
}

func (l *Logger) log(level string, message string, kv ...interface{}) (n int, err error) {
	for _, s := range l.staticKVs {
		kv = append([]interface{}{s.key, s.value}, kv...)
	}
	msg := l.formatter(l, level, message, kv...)
	l.mu.Lock()
	n, err = fmt.Fprintln(l.writer, msg)
	l.mu.Unlock()
	return
}

// Info message
func (l *Logger) Info(message string, kv ...interface{}) (int, error) {
	return l.log("Info", message, kv...)
}

// Notice message
func (l *Logger) Notice(message string, kv ...interface{}) (int, error) {
	return l.log("Notice", message, kv...)
}

// Warning message
func (l *Logger) Warning(message string, kv ...interface{}) (int, error) {
	return l.log("Warning", message, kv...)
}

// Error message
func (l *Logger) Error(message string, kv ...interface{}) (int, error) {
	return l.log("Error", message, kv...)
}

// Fatal message
func (l *Logger) Fatal(message string, kv ...interface{}) (int, error) {
	return l.log("Fatal", message, kv...)
}

// Debug message
func (l *Logger) Debug(message string, kv ...interface{}) (int, error) {
	if !l.debug {
		return 0, nil
	}
	return l.log("Debug", message, kv...)
}

// Trace message
func (l *Logger) Trace(message string, kv ...interface{}) (int, error) {
	if !l.trace {
		return 0, nil
	}
	return l.log("Trace", message, kv...)
}

var logger = New(FormatterSimple)

// SetLogger sets a global logger
func SetLogger(l *Logger) {
	if l == nil {
		panic("can not set nil logger")
	}
	logger = l
}

// Info message with global logger
func Info(message string, kv ...interface{}) (int, error) {
	return logger.Info(message, kv...)
}

// Notice message with global logger
func Notice(message string, kv ...interface{}) (int, error) {
	return logger.Notice(message, kv...)
}

// Warning message with global logger
func Warning(message string, kv ...interface{}) (int, error) {
	return logger.Warning(message, kv...)
}

// Error message with global logger
func Error(message string, kv ...interface{}) (int, error) {
	return logger.Error(message, kv...)
}

// Fatal message with global logger
func Fatal(message string, kv ...interface{}) (int, error) {
	return logger.Fatal(message, kv...)
}

// Debug message with global logger
func Debug(message string, kv ...interface{}) (int, error) {
	return logger.Debug(message, kv...)
}

// Trace message with global logger
func Trace(message string, kv ...interface{}) (int, error) {
	return logger.Trace(message, kv...)
}
