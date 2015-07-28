package arrayList

import (
	"godata/list"
)

/**
 *  编译时检测ArrayList是否实现了list接口
 */
var (
	_ list.List = (*ArrayList)(nil)
)

type ArrayList struct {
	elements []interface{}
	size     int
}

//定义扩容因子和收缩因子

const (
	GROUTH_FACTOR = float32(2.0)  //2倍增长
	SHRINK_FACTOR = float32(0.25) //当大小为容量的1/4时收缩
)


//new a empty arrayList
func New() *ArrayList {
	return &ArayList{}
}

//Append elements at the end of the list
func (list *ArrayList) Add(elements ...interface{}) {

}




//Expand the list if necessary before Add elements
func (list *ArrayList) grow(n int) {
	currentCapacity :=cap(list.elements)
	if list.size+n >= currentCapacity {
		//expand by GROUTH_FACTOR
		newCapacity :=int(GROUTH_FACTOR*float32(currentCapacity+n))
		list.refresh(newCapacity)
	}
}

func (list *ArrayList) refresh(cap int) {
	newElements :=make([]interface{},cap,cap)
	copy(newElements,list.elements)
	list.elements = newElements
}
