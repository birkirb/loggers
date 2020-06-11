package stdlib

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/marcaudefroy/loggers"
	"github.com/marcaudefroy/loggers/mappers"
)

// goLog maps the standard log package logger to an Contextual log interface.
// However it mostly ignores any level info.
type goLog struct {
	logger *log.Logger
	fields []interface{}
}

// NewDefaultLogger returns a Contextual logger using a log.Logger with stderr output.
func NewDefaultLogger() loggers.Contextual {
	var g goLog
	g.logger = log.New(os.Stderr, "", log.Ldate|log.Ltime)
	g.fields = []interface{}{}

	a := mappers.NewContextualMap(&g)
	a.Debug("Now using Go's stdlib log package (via loggers/mappers/stdlib).")

	return a
}

// NewLogger creates a Contextual logger from a log.Logger.
func NewLogger(l *log.Logger) loggers.Contextual {
	var g goLog
	g.logger = l
	g.fields = []interface{}{}
	a := mappers.NewContextualMap(&g)
	a.Debug("Now using Go's stdlib log package (via loggers/mappers/stdlib).")

	return a
}

// LevelPrint is a Mapper method
func (l *goLog) LevelPrint(lev mappers.Level, i ...interface{}) {
	v := []interface{}{lev}
	v = append(v, i...)
	l.logger.Print(v...)
}

// LevelPrintf is a Mapper method
func (l *goLog) LevelPrintf(lev mappers.Level, format string, i ...interface{}) {
	f := "%s" + format
	v := []interface{}{lev}
	v = append(v, i...)
	l.logger.Printf(f, v...)
}

// LevelPrintln is a Mapper method
func (l *goLog) LevelPrintln(lev mappers.Level, i ...interface{}) {
	v := []interface{}{lev}
	v = append(v, i...)
	l.logger.Println(v...)
}

// WithField returns an Contextual logger with a pre-set field.
func (l *goLog) WithField(key string, value interface{}) loggers.Contextual {
	return l.WithFields(key, value)
}

// WithFields returns an Contextual logger with pre-set fields.
func (l *goLog) WithFields(fields ...interface{}) loggers.Contextual {
	if l.fields == nil {
		l.fields = []interface{}{}
	}
	l.fields = append(l.fields, fields...)

	r := gologPostfixLogger{l}
	return mappers.NewContextualMap(&r)
}

func (l *goLog) Fields() []interface{} {
	return l.fields
}

type gologPostfixLogger struct {
	*goLog
}

func (r *gologPostfixLogger) Fields() []interface{} {
	return r.fields
}

func (r *gologPostfixLogger) postfixFromFields() string {
	if len(r.fields) > 1 {
		s := make([]string, 0, len(r.fields)/2)
		for i := 0; i+1 < len(r.fields); i = i + 2 {
			key := r.fields[i]
			value := r.fields[i+1]
			s = append(s, fmt.Sprint(key, "=", value))
		}
		return "[" + strings.Join(s, ", ") + "]"
	}
	return ""
}

func (r *gologPostfixLogger) LevelPrint(lev mappers.Level, i ...interface{}) {
	i = append(i, " ", r.postfixFromFields())

	r.goLog.LevelPrint(lev, i...)
}

func (r *gologPostfixLogger) LevelPrintf(lev mappers.Level, format string, i ...interface{}) {
	format = format + " %s"
	i = append(i, r.postfixFromFields())

	r.goLog.LevelPrintf(lev, format, i...)
}

func (r *gologPostfixLogger) LevelPrintln(lev mappers.Level, i ...interface{}) {
	i = append(i, r.postfixFromFields())
	r.goLog.LevelPrintln(lev, i...)
}
