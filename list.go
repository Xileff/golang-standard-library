package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Felix1")
	data.PushBack("Felix2")
	data.PushBack("Felix3")

	// var head *list.Element = data.Front()
	// fmt.Println(head.Value)

	// next := head.Next()
	// fmt.Println(next.Value)

	// next = head.Next()
	// fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
