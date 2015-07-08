/**
 * ThreadHashSet:Thread safe
 * @Author        LCore
 * @Site          http://kiritor.github.io
*/
package hashSet

import (

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
