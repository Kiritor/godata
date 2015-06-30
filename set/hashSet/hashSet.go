package hashSet

import (
	//"godata/set"
)

type HashSet struct {
	m map[interface{}]bool
}


//init HashSetSet
func NewHashSet() *HashSet {
	return &HashSet{m:make(map[interface{}]bool)}
}

//add an element into set
func (set *HashSet) Add(e interface{}) bool {
	if !set.m[e] {
		set.m[e] = true
		return true
	}
	return false
}

//remove an element from set
func (set *HashSet) Remove(e interface{}) {
	delete(set.m,e)
}

//clear a set
func (set *HashSet) Clear() {
	set.m = make(map[interface{}]bool)
}
