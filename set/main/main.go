package main

import (
	"godata/set/hashSet"
	"fmt"
	"godata/set"
)

func main(){
	set1 := hashSet.NewHashSet()
	set1.Add(5)
	set1.Add(6)
	set2 := hashSet.NewHashSet()
	set2.Add(5)
	set2.Add(88)
	fmt.Println(set1.String())
	fmt.Println(set1.Contains(5))
	fmt.Println(set2.Elements())
	fmt.Println(set.IsSuperSet(set1,set2))
	fmt.Println(set.Union(set1,set2))
	h :=make(map[string][]string)
	h1 :=make([]string,10)
	h1[0] ="hehe"
	h["key1"] = h1
	thread := hashSet.NewThreadHashSet()
	fmt.Println(h)
	     fmt.Println(thread.HashSet.String())
}
