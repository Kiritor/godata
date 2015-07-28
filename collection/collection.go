/**
 * collection interface
 * @Auth           LCore
 * @Site           http://kiritor.github.io
 */
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
