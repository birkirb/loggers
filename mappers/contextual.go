package mappers

import "gopkg.in/birkirb/loggers.v1"

// ContextualMap maps a logger to a contextual logger interface.
type ContextualMap struct {
	AdvancedMap
	ContextualMapper
}

// NewContextualMap returns an contextual logger that is mapped via mapper.
func NewContextualMap(m ContextualMapper) *ContextualMap {
	var a ContextualMap

	if m != nil {
		if am := NewAdvancedMap(m); am != nil {
			a.AdvancedMap = *am
		}
		a.ContextualMapper = m
	}

	return &a
}

// NewContextualMapTesting returns an contextual logger that is mapped via mapper.
// A TestHelper can be passed. This will be called if not nil.
func NewContextualMapTesting(m ContextualMapper, t TestHelper) *ContextualMap {
	var a ContextualMap
	a.t = t

	if m != nil {
		if am := NewAdvancedMapTesting(m, t); am != nil {
			a.AdvancedMap = *am
		}
		a.ContextualMapper = m
	}

	return &a
}


// WithField directly maps the loggers method.
func (c *ContextualMap) WithField(key string, value interface{}) loggers.Advanced {
	return c.ContextualMapper.WithField(key, value)
}

// WithFields directly maps the loggers method.
func (c *ContextualMap) WithFields(fields ...interface{}) loggers.Advanced {
	return c.ContextualMapper.WithFields(fields...)
}
