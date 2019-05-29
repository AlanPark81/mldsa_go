package mldsa_go

import "testing"

func TestStack_EmptyCase(t *testing.T) {
	var stack = CreateStack()

	if stack.Pop() != false {
		t.Error("Create made a non-empty stack")
		return
	}

	if stack.Peek() != nil {
		t.Error("Create made a non-empty stack")
		return
	}
	if stack.Push(1) == false {
		t.Error("push fail")
		return
	}

	if stack.Peek() != 1 {
		t.Error("unexpected peek value")
		return
	}

	if stack.Pop() != true {
		t.Error("pop fail")
		return
	}
}

func TestStack_TwoCase(t *testing.T) {
	var stack = CreateStack()
	if stack.Push(1) == false {
		t.Error("Push failed")
		return
	}

	if stack.Peek() != 1 {
		t.Error("unexpected peek value")
		return
	}

	if stack.Push(2) == false {
		t.Error("Push failed")
		return
	}

	if stack.Peek() != 2 {
		t.Error("unexpected peek value")
		return
	}

	if stack.Pop() != true {
		t.Error("pop failed")
		return
	}

	if stack.Peek() != 1 {
		t.Error("unexpected peek value")
		return
	}

	if stack.Pop() != true {
		t.Error("pop failed")
		return
	}

	if stack.Peek() != nil {
		t.Error("unexpected peek value")
		return
	}
}
