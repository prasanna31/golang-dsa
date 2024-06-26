package set

type HashSet struct {
	set map[interface{}]struct{}
}

func NewHashSet() *HashSet {
	return &HashSet{
		set: make(map[interface{}]struct{}),
	}
}

func (hs *HashSet) Add(item interface{}) {
	hs.set[item] = struct{}{}
}

func (hs *HashSet) Remove(item interface{}) {
	delete(hs.set, item)
}

func (hs *HashSet) Contains(item interface{}) bool {
	_, exists := hs.set[item]
	return exists
}

func (hs *HashSet) Size() int {
	return len(hs.set)
}

func (hs *HashSet) Clear() {
	hs.set = make(map[interface{}]struct{})
}
