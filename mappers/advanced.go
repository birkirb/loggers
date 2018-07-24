package mappers

// AdvancedMap maps a standard logger to an advanced logger interface.
type AdvancedMap struct {
	standardMap
	t TestHelper
}

// NewAdvancedMap returns an advanced logger that is mapped via mapper.
func NewAdvancedMap(m LevelMapper) *AdvancedMap {
	var a AdvancedMap

	if m != nil {
		a.LevelMapper = m
	}

	return &a
}

// TestHelper describes a *testing.T helper.
type TestHelper interface {
	Helper()
}

// NewAdvancedMapTesting returns an advanced logger that is mapped via mapper.
// A TestHelper can be passed. This will be called if not nil.
func NewAdvancedMapTesting(m LevelMapper, t TestHelper) *AdvancedMap {
	var a AdvancedMap

	if m != nil {
		a.LevelMapper = m
	}
	a.t = t
	return &a
}

// Debug should be used when logging exessive debug info.
func (a *AdvancedMap) Debug(v ...interface{}) {
	if a.t != nil {
		a.t.Helper()
	}
	a.LevelPrint(LevelDebug, v...)
}

// Debugf works the same as Debug but supports formatting.
func (a *AdvancedMap) Debugf(format string, v ...interface{}) {
	if a.t != nil {
		a.t.Helper()
	}
	a.LevelPrintf(LevelDebug, format, v...)
}

// Debugln works the same as Debug but supports formatting.
func (a *AdvancedMap) Debugln(v ...interface{}) {
	if a.t != nil {
		a.t.Helper()
	}
	a.LevelPrintln(LevelDebug, v...)
}

// Info is a general function to log something.
func (a *AdvancedMap) Info(v ...interface{}) {
	if a.t != nil {
		a.t.Helper()
	}
	a.LevelPrint(LevelInfo, v...)
}

// Infof works the same as Info but supports formatting.
func (a *AdvancedMap) Infof(format string, v ...interface{}) {
	if a.t != nil {
		a.t.Helper()
	}
	a.LevelPrintf(LevelInfo, format, v...)
}

// Infoln works the same as Info but supports formatting.
func (a *AdvancedMap) Infoln(v ...interface{}) {
	if a.t != nil {
		a.t.Helper()
	}
	a.LevelPrintln(LevelInfo, v...)
}

// Warn is useful for alerting about something wrong.
func (a *AdvancedMap) Warn(v ...interface{}) {
	if a.t != nil {
		a.t.Helper()
	}
	a.LevelPrint(LevelWarn, v...)
}

// Warnf works the same as Warn but supports formatting.
func (a *AdvancedMap) Warnf(format string, v ...interface{}) {
	if a.t != nil {
		a.t.Helper()
	}
	a.LevelPrintf(LevelWarn, format, v...)
}

// Warnln works the same as Warn but supports formatting.
func (a *AdvancedMap) Warnln(v ...interface{}) {
	if a.t != nil {
		a.t.Helper()
	}
	a.LevelPrintln(LevelWarn, v...)
}

// Error should be used only if real error occures.
func (a *AdvancedMap) Error(v ...interface{}) {
	if a.t != nil {
		a.t.Helper()
	}
	a.LevelPrint(LevelError, v...)
}

// Errorf works the same as Error but supports formatting.
func (a *AdvancedMap) Errorf(format string, v ...interface{}) {
	if a.t != nil {
		a.t.Helper()
	}
	a.LevelPrintf(LevelError, format, v...)
}

// Errorln works the same as Error but supports formatting.
func (a *AdvancedMap) Errorln(v ...interface{}) {
	if a.t != nil {
		a.t.Helper()
	}
	a.LevelPrintln(LevelError, v...)
}
