/**
 * Set接口
 * @Author        LCore
 * @Site          http://kiritor.github.io
*/

package set

type Set interface {
	Add(e interface{}) bool                 //添加元素
	Remove(e interface{})                   //删除元素
	Clear()                                 //清空元素
	Contains(e interface{}) bool            //是否包含元素
	Size()  int                             //集合大小
	Same(other Set)   bool                  //判断集合是否相等
	Elements()  []interface{}                 //返回元素列表
	String() string                         //集合字符串表现形式
}
