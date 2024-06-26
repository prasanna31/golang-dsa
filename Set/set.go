package set

type Set struct {
	data map[interface{}]bool
}

func NewSet() *Set {
	return &Set{
		data: make(map[interface{}]bool),
	}
}

func (s *Set) Add(item interface{}) {
	s.data[item] = true
}

func (s *Set) Remove(item interface{}) {
	delete(s.data, item)
}

func (s *Set) Contains(item interface{}) bool {
	_, exists := s.data[item]
	return exists
}

func (s *Set) Size() int {
	return len(s.data)
}

func (s *Set) Clear() {
	s.data = make(map[interface{}]bool)
}

func (s *Set) Values() []interface{} {
	values := make([]interface{}, 0, len(s.data))
	for item := range s.data {
		values = append(values, item)
	}
	return values
}
