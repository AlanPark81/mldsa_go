package mldsa_go

type LinkedListNode struct {
	data interface{}
	next *LinkedListNode
}

type LinkedList struct {
	head *LinkedListNode
}

func CreateLinkedList() *LinkedList {
	return &LinkedList{
		nil,
	}
}

func (c *LinkedList) AddHead(data interface{}) bool {
	var node = &LinkedListNode{data, c.head}
	c.head = node
	return true
}

func (c *LinkedList) RemoveHead() bool {
	if c.head == nil {
		return false
	}
	c.head = c.head.next
	return true
}

func (c *LinkedList) GetHead() interface{} {
	if c.head == nil {
		return nil
	}
	return c.head.data
}

func (c *LinkedList) AddTail(data interface{}) bool {
	var node = c.head
	if node == nil {
		c.head = &LinkedListNode{data, nil}
		return true
	}
	for node.next != nil {
		node = node.next
	}
	node.next = &LinkedListNode{data, nil}
	return true
}

func (c *LinkedList) RemoveTail() bool {
	var node = c.head
	if c.head == nil {
		return false
	}
	if c.head.next == nil {
		c.head = nil
		return true
	}
	var prev_node = node
	for node.next != nil {
		prev_node = node
		node = node.next
	}
	prev_node.next = nil
	return true
}

func (c *LinkedList) GetTail() interface{} {
	var node = c.head
	if node == nil {
		return nil
	}
	for node.next != nil {
		node = node.next
	}
	return node.data
}

func (c *LinkedList) IsEmpty() bool {
	return c.head == nil
}

func (c *LinkedList) Get(index int) interface{} {
	if index < 0 {
		return nil
	}
	var node = c.head

	if node == nil {
		return nil
	}

	for i := 0; i < index; i++ {
		if node.next == nil {
			return nil
		}
		node = node.next
	}
	return node.data
}

func (c *LinkedList) Set(index int, data interface{}) bool {
	if index < 0 {
		return false
	}
	var node = c.head
	for i := 0; i < index; i++ {
		if node.next == nil {
			return false
		}
		node = node.next
	}
	node.data = data
	return true
}

func (c *LinkedList) Size() int {
	var retVal = 0
	for node := c.head; node != nil; node = node.next {
		retVal++
	}
	return retVal
}
