package sparseset

import (
	"fmt"
	"math"
)

type SparseSet[T any] struct {
	Sparse    []uint // Stores index, id
	Dense     []T    // Stores the actual data
	n         int    // Current size of dense
	denseSize int
}

const emptySlot = math.MaxUint

func NewSparseSet[T any](initialSparseSize int, initialDenseSize int) *SparseSet[T] {
	ss := &SparseSet[T]{
		Sparse:    make([]uint, initialSparseSize),
		Dense:     make([]T, initialDenseSize),
		denseSize: initialDenseSize,
		n:         0,
	}

	for i := range ss.Sparse {
		ss.Sparse[i] = emptySlot
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
	}
	diffSparseSize := int(idx) - len(s.Sparse) + 1
	if diffSparseSize > 0 {
		for i := 0; i < diffSparseSize; i++ {
			s.Sparse = append(s.Sparse, emptySlot)
		}
	}
	s.Sparse[idx] = idx

	if s.n < len(s.Dense) {
		s.Dense[denseIdx] = val
	} else {
		s.Dense = append(s.Dense, val)
	}

	return true
}

func (s *SparseSet[T]) Get(id uint) T {
	sIdx := id - 1
	idx := s.Sparse[sIdx]
	fmt.Println("idx", idx, s.Dense)
	if idx == emptySlot {
		var t T
		return pass(t)
	}
	return s.Dense[idx]
}

func pass[T any](tp T) T {
	return tp
}

func (s *SparseSet[T]) Remove(id uint) bool {
	if !s.Contains(id) {
		return false
	}

	lastDense := s.Dense[s.n]
	idx := id - 1
	s.Dense[idx] = lastDense
	s.Sparse[idx] = emptySlot

	s.n -= 1
	return true
}

func (s *SparseSet[T]) PrintDense() {
	fmt.Printf("\n----PrintDense----\n")
	for i := 0; i < s.n; i++ {
		fmt.Println("idx", i, "=>", s.Dense[i])
	}
	fmt.Printf("\n--------\n")
}

func (s *SparseSet[T]) PrintSpase() {
	fmt.Println("PrintSpase")
	fmt.Printf("[ ")
	for i := range s.Sparse {
		if s.Sparse[i] == emptySlot {
			fmt.Print(" nil ,")
		} else {
			fmt.Printf("%d ,", s.Sparse[i])
		}

	}
	fmt.Printf(" ]\n")
}

func (s *SparseSet[T]) PrintSpaseIds() {
	fmt.Println("PrintSpase")
	fmt.Printf(" [\n")
	for i := range s.Sparse {
		if s.Sparse[i] != emptySlot {
			fmt.Printf("\tindex-%d => %d \n", i, s.Sparse[i]+1)
		}

	}
	fmt.Printf(" ]\n")
}

func (s *SparseSet[T]) Contains(id uint) bool {
	idx := id - 1
	if int(idx) > len(s.Sparse)-1 {
		return false
	}
	return s.Sparse[idx] != emptySlot
}
