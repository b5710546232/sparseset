package sparseset

import (
	"fmt"
	"math"
)

type SparseSet[T any] struct {
	sparse    []uint // Stores index, id
	dense     []T    // Stores the actual data
	n         int    // Current size of dense
	denseSize int
}

const emptySlot = math.MaxUint

func NewSparseSet[T any](initialSize int) *SparseSet[T] {
	ss := &SparseSet[T]{
		sparse:    make([]uint, initialSize),
		dense:     make([]T, initialSize),
		denseSize: initialSize,
		n:         0,
	}

	for i := range ss.sparse {
		ss.sparse[i] = emptySlot
	}
	return ss
}

func (s *SparseSet[T]) Put(id uint, val T) bool {
	if id <= 0 {
		return false
	}

	idx := id - 1
	denseIdx := s.n
	if !s.Contains(id) {
		s.n += 1
	} else {
		denseIdx = int(s.sparse[idx])
	}
	diffSparseSize := int(idx) - len(s.sparse)
	if diffSparseSize >= 0 {
		for i := 0; i < diffSparseSize+1; i++ {
			s.sparse = append(s.sparse, emptySlot)
		}
	}
	lastcurrentDenseSizeIdx := uint(s.n) - 1
	s.sparse[idx] = lastcurrentDenseSizeIdx
	if s.n <= len(s.dense) {
		s.dense[denseIdx] = val
	} else {
		s.dense = append(s.dense, val)
	}

	return true
}

func (s *SparseSet[T]) Get(id uint) (T, bool) {
	sIdx := int(id - 1)
	if sIdx > len(s.sparse) {
		var t T
		return pass(t), false
	}
	idx := s.sparse[sIdx]
	if idx == emptySlot {
		var t T
		return pass(t), false
	}
	return s.dense[idx], true
}

func pass[T any](tp T) T {
	return tp
}

func (s *SparseSet[T]) Remove(id uint) bool {
	if !s.Contains(id) {
		return false
	}

	lastDense := s.dense[s.n-1]
	idx := id - 1
	s.dense[idx] = lastDense
	s.sparse[idx] = emptySlot

	s.n -= 1
	return true
}

func (s *SparseSet[T]) PrintDense() {
	fmt.Printf("\n----PrintDense----\n")
	for i := 0; i < s.n; i++ {
		fmt.Println("idx", i, "=>", s.dense[i])
	}
	fmt.Printf("\n--------\n")
}

func (s *SparseSet[T]) PrintSpase() {
	fmt.Println("PrintSpase")
	fmt.Printf("[ ")
	for i := range s.sparse {
		if s.sparse[i] == emptySlot {
			fmt.Print(" nil ,")
		} else {
			fmt.Printf("%d ,", s.sparse[i])
		}

	}
	fmt.Printf(" ]\n")
}

func (s *SparseSet[T]) PrintSpaseIds() {
	fmt.Println("PrintSpase")
	fmt.Printf(" [\n")
	for i := range s.sparse {
		if s.sparse[i] != emptySlot {
			fmt.Printf("\tindex-%d => %d \n", i, s.sparse[i]+1)
		}

	}
	fmt.Printf(" ]\n")
}

func (s *SparseSet[T]) Contains(id uint) bool {
	idx := id - 1
	if int(idx) > len(s.sparse)-1 {
		return false
	}
	return s.sparse[idx] != emptySlot
}

func (s *SparseSet[T]) ForEach(f func(int, T)) {
	for i := range s.dense {
		f(i, s.dense[i])
	}
}

func (s *SparseSet[T]) Dense() []T {
	return s.dense
}
