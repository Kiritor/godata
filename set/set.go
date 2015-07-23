/**
 * Set接口
 * @Author        LCore
 * @Site          http://kiritor.github.io
 */

package set

type Set interface {
	Add(e interface{}) bool      //添加元素
	Remove(e interface{})        //删除元素
	Clear()                      //清空元素
	Contains(e interface{}) bool //是否包含元素
	Size() int                   //集合大小
	Same(other Set) bool         //判断集合是否相等
	Elements() []interface{}     //返回元素列表
	String() string              //集合字符串表现形式
}

/**
 * 高级用法,判断某个set是否是另一个set的超集
 */
func IsSuperSet(one, other Set) bool {
	if one == nil || other == nil {
		return false
	}

	oneLen := one.Size()
	otherLen := other.Size()
	if oneLen == 0 || oneLen <= otherLen {
		return false
	}
	if oneLen > 0 && otherLen == 0 {
		return true
	}
	for _, key := range other.Elements() {
		if !one.Contains(key) && key != nil {
			return false
		}
	}

	return true
}

//合并两个集合
func Union(one, other Set) Set {
	if one.Same(other) {
		return one
	}
	for _, key := range other.Elements() {
		one.Add(key)
	}
	return one
}

