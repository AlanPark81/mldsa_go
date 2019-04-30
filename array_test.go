package mldsa_go

import (
	"fmt"
	"math"
	"testing"
)

func TestGetOutOfBound(t *testing.T) {
	var array = CreateArray(1)

	if array.Get(1) != nil {
		t.Error("get success out of bound")
	}

	if array.Get(-1) != nil {
		t.Error("get success out of bound")
	}
}

func TestEmptyArrayAccess(t *testing.T) {
	var array = CreateArray(1)

	if array.Set(1, 1) != false {
		t.Error("set success out of bound")
	}

	if array.Set(-1, 1) != false {
		t.Error("set success out of bound")
	}
}

func TestCreateAndGet(t *testing.T) {
	var array = CreateArray(2)

	if array.Get(0) != nil {
		t.Error("not nil without assignment")
	}

	if array.Get(1) != nil {
		t.Error("not nil without assignment")
	}
}

func TestCreateSetGet(t *testing.T) {
	var array = CreateArray(2)

	if array.Get(0) != nil {
		t.Error("not assigned element")
	}

	if array.Set(0, 1) == false {
		t.Error("set failed")
	}

	if array.Get(0) != 1 {
		t.Error("element set value changed")
	}

	if array.Get(1) != nil {
		t.Error("not assigned element")
	}

	if array.Set(1, 2) == false {
		t.Error("set failed")
	}

	if array.Get(1) != 2 {
		t.Error("element set value changed")
	}
}

func TestCreateSetGetOutOfBound(t *testing.T) {
	var array = CreateArray(2)

	if array.Set(-1, 1) == true {
		t.Error("set success out of bound")
		return
	}

	if array.Get(-1) == 1 {
		t.Error("get success out of bound")
		return
	}

	if array.Set(2, 2) == true {
		t.Error("set success out of bound")
		return
	}

	if array.Get(2) == 2 {
		t.Error("get success out of bound")
	}
}

func BenchmarkArrayGet(b *testing.B) {
	for k := 0.; k <= 11; k++ {
		n := int(math.Pow(2, k))
		var array = CreateArray(n)
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				array.Get(n - 1)
			}
		})
	}
}

func BenchmarkArraySet(b *testing.B) {
	for k := 0.; k <= 11; k++ {
		n := int(math.Pow(2, k))
		var array = CreateArray(n)
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				array.Set(n-1, 1)
			}
		})
	}
}

func BenchmarkArrayCreate(b *testing.B) {
	for k := 0.; k <= 11; k++ {
		n := int(math.Pow(2, k))
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var _ = CreateArray(n)
			}
		})
	}
}
