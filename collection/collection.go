/**
 * collection interface
 * @Auth           LCore
 * @Site           http://kiritor.github.io
 */
package collection

import (
	"godata/iterable"
)

type Collection interface {
	//return the number of elements in this collection
	Size() int
	//return true if this collection contains no elements
	IsEmpty() bool
	//return true if this collection contains the specified element
	Contains(e interface{})
	//return an arrayble containing all of the elements in this collection
	ToArray() []interface{}
	//return true if the specified element is insert into this collection
	Add(e interface{}) bool
	//return true if the specified element is remove from this collection
	Remove(e interface{}) bool
	//return true if the collection contains the specified collection
	ContainsAll(c Collection) bool
	//return true if the specified collection's elements arerable add to the collection
	AddAll(c Collection) bool
	//return true if the specified collection's elements arerable remove from the collection
	RemoveAll(c Collection) bool
	//clear the collection
	Clear()
	//return true if the specified element equals the collection
	Equals(e interface{}) bool
	//return hashCode
	HashCode() int


	//extend Iterable interface
	iterable.Iterable
}
