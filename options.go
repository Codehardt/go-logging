package log

import "io"

// Option is used to configure the logger on initialization
type Option func(*Logger)

// OptionEnableDebug enables/disables debug mode
func OptionEnableDebug(debugOn bool) func(*Logger) {
	return func(l *Logger) {
		l.debug = debugOn
	}
}

// OptionEnableTrace enables/disables trace mode
func OptionEnableTrace(traceOn bool) func(*Logger) {
	return func(l *Logger) {
		l.trace = traceOn
	}
}

type mutex interface {
	Lock()
	Unlock()
}

type fakeMutex struct{}

func (f *fakeMutex) Lock()   {}
func (f *fakeMutex) Unlock() {}

// OptionDisableMutex disables mutex protection of writer
func OptionDisableMutex(mutexOff bool) func(*Logger) {
	return func(l *Logger) {
		if mutexOff {
			l.mu = &fakeMutex{}
		}
	}
}

// OptionWithTimeFormat changes the time format
func OptionWithTimeFormat(format string) func(*Logger) {
	return func(l *Logger) {
		l.timeformat = format
	}
}

// OptionEnableLocalTime enables/disables local time instead of utc time
func OptionEnableLocalTime(localTimeOn bool) func(*Logger) {
	return func(l *Logger) {
		l.localtime = localTimeOn
	}
}

type staticKV struct {
	key   string
	value interface{}
}

// OptionWithStaticKV adds a key value pair to each upcoming log line .. This option can be passed multiple times.
func OptionWithStaticKV(key string, value interface{}) func(*Logger) {
	return func(l *Logger) {
		l.staticKVs = append(l.staticKVs, staticKV{key, value})
	}
}

// OptionWithMessageKey overwrites the key that is used for message
func OptionWithMessageKey(key string) func(l *Logger) {
	if key == "" {
		panic("missing key for message")
	}
	return func(l *Logger) {
		l.messageKey = key
	}
}

// OptionWithTimeKey overwrites the key that is used for time
func OptionWithTimeKey(key string) func(l *Logger) {
	if key == "" {
		panic("missing key for time")
	}
	return func(l *Logger) {
		l.timeKey = key
	}
}

// OptionWithLevelKey overwrites the key that is used for level
func OptionWithLevelKey(key string) func(l *Logger) {
	if key == "" {
		panic("missing key for level")
	}
	return func(l *Logger) {
		l.levelKey = key
	}
}

// OptionWithWriter overwrites the writer of the logger
func OptionWithWriter(w io.Writer) func(l *Logger) {
	if w == nil {
		panic("can not use nil writer for logging")
	}
	return func(l *Logger) {
		l.writer = w
	}
}

// OptionDisableTime removes time from log lines
func OptionDisableTime(timeOff bool) func(l *Logger) {
	return func(l *Logger) {
		l.withoutTime = timeOff
	}
}

// OptionDisableLevel removes level from log lines
func OptionDisableLevel(levelOff bool) func(l *Logger) {
	return func(l *Logger) {
		l.withoutLevel = levelOff
	}
}

func (l *Logger) validateOptions() {
	keys := make(map[string]struct{})
	keys[l.messageKey] = struct{}{}
	keys[l.levelKey] = struct{}{}
	keys[l.timeKey] = struct{}{}
	if len(keys) != 3 {
		panic("key names have to be unique")
	}
}
