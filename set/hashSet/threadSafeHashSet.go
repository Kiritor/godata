/**
 * ThreadHashSet:Thread safe
 * @Author        LCore
 * @Site          http://kiritor.github.io
*/
package hashSet

import (
    "godata/set"
	"fmt"
	"bytes"
	"sync"
)

type ThreadHashSet struct {
	HashSet                      //匿名字段
	mutex sync.Mutex
}

func NewThreadHashSet() *ThreadHashSet{
	var mutex sync.Mutex
	return &ThreadHashSet{HashSet{m:make(map[interface{}]bool)},mutex}
}

func(hashSet *ThreadHashSet) Add(e interface{}) bool {
	hashSet.mutex.Lock()
	defer hashSet.mutex.Unlock()
	return hashSet.HashSet.Add(e)
}

func (hashSet *ThreadHashSet) Remove(e interface{}) {
	hashSet.mutex.Lock()
	defer hashSet.mutex.Unlock()
	hashSet.HashSet.Remove(e)
}

func (hashSet *ThreadHashSet) Clear() {
	hashSet.mutex.Lock()
	defer hashSet.mutex.Unlock()
	hashSet.HashSet.Clear()
}

func (hashSet *ThreadHashSet) Contains(e interface{}) bool{
	hashSet.mutex.Lock()
	defer hashSet.mutex.Unlock()
	return hashSet.HashSet.Contains(e)
}

func (hashSet *ThreadHashSet) Size() int {
	hashSet.mutex.Lock()
	defer hashSet.mutex.Unlock()
	return hashSet.HashSet.Size()
}

func (hashSet *ThreadHashSet) Same(set Set) bool {
	hashSet.mutex.Lock()
	defer hashSet.mutex.Unlock()
	return hashSet.HashSet.Same(set)
}

func (hashSet *ThreadHashSet) Elements() interface{} {
	hashSet.mutex.Lock()
	defer hashSet.mutex.Unlock()
	return hashSet.HashSet.Elements()
}

func (hashSet *ThreadHashSet) String() string {
	hashSet.mutex.Lock()
	defer hashSet.mutex.Unlock()
	return hashSet.HashSet.String()
}
