package Map

// Map :
type Map struct {
	// Key :
	Key any
	// Value :
	Value any
}

// NewMap :
func NewMap() (m *Map, err error) {
	m = new(Map)
	return m, nil
}

// SetKey :
func (m *Map) SetKey(key any) {
	m.Key = key
}

// SetValue :
func (m *Map) SetValue(value any) {
	m.Value = value
}

// GetKey :
func (m *Map) GetKey() any {
	return m.Key
}

// GetValue :
func (m *Map) GetValue() any {
	return m.Value
}

// GetMap :
func (m *Map) GetMap() map[any]any {
	return map[any]any{
		m.Key: m.Value,
	}
}
