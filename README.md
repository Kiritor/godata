#godata (Go Data Structrues)
Golang基础数据结构实现

##数据结构
- [collection](#collection)
  - [List](#list)
    - [ArrayList](#arrayList)

###Collection

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
####List	

&nbsp;&nbsp;&nbsp;&nbsp;List(列表),主要的特性是元素以线性的方式存储,没有一个排序规则,且集合中可以存放重复的对象.通过索引访问和删除元素.
&nbsp;&nbsp;&nbsp;&nbsp;接口List继承collection接口,包含的方法如下:
```go
	Get(index int) (interface{}, bool)
	Remove(index int)
	Add(elements ...interface{})
	Contains(elements ...interface{}) bool
	collection.Collection
```
#####ArrayList
&nbsp;&nbsp;&nbsp;&nbsp;ArrayList实现List接口,ArrayList是基于动态数组的数据结构,会自动进行扩容(100%扩容)和压缩操作(当实际大小为容器的1/4时收缩为实际大小),ArrayList的Get、Remove方法都是O(n)时间复杂度,Contains方法为O(n*n)复杂度.

```go
package main

import (
	"github.com/kiritor/godata/list/arrayList"
	"fmt"
)

func main(){
	list := arrayList.New()
	fmt.Println(list.String())             //ArrayList[]
	list.Add(`A`,'B',"CDE")        
	fmt.Println(list.String())             //ArrayList[A,66,CDE]
	fmt.Println(list.Get(0))               //A true
    list.Remove(1)                       
	fmt.Println(list.String())             //ArrayList[A,CDE]]
	fmt.Println(list.Values())             //[A CDE
	fmt.Println(list.Contains("CDE"))      //true
	list.Clear()
	fmt.Println(list.String())             //ArrayList[]
	fmt.Println(list.Size())               //0
}

```
