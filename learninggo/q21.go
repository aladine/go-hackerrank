package main

import "fmt"
import "container/list"

func main() {
	l := list.New()
	// []int{1, 2, 3})
	l.PushFront(1)
	l.PushFront(2)
	for e := l.Front(); e != nil; e = e.Next() {

		fmt.Printf("%#v", e.Value)
	}
}
