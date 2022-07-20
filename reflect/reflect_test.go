package reflect

import "testing"

func TestIsIntType(t *testing.T) {
	if IsIntType("") {
		t.Fatal("string not int")
	}

	if IsIntType(1.111) {
		t.Fatal("float not int")
	}

	if !IsIntType(uint64(12)) {
		t.Fatal("uint is int")
	}

	if IsIntType(make(map[int]int)) {
		t.Fatal("map not int")
	}
	var a int32
	if IsIntType(&a) {
		t.Fatal("pointer not int")
	}
}

func BenchmarkIsIntType(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	// 设置并发数
	b.SetParallelism(5000)
	// 测试多线程并发模式
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			IsIntType(12)
		}
	})
}
