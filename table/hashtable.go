package table

type HashTable struct {
	data map[string]interface{}
}

func NewHashTable() *HashTable {
	return &HashTable{
		data: make(map[string]interface{}),
	}
}

func (ht *HashTable) Put(key string, value interface{}) {
	ht.data[key] = value
}

func (ht *HashTable) Get(key string) (interface{}, bool) {
	value, ok := ht.data[key]
	return value, ok
}

func (ht *HashTable) Delete(key string) {
	delete(ht.data, key)
}
