package arrayList

import (
	"godata/list"
	"strings"
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
	return &ArrayList{}
}

//Append elements at the end of the list
func (list *ArrayList) Add(elements ...interface{}) {
	//是否扩容
	list.grow(len(elements))
	for _, element := range elements {
		list.elements[list.size] = element
		list.size += 1
	}
}

//Retrun the element at the specified index
func (list *ArrayList) Get(index int) (interface{}, bool) {
	if !list.boundCheck(index) {
		return nil, false
	}

	return list.elements[index], true
}

//Remove the element at te specified index
func (list *ArrayList) Remove(index int) {
	if !list.boundCheck(index) {
		return
	}

	list.elements[index] = nil
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size -= 1
	//shirnk if necessary
	list.shrink()
}

//return true if the list contains the specified elements

func (list *ArrayList) Contains(elements ...interface{}) bool {
	var flag bool = true
	for _, e := range elements {
		if !list.contain(e) {
			flag = false
			break
		}
	}

	return flag
}

func (list *ArrayList) String() string {
	str := "ArrayList"
	var values string = ""
	for _, value := range list.elements[:list.Size()] {
		values += strings.Join(value, " , ")
	}
	str += "["
	str += values
	str += "]"
	return str
}

//collection interface method
//Clear the list
func (list *ArrayList) Clear() {
	list.size = 0
	list.elements = []interface{}{}
}

//return the true if the list is empty
func (list *ArrayList) IsEmpty() bool {
	return list.size == 0
}

//return the size of list
func (list *ArrayList) Size() int {
	return list.size
}

//return values of the list
func (list *ArrayList) Values() []interface{} {
	newE := make([]interface{}, list.size, list.size)
	copy(newE, list.elements[:list.size])
	return newE
}
func (list *ArrayList) contain(element interface{}) bool {
	var flag bool = false
	for _, e := range list.elements {
		if e == element {
			flag = true
			break
		}
	}
	return flag
}

//check the index is withing bounds of the list
func (list *ArrayList) boundCheck(index int) bool {
	return index >= 0 && index < list.size
}

//Expand the list if necessary before Add elements
func (list *ArrayList) grow(n int) {
	currentCapacity := cap(list.elements)
	if list.size+n >= currentCapacity {
		//expand by GROUTH_FACTOR
		newCapacity := int(GROUTH_FACTOR * float32(currentCapacity+n))
		list.refresh(newCapacity)
	}
}

//shrink the list if necessary after remove elements
func (list *ArrayList) shrink() {
	if SHRINK_FACTOR == 0.0 {
		return
	}
	currentCapacity := cap(list.elements)
	if list.size <= int(float32(currentCapacity)*SHRINK_FACTOR) {
		//调整容量为实际大小
		list.refresh(list.size)
	}

}

func (list *ArrayList) refresh(cap int) {
	newElements := make([]interface{}, cap, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}
