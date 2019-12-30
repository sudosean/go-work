package main

import (
	"container/list"

	"fmt"
)

func main()  {
	var list list.List
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	for e := list.Front(); e != nil; e=e.Next(){
		fmt.Println(e.Prev())
	}
}