package sparseset

type SparseSet[T any] struct {
	denseMap []uint
	dense    []T
	sparse   []uint
}

func NewSparseSet[T any](initialSize uint) *SparseSet[T] {
	return &SparseSet[T]{
		sparse: make([]uint, initialSize+1),
	}
}

const emptySlot = 0

func (s *SparseSet[T]) PutVal(id uint, val T) bool {
	return s.Put(id, &val)
}

func (s *SparseSet[T]) newMaxValue(maxValue uint) {
	if len(s.sparse) >= int(maxValue+1) {
		panic("only increasing is possible")
	}
	newSparse := make([]uint, maxValue+1)
	copy(newSparse, s.sparse[:])
	s.sparse = newSparse
}

func (s *SparseSet[T]) Put(id uint, val *T) bool {
	if int(id) >= len(s.sparse) {
		newMaxValue := uint(len(s.sparse) * 2)
		if newMaxValue < id {
			newMaxValue = id + 1
		}
		s.newMaxValue(newMaxValue)
	}

	s.dense = append(s.dense, *val)
	s.denseMap = append(s.denseMap, id)
	s.sparse[id] = uint(len(s.dense))
	return true
}

func (s *SparseSet[T]) Get(id uint) *T {
	if int(id) >= len(s.sparse) {
		return nil
	}
	idx := s.sparse[id]
	if idx == 0 || int(idx) > len(s.dense) {
		return nil
	}
	return &s.dense[idx-1]
}

func (s *SparseSet[T]) Contains(id uint) bool {
	return s.Get(id) != nil
}

func (s *SparseSet[T]) Remove(id uint) {
	if s.Get(id) == nil {
		return
	}

	idx := s.sparse[id]
	s.sparse[id] = 0
	last := s.dense[len(s.dense)-1]
	lastSparse := s.denseMap[len(s.denseMap)-1]

	if int(idx) < len(s.dense) {
		s.dense[idx-1] = last
		s.denseMap[idx-1] = lastSparse
		s.sparse[lastSparse] = idx
	}

	s.dense = s.dense[:len(s.dense)-1]
	s.denseMap = s.denseMap[:len(s.denseMap)-1]
}

func (s *SparseSet[T]) Clear() {
	s.dense = s.dense[:0]
	s.denseMap = s.denseMap[:0]
	s.sparse = make([]uint, len(s.sparse))
}

func (s *SparseSet[T]) Values() []T {
	return s.dense
}
