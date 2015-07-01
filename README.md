# godata (Go Data Structrues)
Golang基础数据结构实现

### 数据结构
* 容器
	* Set
		* HashSet
 
### Set
&nbsp;&nbsp;&nbsp;&nbsp;Set直观上和数学上的集(set)概念是相同的。Set最大的特性就是不允许在其中存放的元素是重复的。set可以用来过滤在其他集合中存放的元素,从而得到一个没有包含重复元素的新集合。

&nbsp;&nbsp;&nbsp;&nbsp;接口 Set包含的方法如下:

```go
type Set interface {
	Add(e interface{}) bool                 //添加元素
	Remove(e interface{})                   //删除元素
	Clear()                                 //清空元素
	Contains(e interface{}) bool            //是否包含元素
	Size()  int                             //集合大小
	Same(other Set)   bool                  //判断集合是否相等
	Elements()  []interface{}               //返回元素列表
	String() string                         //集合字符串表现形式
}

//高级用法
//判断是否为超集
func IsSuperSet(one,other Set) bool {}
//合并集合
func Union(one,other Set) Set {}
```