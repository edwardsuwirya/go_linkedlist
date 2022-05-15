package main

import "fmt"

type node struct {
	data string
	next *node
}

type mySingleLinkedList struct {
	size int
	head *node
}

func newLinkedList() *mySingleLinkedList {
	return &mySingleLinkedList{}
}

func main() {
	singleList := newLinkedList()
	fmt.Println(singleList.size)
}
