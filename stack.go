package mldsa_go

type Stack struct {
	list *LinkedList
}

func CreateStack() *Stack {
	return &Stack{
		CreateLinkedList(),
	}
}

func (c *Stack) Push(data interface{}) bool {
	return c.list.AddTail(data)
}

func (c *Stack) Pop() bool {
	var result = c.list.RemoveTail()
	return result
}

func (c *Stack) Peek() interface{} {
	return c.list.GetTail()
}
