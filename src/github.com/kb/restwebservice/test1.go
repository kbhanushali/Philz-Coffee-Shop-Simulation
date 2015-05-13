package main

import (
	"./queue"
	"fmt"
)


func main() {

	q := queue.NewQueue()
	q.Push(4)
	q.Push(9)

	if ele, ok := q.Pop(); ok {
		fmt.Printf("Content : %v\n", ele)
		fmt.Printf("Size : %d\n", q.Size())
	}
}