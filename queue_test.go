package mldsa_go

import (
	"testing"
)

func TestQueue_EmptyCase(t *testing.T) {
	var queue = CreateQueue()

	if queue.Pop() != false {
		t.Error("Create made a non-empty list")
		return
	}
}

func TestQueue_OneCase(t *testing.T) {
	var queue = CreateQueue()
	if queue.Push(1) == false {
		t.Error("Push failed")
		return
	}

	if queue.Pop() != 1 {
		t.Error("non-pushed value poped")
		return
	}
}

func TestQueue_TwoCase(t *testing.T) {
	var queue = CreateQueue()
	if queue.Push(1) == false {
		t.Error("Push failed")
		return
	}

	if queue.Push(2) == false {
		t.Error("Push failed")
		return
	}

	if queue.Pop() != 1 {
		t.Error("not expected value poped")
		return
	}

	if queue.Pop() != 2 {
		t.Error("not expected value poped")
		return
	}
}
