package log

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Formatter defines the log format for a single log line
type Formatter func(*Logger, string, string, ...interface{}) string

var now = time.Now

// FormatterJSON formats every log line as JSON format
func FormatterJSON(l *Logger, level string, message string, kv ...interface{}) string {
	m := make(map[string]interface{})
	m[l.messageKey] = message
	if !l.withoutLevel {
		m[l.levelKey] = level
	}
	if !l.withoutTime {
		t := now().Local()
		if !l.localtime {
			t = t.UTC()
		}
		m[l.timeKey] = t.Format(l.timeformat)
	}
	for i := 0; i < len(kv)-1; i += 2 {
		if key, ok := kv[i].(string); ok {
			m[key] = l.adjustValue(kv[i+1])
		} else {
			m[fmt.Sprintf("_invalid_string_%d", i/2)] = l.adjustValue(kv[i+1])
		}
	}
	b, err := json.Marshal(m)
	if err != nil {
		return "JSON ERROR: " + strings.Replace(err.Error(), "\n", " ", -1) + "\n"
	}
	return string(b)
}

var kvKeyReplacer = strings.NewReplacer(":", "", "\n", "", "\"", "")

// FormatterKV formats every log line with key: "value" format using strconv.Quote(...)
func FormatterKV(l *Logger, level string, message string, kv ...interface{}) string {
	var res string
	if !l.withoutTime {
		t := now().Local()
		if !l.localtime {
			t = t.UTC()
		}
		res += l.timeKey + ": " + strconv.Quote(t.Format(l.timeformat))
	}
	if !l.withoutLevel {
		if res != "" {
			res += " "
		}
		res += l.levelKey + ": " + strconv.Quote(level)
	}
	if res != "" {
		res += " "
	}
	res += l.messageKey + ": " + strconv.Quote(message)
	for i := 0; i < len(kv)-1; i += 2 {
		res += " "
		if key, ok := kv[i].(string); ok {
			res += kvKeyReplacer.Replace(key)
		} else {
			res += fmt.Sprintf("_invalid_string_%d", i/2)
		}
		res += ": " + strconv.Quote(fmt.Sprint(l.adjustValue(kv[i+1])))
	}
	return res
}

var simpleValueReplacer = strings.NewReplacer("\"", "", "\\\"", "\"")

// FormatterSimple formats every log line in a simple format %timestamp% [%level%] %message% %key-values%
func FormatterSimple(l *Logger, level string, message string, kv ...interface{}) string {
	var res string
	if !l.withoutTime {
		t := now().Local()
		if !l.localtime {
			t = t.UTC()
		}
		res += t.Format(l.timeformat)
	}
	if !l.withoutLevel {
		if res != "" {
			res += " "
		}
		res += "[" + strings.ToUpper(level[:3]) + "]"
	}
	if res != "" {
		res += " "
	}
	res += simpleValueReplacer.Replace(strconv.Quote(message))
	for i := 0; i < len(kv)-1; i += 2 {
		res += " "
		if key, ok := kv[i].(string); ok {
			res += strings.ToUpper(kvKeyReplacer.Replace(key))
		} else {
			res += fmt.Sprintf("_INVALID_STRING_%d", i/2)
		}
		res += ": " + simpleValueReplacer.Replace(strconv.Quote(fmt.Sprint(l.adjustValue(kv[i+1]))))
	}
	return res
}

func (l *Logger) adjustValue(value interface{}) interface{} {
	if value == nil {
		return nil
	}
	if t, ok := value.(time.Time); ok {
		if l.localtime {
			t = t.Local()
		} else {
			t = t.UTC()
		}
		return t.Format(l.timeformat)
	}
	if t, ok := value.(*time.Time); ok {
		if l.localtime {
			tt := t.Local()
			t = &tt
		} else {
			tt := t.UTC()
			t = &tt
		}
		return t.Format(l.timeformat)
	}
	if err, ok := value.(error); ok {
		return err.Error()
	}
	return value
}
