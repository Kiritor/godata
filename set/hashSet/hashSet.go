package hashSet

import (
	"godata/set"
	"fmt"
	"bytes"
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

/**
   if element is contain in the set ,the element
   can not be function、map、slice,otherwise it can be panic
*/
func (set *HashSet) Contains(e interface{}) bool {
	return set.m[e]
}

//return the size of the set
func (set *HashSet) Size() int {
	return len(set.m)
}

//return a set is eaquals an other set
func (set *HashSet) Same(other set.Set) bool {
	if other == nil {
		return false
	}
	if set.Size() != other.Size() {
		return false
	}

	for key:=range set.m {
		if !other.Contains(key) {
			return false
		}
	}

	return true
}
//iteral the set
func (set *HashSet) Elements() []interface{} {
    initialLen := len(set.m)
	snapshot := make([]interface{}, initialLen)
	actLen:=0
	for key := range set.m {
		snapshot[actLen] = key
		actLen++
	}
	return snapshot
}

//Stirng the set
func (set *HashSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("Set{")
	first:=true
	for key:=range set.m {
		if first {
			first = false
		}else {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf("%v",key))
	}
	buf.WriteString("}")
	return buf.String()
}

