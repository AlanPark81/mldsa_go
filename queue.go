package mldsa_go

type Queue struct {
	list *LinkedList
}

func CreateQueue() *Queue {
	return &Queue{
		CreateLinkedList(),
	}
}

func (c *Queue) Push(data interface{}) bool {
	return c.list.AddTail(data)
}

func (c *Queue) Pop() interface{} {
	var ret_val = c.list.GetHead()
	if !c.list.RemoveHead() {
		return false
	}
	return ret_val
}
