package loggers

import (
	"log"
	"testing"
)

func TestInterface(t *testing.T) {
	var _ Standard = &log.Logger{}
}
