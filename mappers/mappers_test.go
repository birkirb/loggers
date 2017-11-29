package mappers

import (
	"testing"

	"github.com/marcaudefroy/loggers"
)

func TestInterface(t *testing.T) {
	var _ LevelMapper = &standardMap{}
	var _ loggers.Standard = &standardMap{}

	var _ LevelMapper = &AdvancedMap{}
	var _ loggers.Advanced = &AdvancedMap{}

	var _ LevelMapper = &ContextualMap{}
	var _ loggers.Contextual = &ContextualMap{}
}
