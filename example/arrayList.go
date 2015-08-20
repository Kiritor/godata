package main

import (
	"github.com/kiritor/godata/list/arrayList"
	"fmt"
)

func main(){
	list := arrayList.New()
	fmt.Println(list.String())             //ArrayList[]
	list.Add(`A`,'B',"CDE")
	fmt.Println(list.String())             //ArrayList[A,66,CDE]
	fmt.Println(list.Get(0))               //A true
    list.Remove(1)
	fmt.Println(list.String())             //ArrayList[A,CDE]]
	fmt.Println(list.Values())             //[A CDE
	fmt.Println(list.Contains("CDE"))      //true
	list.Clear()
	fmt.Println(list.String())             //ArrayList[]
	fmt.Println(list.Size())               //0
}
