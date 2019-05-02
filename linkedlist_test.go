package mldsa_go

import (
	"fmt"
	"math"
	"testing"
)

func TestCreateLinkedList(t *testing.T) {
	var linkedlist = CreateLinkedList()

	if linkedlist.Get(0) != nil {
		t.Error("Create made a non-empty list")
		return
	}

	if linkedlist.IsEmpty() == false {
		t.Error("Create made a non-empty list")
		return
	}

	if linkedlist.GetHead() != nil {
		t.Error("Create made a non-empty list")
		return
	}

	if linkedlist.GetTail() != nil {
		t.Error("Create made a non-empty list")
		return
	}

	if linkedlist.RemoveHead() == true {
		t.Error("RemoveHead succeed in empty list")
		return
	}

	if linkedlist.RemoveTail() == true {
		t.Error("RemoveTail succeed in empty list")
		return
	}
}

func TestLinkedList_AddHead(t *testing.T) {
	var linkedlist = CreateLinkedList()
	if linkedlist.AddHead(3) == false {
		t.Error("AddHead failed")
		return
	}

	if linkedlist.GetHead() != 3 {
		t.Error("value mismatch after AddHead and GetHead")
		return
	}

	if linkedlist.GetTail() != 3 {
		t.Error("value mismatch after AddHead and GetTail for the empty list")
		return
	}
}

func TestLinkedList_AddTail(t *testing.T) {
	var linkedlist = CreateLinkedList()
	if linkedlist.AddTail(3) == false {
		t.Error("AddTail failed")
		return
	}

	if linkedlist.GetHead() != 3 {
		t.Error("value mismatch after AddHead and GetHead")
		return
	}

	if linkedlist.GetTail() != 3 {
		t.Error("value mismatch after AddHead and GetTail for the empty list")
		return
	}
}

func TestLinkedList_AddTailRepetitionSize(t *testing.T) {
	var linkedlist = CreateLinkedList()
	if linkedlist.AddTail(1) == false {
		t.Error("AddTail failed")
		return
	}
	if linkedlist.AddTail(2) == false {
		t.Error("AddTail failed")
		return
	}
	if linkedlist.AddTail(3) == false {
		t.Error("AddTail failed")
		return
	}

	if linkedlist.Size() != 3 {
		t.Error("size mismatch")
		return
	}
}

func TestLinkedList_AddHeadRepetitionSize(t *testing.T) {
	var linkedlist = CreateLinkedList()
	if linkedlist.AddHead(1) == false {
		t.Error("AddTail failed")
		return
	}
	if linkedlist.AddHead(2) == false {
		t.Error("AddTail failed")
		return
	}
	if linkedlist.AddHead(3) == false {
		t.Error("AddTail failed")
		return
	}

	if linkedlist.Size() != 3 {
		t.Error("size mismatch")
		return
	}
}

func TestLinkedList_MixedAddRepetitionSize(t *testing.T) {
	var linkedlist = CreateLinkedList()
	if linkedlist.AddHead(1) == false {
		t.Error("AddTail failed")
		return
	}
	if linkedlist.AddTail(2) == false {
		t.Error("AddTail failed")
		return
	}
	if linkedlist.AddHead(3) == false {
		t.Error("AddTail failed")
		return
	}
	if linkedlist.AddTail(4) == false {
		t.Error("AddTail failed")
		return
	}
	if linkedlist.AddHead(5) == false {
		t.Error("AddTail failed")
		return
	}

	if linkedlist.Size() != 5 {
		t.Error("size mismatch")
		return
	}
}

func TestLinkedList_MixedAddRepetitionGet(t *testing.T) {
	var linkedlist = CreateLinkedList()
	linkedlist.AddHead(1)
	linkedlist.AddTail(2)
	linkedlist.AddHead(3)
	linkedlist.AddTail(4)
	linkedlist.AddHead(5)

	if linkedlist.GetHead() != 5 {
		t.Error("AddHead GetHead mismatch")
		return
	}

	if linkedlist.GetTail() != 4 {
		t.Error("AddTail GetTail mismatch")
		return
	}

	if linkedlist.Get(-1) != nil {
		t.Error("value found out of bound")
		return
	}

	if linkedlist.Get(0) != 5 {
		t.Error("value mismatch after Added Node and Get Index")
		return
	}

	if linkedlist.Get(1) != 3 {
		t.Error("value mismatch after Added Node and Get Index")
		return
	}

	if linkedlist.Get(2) != 1 {
		t.Error("value mismatch after Added Node and Get Index")
		return
	}

	if linkedlist.Get(3) != 2 {
		t.Error("value mismatch after Added Node and Get Index")
		return
	}

	if linkedlist.Get(4) != 4 {
		t.Error("value mismatch after Added Node and Get Index")
		return
	}

	if linkedlist.Get(5) != nil {
		t.Error("value found out of bound")
		return
	}
}

func TestLinkedList_MixedAddRepetitionRemoveHeadRepetition(t *testing.T) {
	var linkedlist = CreateLinkedList()
	linkedlist.AddHead(1)
	linkedlist.AddTail(2)
	linkedlist.AddHead(3)
	linkedlist.AddTail(4)
	linkedlist.AddHead(5)

	if linkedlist.GetHead() != 5 {
		t.Error("value mismatch after Added Node and Get Index")
		return
	}

	if linkedlist.RemoveHead() == false {
		t.Error("RemoveHead failed")
		return
	}

	if linkedlist.GetHead() != 3 {
		t.Error("value mismatch after Added Node and GetHead after remove head")
		return
	}

	if linkedlist.RemoveHead() == false {
		t.Error("RemoveHead failed")
		return
	}

	if linkedlist.GetHead() != 1 {
		t.Error("value mismatch after Added Node and Get Index")
		return
	}

	if linkedlist.RemoveHead() == false {
		t.Error("RemoveHead failed")
		return
	}

	if linkedlist.GetHead() != 2 {
		t.Error("value mismatch after Added Node and Get Index")
		return
	}

	if linkedlist.RemoveHead() == false {
		t.Error("RemoveHead failed")
		return
	}

	if linkedlist.GetHead() != 4 {
		t.Error("value mismatch after Added Node and Get Index")
		return
	}

	if linkedlist.RemoveHead() == false {
		t.Error("RemoveHead failed")
		return
	}

	if linkedlist.GetHead() != nil {
		t.Error("value found out of bound")
		return
	}
}

func TestLinkedList_MixedAddRepetitionRemoveTailRepetition(t *testing.T) {
	var linkedlist = CreateLinkedList()
	linkedlist.AddHead(1)
	linkedlist.AddTail(2)
	linkedlist.AddHead(3)
	linkedlist.AddTail(4)
	linkedlist.AddHead(5)

	if linkedlist.GetTail() != 4 {
		t.Error("value mismatch after Added Node and Get Index")
		return
	}

	if linkedlist.RemoveTail() == false {
		t.Error("RemoveHead failed")
		return
	}

	if linkedlist.GetTail() != 2 {
		t.Error("value mismatch after Added Node and GetHead after remove head")
		return
	}

	if linkedlist.RemoveTail() == false {
		t.Error("RemoveHead failed")
		return
	}

	if linkedlist.GetTail() != 1 {
		t.Error("value mismatch after Added Node and Get Index")
		return
	}

	if linkedlist.RemoveTail() == false {
		t.Error("RemoveHead failed")
		return
	}

	if linkedlist.GetTail() != 3 {
		t.Error("value mismatch after Added Node and Get Index")
		return
	}

	if linkedlist.RemoveTail() == false {
		t.Error("RemoveHead failed")
		return
	}

	if linkedlist.GetTail() != 5 {
		t.Error("value mismatch after Added Node and Get Index")
		return
	}

	if linkedlist.RemoveTail() == false {
		t.Error("RemoveHead failed")
		return
	}

	if linkedlist.GetTail() != nil {
		t.Error("value found out of bound")
		return
	}
}

func TestLinkedList_AddTailRepetitionGet(t *testing.T) {
	var linkedlist = CreateLinkedList()
	if linkedlist.AddTail(1) == false {
		t.Error("AddTail failed")
		return
	}
	if linkedlist.AddTail(2) == false {
		t.Error("AddTail failed")
		return
	}
	if linkedlist.AddTail(3) == false {
		t.Error("AddTail failed")
		return
	}

	if linkedlist.Get(-1) != nil {
		t.Error("value found out of bound")
		return
	}

	if linkedlist.Get(0) != 1 {
		t.Error("value mismatch after Added Tail and Get Index")
		return
	}

	if linkedlist.Get(1) != 2 {
		t.Error("value mismatch after Added Tail and Get Index")
		return
	}

	if linkedlist.Get(2) != 3 {
		t.Error("value mismatch after Added Tail and Get Index")
		return
	}

	if linkedlist.Get(3) != nil {
		t.Error("value found out of bound")
		return
	}
}

func TestLinkedList_AddHeadRepetitionGet(t *testing.T) {
	var linkedlist = CreateLinkedList()
	if linkedlist.AddHead(1) == false {
		t.Error("AddTail failed")
		return
	}
	if linkedlist.AddHead(2) == false {
		t.Error("AddTail failed")
		return
	}
	if linkedlist.AddHead(3) == false {
		t.Error("AddTail failed")
		return
	}

	if linkedlist.Get(-1) != nil {
		t.Error("value found out of bound")
		return
	}

	if linkedlist.Get(0) != 3 {
		t.Error("value mismatch after Added Tail and Get Index")
		return
	}

	if linkedlist.Get(1) != 2 {
		t.Error("value mismatch after Added Tail and Get Index")
		return
	}

	if linkedlist.Get(2) != 1 {
		t.Error("value mismatch after Added Tail and Get Index")
		return
	}

	if linkedlist.Get(3) != nil {
		t.Error("value found out of bound")
		return
	}
}

func TestLinkedList_AddHeadRepetitionSet(t *testing.T) {
	var linkedlist = CreateLinkedList()
	if linkedlist.AddHead(1) == false {
		t.Error("AddTail failed")
		return
	}
	if linkedlist.AddHead(2) == false {
		t.Error("AddTail failed")
		return
	}
	if linkedlist.AddHead(3) == false {
		t.Error("AddTail failed")
		return
	}

	if linkedlist.Set(-1, 3) == true {
		t.Error("Set Out Of Bound succeed")
		return
	}

	if linkedlist.Set(0, 4) == false {
		t.Error("Set failed")
		return
	}

	if linkedlist.Set(1, 5) == false {
		t.Error("Set failed")
		return
	}

	if linkedlist.Set(2, 6) == false {
		t.Error("Set failed")
		return
	}

	if linkedlist.Set(3, 3) == true {
		t.Error("Set Out Of Bound succeed")
		return
	}

	if linkedlist.Get(0) != 4 {
		t.Error("value mismatch Set and Get Index")
		return
	}

	if linkedlist.Get(1) != 5 {
		t.Error("value mismatch Set and Get Index")
		return
	}

	if linkedlist.Get(2) != 6 {
		t.Error("value mismatch Set and Get Index")
		return
	}
}

func TestLinkedList_AddTailRepetitionSet(t *testing.T) {
	var linkedlist = CreateLinkedList()
	if linkedlist.AddTail(1) == false {
		t.Error("AddTail failed")
		return
	}
	if linkedlist.AddTail(2) == false {
		t.Error("AddTail failed")
		return
	}
	if linkedlist.AddTail(3) == false {
		t.Error("AddTail failed")
		return
	}

	if linkedlist.Set(-1, 3) == true {
		t.Error("Set Out Of Bound succeed")
		return
	}

	if linkedlist.Set(0, 4) == false {
		t.Error("Set failed")
		return
	}

	if linkedlist.Set(1, 5) == false {
		t.Error("Set failed")
		return
	}

	if linkedlist.Set(2, 6) == false {
		t.Error("Set failed")
		return
	}

	if linkedlist.Set(3, 7) == true {
		t.Error("Set Out Of Bound succeed")
		return
	}

	if linkedlist.Get(0) != 4 {
		t.Error("value mismatch Set and Get Index")
		return
	}

	if linkedlist.Get(1) != 5 {
		t.Error("value mismatch Set and Get Index")
		return
	}

	if linkedlist.Get(2) != 6 {
		t.Error("value mismatch Set and Get Index")
		return
	}
}

func BenchmarkLinkedList_Get(b *testing.B) {
	for k := 0.; k <= 11; k++ {
		n := int(math.Pow(2, k))
		var list = CreateLinkedList()
		for i := 0; i < n; i++ {
			list.AddTail(i)
		}
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				list.Get(n - 1)
			}
		})
	}
}

func BenchmarkLinkedList_Set(b *testing.B) {
	for k := 0.; k <= 11; k++ {
		n := int(math.Pow(2, k))
		var list = CreateLinkedList()
		for i := 0; i < n; i++ {
			list.AddTail(i)
		}
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				list.Set(n-1, 1)
			}
		})
	}
}

func BenchmarkLinkedList_Size(b *testing.B) {
	for k := 0.; k <= 11; k++ {
		n := int(math.Pow(2, k))
		var list = CreateLinkedList()
		for i := 0; i < n; i++ {
			list.AddTail(i)
		}
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var _ = list.Size()
			}
		})
	}
}

func BenchmarkLinkedList_AddHead(b *testing.B) {
	for k := 0.; k <= 11; k++ {
		n := int(math.Pow(2, k))
		var list = CreateLinkedList()
		for i := 0; i < n; i++ {
			list.AddTail(i)
		}
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				list.AddHead(i)
			}
		})
	}
}

func BenchmarkLinkedList_GetHead(b *testing.B) {
	for k := 0.; k <= 11; k++ {
		n := int(math.Pow(2, k))
		var list = CreateLinkedList()
		for i := 0; i < n; i++ {
			list.AddTail(i)
		}
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var _ = list.GetHead()
			}
		})
	}
}

func BenchmarkLinkedList_AddTail(b *testing.B) {
	for k := 0.; k <= 11; k++ {
		n := int(math.Pow(2, k))
		var list = CreateLinkedList()
		for i := 0; i < n; i++ {
			list.AddTail(i)
		}
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				list.AddTail(i)
				list.RemoveHead()
			}
		})
	}
}

func BenchmarkLinkedList_GetTail(b *testing.B) {
	for k := 0.; k <= 11; k++ {
		n := int(math.Pow(2, k))
		var list = CreateLinkedList()
		for i := 0; i < n; i++ {
			list.AddTail(i)
		}
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var _ = list.GetTail()
			}
		})
	}
}
