/**
 * list interface
 * @Auth           LCore
 * @Site           http://kiritor.github.io
 */
package list

import (
	"godata/collection"
)

type List interface {
	Get(index int) (interface{}, bool)
	Remove(index int)
	Add(elements ...interface{})
	Contains(elements ...interface{}) bool
	collection.Collection
}
