package LRUCache

import "container/list"

type Node struct {
	key   int
	value int
}

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	if elem, ok := c.cache[key]; ok {
		c.list.MoveToFront(elem)
		return elem.Value.(*Node).value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if elem, ok := c.cache[key]; ok {
		c.list.MoveToFront(elem)
		elem.Value.(*Node).value = value
		return
	}

	if c.list.Len() == c.capacity {
		back := c.list.Back()
		if back != nil {
			node := back.Value.(*Node)
			delete(c.cache, node.key) // deleting key from the map
			c.list.Remove(back)       // removing node from the list
		}
	}

	newNode := new(Node)
	newNode.key = key
	newNode.value = value

	elem := c.list.PushFront(newNode)
	c.cache[key] = elem
}
