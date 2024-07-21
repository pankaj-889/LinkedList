package container

import "container/list"

type Node struct {
	key   int
	value int
}

type DoublyLL struct {
	list *list.List
}

func NewDoublyLL() *DoublyLL {
	return &DoublyLL{list: list.New()}
}

func (ll *DoublyLL) InsertFront(value int) {
	ll.list.PushFront(value)
}

func (ll *DoublyLL) InsertBack(value int) {
	ll.list.PushBack(value)
}

func (ll *DoublyLL) InsertInBetween(value int, elem *list.Element) {
	ll.list.InsertAfter(value, elem)
}

func (ll *DoublyLL) Delete(value int) {
	for e := ll.list.Front(); e != nil; e = e.Next() {
		if e.Value.(int) == value {
			ll.list.Remove(e)
			break
		}
	}
}

func (ll *DoublyLL) DisplayForward() {
	for e := ll.list.Front(); e != nil; e = e.Next() {
		println(e.Value.(int))
	}
}

func (ll *DoublyLL) DisplayBackward() {
	for e := ll.list.Back(); e != nil; e = e.Prev() {
		println(e.Value.(int))
	}
}
