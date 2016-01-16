package mappers

import (
	"testing"

	"gopkg.in/birkirb/loggers.v1"
)

func TestInterface(t *testing.T) {
	var _ LevelMapper = &standardMap{}
	var _ loggers.Standard = &standardMap{}

	var _ LevelMapper = &AdvancedMap{}
	var _ loggers.Advanced = &AdvancedMap{}

	var _ LevelMapper = &ContextualMap{}
	var _ loggers.Contextual = &ContextualMap{}
}
