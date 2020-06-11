package logrus

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/marcaudefroy/loggers"
	"github.com/sirupsen/logrus"
)

func TestLogrusInterface(t *testing.T) {
	var _ loggers.Contextual = NewDefaultLogger()
	var _ loggers.Advanced = &logrus.Logger{}
}

func TestLogrusLevelOutput(t *testing.T) {
	l, b := newBufferedLogrusLog()
	l.Info("This is a test")

	expectedMatch := "(?i)info.*This is a test"
	actual := b.String()
	if ok, _ := regexp.Match(expectedMatch, []byte(actual)); !ok {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expectedMatch)
	}
}

func TestLogrusLevelfOutput(t *testing.T) {
	l, b := newBufferedLogrusLog()
	l.Errorf("This is %s test", "a")

	expectedMatch := "(?i)erro.*This is a test"
	actual := b.String()
	if ok, _ := regexp.Match(expectedMatch, []byte(actual)); !ok {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expectedMatch)
	}
}

func TestLogrusLevellnOutput(t *testing.T) {
	l, b := newBufferedLogrusLog()
	l.Debugln("This is a test.", "So is this.")

	expectedMatch := "(?i)debu.*This is a test. So is this."
	actual := b.String()
	if ok, _ := regexp.Match(expectedMatch, []byte(actual)); !ok {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expectedMatch)
	}
}

func TestLogrusWithFieldsOutput(t *testing.T) {
	l, b := newBufferedLogrusLog()
	l.WithFields("test", true).Warn("This is a message.")

	expectedMatch := "(?i)warn.*This is a message.*test.*=true"
	actual := b.String()
	if ok, _ := regexp.Match(expectedMatch, []byte(actual)); !ok {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expectedMatch)
	}
}

func TestLogrusChainedWithFieldsOutput(t *testing.T) {
	l, b := newBufferedLogrusLog()
	l.WithFields("test", true).WithFields("test2", false).Warn("This is a message.")

	expectedMatch := "(?i)warn.*This is a message.*test.*=true.*test2.*=false"
	actual := b.String()
	if ok, _ := regexp.Match(expectedMatch, []byte(actual)); !ok {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expectedMatch)
	}
}

func TestLogrusWithFieldsfOutput(t *testing.T) {
	l, b := newBufferedLogrusLog()
	l.WithFields("test", true, "Error", "serious").Errorf("This is a %s.", "message")

	expectedMatch := "(?i)erro.*This is a message.*Error.*=serious.*test.*=true"
	actual := b.String()
	if ok, _ := regexp.Match(expectedMatch, []byte(actual)); !ok {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expectedMatch)
	}
}

func TestLogrusFieldsfOutput(t *testing.T) {
	l := NewDefaultLogger()
	l = l.WithFields("test", true, "Error", "serious")
	nl := l.WithField("foo", "bar")

	lFields := l.Fields()
	nlFields := nl.Fields()

	if len(lFields) != 4 {
		t.Errorf("Log fields must have %d elements, it have %d", 4, len(lFields))
	}
	if len(nlFields) != 6 {
		t.Errorf("Log fields must have %d elements, it have %d", 4, len(nlFields))
	}
}

func newBufferedLogrusLog() (loggers.Contextual, *bytes.Buffer) {
	var b []byte
	var bb = bytes.NewBuffer(b)

	l := logrus.New()
	l.Out = bb
	l.Level = logrus.DebugLevel
	return NewLogger(l), bb
}
