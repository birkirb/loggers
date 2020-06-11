package logrus

import (
	"fmt"

	"github.com/marcaudefroy/loggers"
	"github.com/sirupsen/logrus"
)

// Logger is an Contextual logger wrapper over Logrus's logger.
type Logger struct {
	*logrus.Entry
}

// NewLogger returns a Contextual Logger for Logrus's logger.
// Note that any initialization must be done on the input logrus.
func NewLogger(log *logrus.Logger) loggers.Contextual {
	var l Logger
	l.Entry = logrus.NewEntry(log)
	l.Info("Now using Logrus logger package (via loggers/mappers/logrus).")
	return &l
}

// NewDefaultLogger returns a Contextual Logger for Logrus's logger.
// The logger will contain whatever defaults Logrus uses.
func NewDefaultLogger() loggers.Contextual {
	var l Logger
	l.Entry = logrus.NewEntry(logrus.New())
	l.Info("Now using Logrus logger package (via loggers/mappers/logrus).")
	return &l
}

// WithField returns an advanced logger with a pre-set field.
func (l *Logger) WithField(key string, value interface{}) loggers.Contextual {
	var nl Logger
	nl.Entry = l.Entry.WithField(key, value)
	return &nl
}

// WithFields returns an advanced logger with pre-set fields.
func (l *Logger) WithFields(fields ...interface{}) loggers.Contextual {
	var nl Logger
	nl.Entry = l.Entry.WithFields(sliceToMap(fields...))
	return &nl
}

func (l *Logger) Fields() []interface{} {
	return fieldsToSlice(l.Entry.Data)
}

// type LoggerEntry struct {
// 	logger *Logger
// 	*logrus.Entry
// }

// // WithField returns an advanced logger with a pre-set field.
// func (l *LoggerEntry) WithField(key string, value interface{}) loggers.Contextual {
// 	l.Entry = l.Entry.WithField(key, value)
// 	return &l
// }

// // WithFields returns an advanced logger with a pre-set field.
// func (l *LoggerEntry) WithFields(fields ...interface{}) loggers.Contextual {

// 	l.Entry = l.Entry.WithFields(sliceToMap(fields...))
// 	return l
// }

func fieldsToSlice(fields logrus.Fields) []interface{} {
	res := []interface{}{}
	for key, value := range fields {
		res = append(res, key)
		res = append(res, value)
	}
	return res
}

func sliceToMap(fields ...interface{}) map[string]interface{} {
	f := make(map[string]interface{}, len(fields)/2)
	var key, value interface{}
	for i := 0; i+1 < len(fields); i = i + 2 {
		key = fields[i]
		value = fields[i+1]
		if s, ok := key.(string); ok {
			f[s] = value
		} else if s, ok := key.(fmt.Stringer); ok {
			f[s.String()] = value
		}
	}
	return f
}
