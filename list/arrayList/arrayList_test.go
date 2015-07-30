package arrayList

import (
	"testing"
)

func TestArrayList(t *testing.T) {
	list := New()
	list.Add(1, 2, 3, 4, 'A')
	if !list.Contains('A') {
		t.Errorf("Not found A,the Contains method is bug!")
	}
	if !list.Contains(2, 4, 3, 1, 'A') {
		t.Errorf("Not found the specified elements,the Contains method is bug!")
	}
	if list.Size() != 5 {
		t.Errorf("the Size method is bug!")
	}

	list.Remove(4)
	if list.Size() != 4 {
		t.Errorf("the Remove method is bug!")
	}

	if list.Contains('A') {
		t.Errorf("the Remove method is bug!")
	}
	value, _ := list.Get(1)
	if value != 2 {
		t.Errorf("the Get method is bug!")
	}
	list.Clear()

	if list.Size() != 0 {
		t.Errorf("the Clear method is bug!")
	}
}
