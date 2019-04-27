package mldsa_go

type Array struct {
	arrayIn []interface{}
}

func CreateArray(size int) *Array {
	return &Array {
		make([]interface{}, size),
	}
}

func (c *Array) Get(index int) interface{} {
	if index >= len(c.arrayIn) || index < 0 {
		return nil
	}
	return c.arrayIn[index]
}

func (c *Array) Set(index int, value interface{}) bool {
	if index >= len(c.arrayIn) || index < 0 {
		return false
	}
	c.arrayIn[index] = value
	return true
}