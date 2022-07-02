package brechmark

import (
	"fmt"
	"testing"

	"github.com/b5710546232/sparseset/sparseset"
)

func Benchmark_map_int_write_10_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmark_map_write(10_000)
	}
}

func Benchmark_map_str_write_10_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmark_map_str_write(10_000)
	}
}

func Benchmark_map_int_read_10_000(b *testing.B) {
	n := 10_000
	m := make(map[int]int, n)
	for i := 1; i <= n; i++ {
		m[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmark_map_read(m)
	}
}

func Benchmark_map_str_read_10_000(b *testing.B) {
	n := 10_000
	m := make(map[int]string, n)
	for i := 1; i <= n; i++ {
		m[i] = string(rune(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmark_map_read_str(m)
	}
}

func benchmark_map_write(n int) {
	m := make(map[int]int, n)
	for i := 1; i <= n; i++ {
		m[i] = i
	}
}

func benchmark_map_str_write(n int) {
	m := make(map[int]string, n)
	for i := 1; i <= n; i++ {
		m[i] = fmt.Sprintf("%d", i)
	}
}

func Benchmark_spareset_int_write_10_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmark_spareset_int_write(10_000)
	}
}

func Benchmark_spareset_str_write_10_000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmark_spareset_str_write(10_000)
	}
}

func Benchmark_spareset_read_int_10_000(b *testing.B) {
	n := 10_000
	m := sparseset.NewSparseSet[int](n, n)
	for i := 1; i <= n; i++ {
		m.Put(uint(i), i)
	}
	for i := 0; i < b.N; i++ {
		benchmark_spareset_read(m)
	}
}

func Benchmark_spareset_read_str_10_000(b *testing.B) {
	n := 10_000
	m := sparseset.NewSparseSet[string](n, n)
	for i := 1; i <= n; i++ {
		m.Put(uint(i), string(rune(i)))
	}
	for i := 0; i < b.N; i++ {
		benchmark_spareset_str_read(m)
	}
}

func benchmark_map_read(m map[int]int) {
	for i := range m {
		_ = m[i]
	}
}

func benchmark_map_read_str(m map[int]string) {
	for i := range m {
		_ = m[i]
	}
}

func benchmark_spareset_int_write(n int) {
	m := sparseset.NewSparseSet[int](n, n)
	for i := 1; i <= n; i++ {
		m.Put(uint(i), i)
	}
}

func benchmark_spareset_str_write(n int) {
	m := sparseset.NewSparseSet[string](n, n)
	for i := 1; i <= n; i++ {
		m.Put(uint(i), fmt.Sprintf("%d", i))
	}
}

func benchmark_spareset_read(ss *sparseset.SparseSet[int]) {
	for i := range ss.Dense {
		_ = ss.Dense[i]
	}
}

func benchmark_spareset_str_read(ss *sparseset.SparseSet[string]) {
	for i := range ss.Dense {
		_ = ss.Dense[i]
	}
}