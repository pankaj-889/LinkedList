package Doubly

import (
	"fmt"
	"strings"
)

type Node struct {
	Value int
	Prev  *Node
	Next  *Node
}

type DoublyLL struct {
	head *Node
}

func (n Node) String() string {
	return fmt.Sprintf("%d", n.Value)
}

func (l *DoublyLL) Add(value int) {
	newNode := new(Node)
	newNode.Value = value
	if l.head == nil {
		l.head = newNode
		return
	}

	temp := l.head
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = newNode
	newNode.Prev = temp
}

func (l *DoublyLL) Remove(value int) {

	if l.head == nil {
		return
	}

	if l.head.Value == value {
		l.head = l.head.Next
		return
	}

	prev := l.head
	curr := l.head.Next

	for curr.Next != nil {
		if curr.Value == value {
			temp := curr.Next
			prev.Next = temp
			curr.Prev = nil
			curr.Next = nil
			temp.Prev = prev
			return
		}
		prev = curr
		curr = curr.Next
	}
	prev.Next = nil
	curr.Prev = nil
}

func (l DoublyLL) Get(value int) *Node {

	temp := l.head
	for temp != nil {
		if temp.Value == value {
			return temp
		}
		temp = temp.Next
	}
	return nil
}

func (l DoublyLL) String() string {
	temp := l.head
	sb := strings.Builder{}
	for temp != nil {
		sb.WriteString(fmt.Sprintf("%d -> ", temp.Value))
		temp = temp.Next
	}
	return sb.String()
}
