package container

type Container interface {
	Clear()                //清空元素
	Size() int             //集合大小
	Empty() bool           //是否为空
	Values() []interface{} //元素数组
}
