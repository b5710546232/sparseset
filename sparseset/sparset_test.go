package sparseset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemove(t *testing.T) {
	ss := NewSparseSet[int](10, 10)
	ss.Put(1, 9991)
	ss.Put(2, 21)
	ss.Put(3, 21)
	ss.Put(3, 0)
	ss.Remove(2)
	assert.Equal(t, ss.n, 2)
	assert.Equal(t, ss.Dense[0], 9991)
	assert.Equal(t, ss.Dense[1], 0)
}
