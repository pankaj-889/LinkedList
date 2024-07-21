package main

import (
	. "LinkedList/Doubly"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	l := DoublyLL{}
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Remove(2)
	fmt.Println(l)

}
