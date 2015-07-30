package main

import (
	"godata/list/arrayList"
	"fmt"
)

func main(){
	list := arrayList.New()
	fmt.Println(list.String())
	list.Add(`A`,'B',"CDE")
	fmt.Println(list.String())
	fmt.Println(list.Get(0))
    list.Remove(1)
	fmt.Println(list.String())
	fmt.Println(list.Values())
	fmt.Println(list.Contains("CDE"))
	list.Clear()
	fmt.Println(list.String())
	fmt.Println(list.Size())
}
