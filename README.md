# godata (Go Data Structrues)
Golang基础数据结构实现
### 数据结构
- [collection](#collection)
  - [List](#list)
    - [ArrayList](#arrayList)

### Collection
&nbsp;&nbsp;&nbsp;&nbsp;集合接口,所有结构都应该实现的接口,提供集合都应该有的方法集合。

Collection接口包含的方法如下:
```go
package collection

type Collection interface {
	//return the number of elements in this collection
	Size() int
	//return true if this collection contains no elements
	IsEmpty() bool
	//clear the collection
	Clear()
	//return slice of the collection's element
	Values() []interface{}
}
```
### List	
&nbsp;&nbsp;&nbsp;&nbsp;List(列表),主要的特性是元素以线性的方式存储,有序,且集合中可以存放重复的对象.

&nbsp;&nbsp;&nbsp;&nbsp;接口List继承collection接口,包含的方法如下:
```go
	Get(index int) (interface{}, bool)
	Remove(index int)
	Add(elements ...interface{})
	Contains(elements ...interface{}) bool
	collection.Collection
```