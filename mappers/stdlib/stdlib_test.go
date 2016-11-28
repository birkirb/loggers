package stdlib

import (
	"bytes"
	"log"
	"strings"
	"testing"

	"gopkg.in/birkirb/loggers.v1"
)

func TestLogInterface(t *testing.T) {
	var _ loggers.Contextual = NewDefaultLogger()
	var _ loggers.Advanced = NewDefaultLogger()
	var _ loggers.Standard = &log.Logger{}
}

func TestLogLevelOutput(t *testing.T) {
	l, b := NewBufferedLog()
	l.Info("This is a test")

	expected := "INFO  This is a test\n"
	s := b.String()
	start := strings.Index(s, "INFO")
	actual := s[start:]
	if start < 0 || actual != expected {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expected)
	}
}

func TestLogLevelfOutput(t *testing.T) {
	l, b := NewBufferedLog()
	l.Errorf("This is %s test", "a")

	expected := "ERROR This is a test\n"
	s := b.String()
	start := strings.Index(s, "ERROR")
	actual := s[start:]
	if start < 0 || actual != expected {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expected)
	}
}

func TestLogLevellnOutput(t *testing.T) {
	l, b := NewBufferedLog()
	l.Debugln("This is a test.", "So is this.")

	expected := "DEBUG  This is a test. So is this.\n"
	s := b.String()
	start := strings.LastIndex(s, "DEBUG")
	actual := s[start:]
	if start < 0 || actual != expected {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expected)
	}
}

func TestLogWithFieldsOutput(t *testing.T) {
	l, b := NewBufferedLog()
	l.WithFields("test", true).Warn("This is a message.")

	expected := "WARN  This is a message. [test=true]\n"
	s := b.String()
	start := strings.Index(s, "WARN")
	actual := s[start:]
	if actual != expected {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expected)
	}
}

func TestLogWithFieldsfOutput(t *testing.T) {
	l, b := NewBufferedLog()
	l.WithFields("test", true, "Error", "serious").Errorf("This is a %s.", "message")

	expected := "ERROR This is a message. [test=true, Error=serious]\n"
	s := b.String()
	start := strings.Index(s, "ERROR")
	actual := s[start:]
	if actual != expected {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expected)
	}
}

func TestLogWithFieldsLnOutput(t *testing.T) {
	l, b := NewBufferedLog()
	l.WithFields("test", true, "Error", "not so serious").Warnln("This is your last.")

	expected := "WARN   This is your last. [test=true, Error=not so serious]\n"
	s := b.String()
	start := strings.Index(s, "WARN")
	actual := s[start:]
	if actual != expected {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expected)
	}
}

func NewBufferedLog() (loggers.Contextual, *bytes.Buffer) {
	var b []byte
	var bb = bytes.NewBuffer(b)
	l := log.New(bb, "", log.Ldate|log.Ltime)
	return NewLogger(l), bb
}
