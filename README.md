# godata (Go Data Structrues)
Golang基础数据结构实现
### 数据结构
- [collection](#Collection)
  - [Set](#Set)
    - [HashSet](#HashSet)

### Collection
&nbsp;&nbsp;集合接口,所有结构都应该实现的接口,提供集合都应该有的方法集合。

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
