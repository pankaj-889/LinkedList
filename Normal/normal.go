package main

import (
	"fmt"
	"strings"
)

type node struct {
	value int
	next  *node
}

type linkedList struct {
	head *node
	len  int
}

func (n node) String() string {
	return fmt.Sprintf("%d", n.value)
}

func (l *linkedList) add(value int) {
	newNode := new(node)
	newNode.value = value

	if l.head == nil {
		l.head = newNode
		return
	}
	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}

func (l *linkedList) remove(value int) {
	if l.head == nil {
		return
	}

	if l.head.value == value {
		l.head = l.head.next
		l.len--
		return
	}

	prev := l.head
	curr := l.head.next

	for curr != nil {
		if curr.value == value {
			prev.next = curr.next
			l.len--
			return
		}
		prev = curr
		curr = curr.next
	}
}

func (l linkedList) get(value int) *node {
	temp := l.head
	for temp != nil {
		if temp.value == value {
			return temp
		}
		temp = temp.next
	}
	return nil
}

func (l linkedList) String() string {
	temp := l.head
	sb := strings.Builder{}
	for temp != nil {
		sb.WriteString(fmt.Sprintf("%d -> ", temp.value))
		temp = temp.next
	}
	return sb.String()
}

func main() {
	fmt.Println("hello world")
	l := linkedList{}
	l.add(1)
	l.add(2)
	l.add(3)
	l.add(4)
	l.remove(5)
	fmt.Println(l)

}
