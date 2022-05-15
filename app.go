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

func (s *mySingleLinkedList) addToHead(name string) {
	newNode := &node{
		data: name,
	}
	if s.head == nil {
		s.head = newNode
	} else {
		newNode.next = s.head
		s.head = newNode
	}
	s.size++
}
func (s *mySingleLinkedList) addToTail(name string) {
	newNode := &node{
		data: name,
	}
	if s.head == nil {
		s.head = newNode
	} else {
		current := s.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	s.size++
}

func (s *mySingleLinkedList) iterateList() {
	for node := s.head; node != nil; node = node.next {
		fmt.Println(node.data)
	}
}

func newLinkedList() *mySingleLinkedList {
	return &mySingleLinkedList{}
}

func main() {
	singleList := newLinkedList()
	singleList.addToHead("Maya")
	singleList.addToHead("Dani")
	singleList.addToTail("Victor")
	singleList.addToTail("Zupri")
	singleList.iterateList()
	fmt.Println(singleList.size)
}
