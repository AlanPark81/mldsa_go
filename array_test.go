package mldsa_go

import (
	"testing"
)

func TestGetOutOfBound(t *testing.T) {
	var array = CreateArray(1)

	if array.Get(1) != nil {
		t.Error("failed")
	}

	if array.Get(-1) != nil {
		t.Error("failed")
	}
}

func TestEmptyArrayAccess(t *testing.T) {
	var array = CreateArray(1)

	if array.Set(1, 1) != false {
		t.Error("failed")
	}

	if array.Set(-1, 1) != false {
		t.Error("failed")
	}
}

func TestCreateAndGet(t *testing.T) {
	var array = CreateArray(2)

	if array.Get(0) != nil {
		t.Error("failed")
	}

	if array.Get(1) != nil {
		t.Error("failed")
	}
}

func TestCreateSetGet(t *testing.T) {
	var array = CreateArray(2)

	if array.Get(0) != nil {
		t.Error("failed")
	}

	if array.Set(0, 1) == false {
		t.Error("failed")
	}

	if array.Get(0) != 1 {
		t.Error("failed")
	}

	if array.Get(1) != nil {
		t.Error("failed")
	}

	if array.Set(1, 2) == false {
		t.Error("failed")
	}

	if array.Get(1) != 2 {
		t.Error("failed")
	}
}

func TestCreateSetGetOutOfBound(t *testing.T) {
	var array = CreateArray(2)

	if array.Set(-1, 1) == true {
		t.Error("failed")
		return
	}

	if array.Get(-1) == 1 {
		t.Error("failed")
		return
	}

	if array.Set(2, 2) == true {
		t.Error("failed")
		return
	}

	if array.Get(2) == 2 {
		t.Error("failed")
	}
}

func benchmarkArrayGet(size int, b *testing.B) {
	var array = CreateArray(size)
	for n := 0; n < b.N; n++ {
		array.Get(size - 1)
	}
}

func BenchmarkArrayGet1(b *testing.B) {
	benchmarkArrayGet(1, b)
}

func BenchmarkArrayGet2(b *testing.B) {
	benchmarkArrayGet(2, b)
}

func BenchmarkArrayGet4(b *testing.B) {
	benchmarkArrayGet(4, b)
}

func BenchmarkArrayGet8(b *testing.B) {
	benchmarkArrayGet(8, b)
}

func BenchmarkArrayGet16(b *testing.B) {
	benchmarkArrayGet(16, b)
}

func BenchmarkArrayGet32(b *testing.B) {
	benchmarkArrayGet(32, b)
}

func BenchmarkArrayGet64(b *testing.B) {
	benchmarkArrayGet(64, b)
}

func BenchmarkArrayGet128(b *testing.B) {
	benchmarkArrayGet(128, b)
}

func BenchmarkArrayGet256(b *testing.B) {
	benchmarkArrayGet(256, b)
}

func BenchmarkArrayGet512(b *testing.B) {
	benchmarkArrayGet(512, b)
}

func BenchmarkArrayGet1024(b *testing.B) {
	benchmarkArrayGet(1024, b)
}

func BenchmarkArrayGet2048(b *testing.B) {
	benchmarkArrayGet(2048, b)
}

func benchmarkArraySet(size int, b *testing.B) {
	var array = CreateArray(size)
	for n := 0; n < b.N; n++ {
		array.Set(size-1, 1)
	}
}

func BenchmarkArraySet1(b *testing.B) {
	benchmarkArraySet(1, b)
}

func BenchmarkArraySet2(b *testing.B) {
	benchmarkArraySet(2, b)
}

func BenchmarkArraySet4(b *testing.B) {
	benchmarkArraySet(4, b)
}

func BenchmarkArraySet8(b *testing.B) {
	benchmarkArraySet(8, b)
}

func BenchmarkArraySet16(b *testing.B) {
	benchmarkArraySet(16, b)
}

func BenchmarkArraySet32(b *testing.B) {
	benchmarkArraySet(32, b)
}

func BenchmarkArraySet64(b *testing.B) {
	benchmarkArraySet(64, b)
}

func BenchmarkArraySet128(b *testing.B) {
	benchmarkArraySet(128, b)
}

func BenchmarkArraySet256(b *testing.B) {
	benchmarkArraySet(256, b)
}

func BenchmarkArraySet512(b *testing.B) {
	benchmarkArraySet(512, b)
}

func BenchmarkArraySet1024(b *testing.B) {
	benchmarkArraySet(1024, b)
}

func BenchmarkArraySet2048(b *testing.B) {
	benchmarkArraySet(2048, b)
}

func benchmarkArrayCreate(size int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		var _ = CreateArray(size)
	}
}

func BenchmarkArrayCreate1(b *testing.B) {
	benchmarkArrayCreate(1, b)
}

func BenchmarkArrayCreate2(b *testing.B) {
	benchmarkArrayCreate(2, b)
}

func BenchmarkArrayCreate4(b *testing.B) {
	benchmarkArrayCreate(4, b)
}

func BenchmarkArrayCreate8(b *testing.B) {
	benchmarkArrayCreate(8, b)
}

func BenchmarkArrayCreate16(b *testing.B) {
	benchmarkArrayCreate(16, b)
}

func BenchmarkArrayCreate32(b *testing.B) {
	benchmarkArrayCreate(32, b)
}

func BenchmarkArrayCreate64(b *testing.B) {
	benchmarkArrayCreate(64, b)
}

func BenchmarkArrayCreate128(b *testing.B) {
	benchmarkArrayCreate(128, b)
}

func BenchmarkArrayCreate256(b *testing.B) {
	benchmarkArrayCreate(256, b)
}

func BenchmarkArrayCreate512(b *testing.B) {
	benchmarkArrayCreate(512, b)
}

func BenchmarkArrayCreate1024(b *testing.B) {
	benchmarkArrayCreate(1024, b)
}

func BenchmarkArrayCreate2048(b *testing.B) {
	benchmarkArrayCreate(2048, b)
}
