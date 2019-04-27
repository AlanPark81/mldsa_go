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

func benchmarkArray_Get(size int, b *testing.B) {
	var array = CreateArray(size)
	for n := 0; n < b.N; n++ {
		array.Get(size - 1)
	}
}

func BenchmarkArray_Get1(b *testing.B) {
	benchmarkArray_Get(1, b)
}

func BenchmarkArray_Get2(b *testing.B) {
	benchmarkArray_Get(2, b)
}

func BenchmarkArray_Get4(b *testing.B) {
	benchmarkArray_Get(4, b)
}

func BenchmarkArray_Get8(b *testing.B) {
	benchmarkArray_Get(8, b)
}

func BenchmarkArray_Get16(b *testing.B) {
	benchmarkArray_Get(16, b)
}

func BenchmarkArray_Get32(b *testing.B) {
	benchmarkArray_Get(32, b)
}

func BenchmarkArray_Get64(b *testing.B) {
	benchmarkArray_Get(64, b)
}

func BenchmarkArray_Get128(b *testing.B) {
	benchmarkArray_Get(128, b)
}

func BenchmarkArray_Get256(b *testing.B) {
	benchmarkArray_Get(256, b)
}

func BenchmarkArray_Get512(b *testing.B) {
	benchmarkArray_Get(512, b)
}

func BenchmarkArray_Get1024(b *testing.B) {
	benchmarkArray_Get(1024, b)
}

func BenchmarkArray_Get2048(b *testing.B) {
	benchmarkArray_Get(2048, b)
}

func benchmarkArray_Set(size int, b *testing.B) {
	var array = CreateArray(size)
	for n := 0; n < b.N; n++ {
		array.Set(size - 1, 1)
	}
}

func BenchmarkArray_Set1(b *testing.B) {
	benchmarkArray_Set(1, b)
}

func BenchmarkArray_Set2(b *testing.B) {
	benchmarkArray_Set(2, b)
}

func BenchmarkArray_Set4(b *testing.B) {
	benchmarkArray_Set(4, b)
}

func BenchmarkArray_Set8(b *testing.B) {
	benchmarkArray_Set(8, b)
}

func BenchmarkArray_Set16(b *testing.B) {
	benchmarkArray_Set(16, b)
}

func BenchmarkArray_Set32(b *testing.B) {
	benchmarkArray_Set(32, b)
}

func BenchmarkArray_Set64(b *testing.B) {
	benchmarkArray_Set(64, b)
}

func BenchmarkArray_Set128(b *testing.B) {
	benchmarkArray_Set(128, b)
}

func BenchmarkArray_Set256(b *testing.B) {
	benchmarkArray_Set(256, b)
}

func BenchmarkArray_Set512(b *testing.B) {
	benchmarkArray_Set(512, b)
}

func BenchmarkArray_Set1024(b *testing.B) {
	benchmarkArray_Set(1024, b)
}

func BenchmarkArray_Set2048(b *testing.B) {
	benchmarkArray_Set(2048, b)
}

func benchmarkArray_Create(size int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		var _ = CreateArray(size)
	}
}

func BenchmarkArray_Create1(b *testing.B) {
	benchmarkArray_Create(1, b)
}

func BenchmarkArray_Create2(b *testing.B) {
	benchmarkArray_Create(2, b)
}

func BenchmarkArray_Create4(b *testing.B) {
	benchmarkArray_Create(4, b)
}

func BenchmarkArray_Create8(b *testing.B) {
	benchmarkArray_Create(8, b)
}

func BenchmarkArray_Create16(b *testing.B) {
	benchmarkArray_Create(16, b)
}

func BenchmarkArray_Create32(b *testing.B) {
	benchmarkArray_Create(32, b)
}

func BenchmarkArray_Create64(b *testing.B) {
	benchmarkArray_Create(64, b)
}

func BenchmarkArray_Create128(b *testing.B) {
	benchmarkArray_Create(128, b)
}

func BenchmarkArray_Create256(b *testing.B) {
	benchmarkArray_Create(256, b)
}

func BenchmarkArray_Create512(b *testing.B) {
	benchmarkArray_Create(512, b)
}

func BenchmarkArray_Create1024(b *testing.B) {
	benchmarkArray_Create(1024, b)
}

func BenchmarkArray_Create2048(b *testing.B) {
	benchmarkArray_Create(2048, b)
}