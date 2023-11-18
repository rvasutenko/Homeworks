package main

import "fmt"

type OneLinkedList struct {
	head     *Node
	quantity int
}

type Node struct {
	value    int
	nextNode *Node
}

func (s *OneLinkedList) Add(value int) {
	newElem := new(Node)
	newElem.value = value
	if s.quantity <= 0 {
		s.head = newElem
	} else {
		elem := s.head
		for i := 0; i < s.quantity; i++ {
			if elem.nextNode != nil {
				elem = elem.nextNode
			}
		}
		elem.nextNode = newElem
	}
	s.quantity += 1
}

func (s *OneLinkedList) Pop() Node {
	var popElem Node
	if s.quantity <= 0 {
		fmt.Println("ERROR: NO ELEM")
		return Node{}
	} else {
		elem := s.head
		for i := 1; i < s.quantity-1; i++ {
			elem = elem.nextNode
		}
		if s.quantity != 1 {
			popElem = *
		}
		popElem = *elem.nextNode
		elem.nextNode = nil
	}
	s.quantity -= 1
	return popElem
}

func main() {
	LinkedList := OneLinkedList{nil, 0}
	LinkedList.Add(3)
	LinkedList.Add(4)
	LinkedList.Add(5)
	fmt.Println(LinkedList.Pop().value)
}
