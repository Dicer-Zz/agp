package main

import (
	"container/list"
	"fmt"
)

func main() {
	li := list.New()
	li.PushBack(101)
	li.PushBack(102)
	li.PushBack(103)

	for e := li.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
