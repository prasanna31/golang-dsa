package Map

type HashMap map[string]interface{}

func NewHashMap() HashMap {
	return make(HashMap)
}

func (m HashMap) Put(key string, value interface{}) {
	m[key] = value
}

func (m HashMap) Get(key string) (interface{}, bool) {
	value, ok := m[key]
	return value, ok
}

func (m HashMap) Delete(key string) {
	delete(m, key)
}
