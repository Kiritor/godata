package hashSet


import (
	"fmt"
	"testing"
)


func TestHashSetAdd(t *testing.T) {

    set1 :=NewHashSet()
	set1.Add(1)
	fmt.Println(set1)
}

func TestHashSetRemove(t *testing.T) {
	set1 :=NewHashSet()
	set1.Add(1)
	fmt.Println(set1.String())
	set1.Remove(1)
	fmt.Println(set1.String())
}
