/**
 * list interface
 * @Auth           LCore
 * @Site           http://kiritor.github.io
 */
package list

type List interface {
	Get(index int) (interface{}, bool)
	Remove(index int)
	Add(elements ...interface{}) bool
	Contains(elements ...interface{}) bool
}
