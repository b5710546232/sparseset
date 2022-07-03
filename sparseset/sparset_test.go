package sparseset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemove(t *testing.T) {
	ss := NewSparseSet[int](10)
	ss.Put(1, 9991)
	ss.Put(2, 21)
	ss.Put(3, 21)
	ss.Put(3, 0)
	ss.Remove(2)

	assert.Equal(t, ss.n, 2)
	val1, _ := ss.Get(1)
	assert.Equal(t, val1, 9991)
	val3, _ := ss.Get(3)
	assert.Equal(t, val3, 0)
}

func TestAddStr(t *testing.T) {
	ss := NewSparseSet[string](1)
	ss.Put(1, "9991")
	ss.Put(2, "21")
	ss.Put(3, "21")
	ss.Put(3, "-1")

	assert.Equal(t, ss.n, 3)
	val1, _ := ss.Get(1)
	assert.Equal(t, "9991", val1)
	val3, _ := ss.Get(3)
	assert.Equal(t, val3, "-1")
}

func TestAddInt(t *testing.T) {
	n := 10_000
	ss := NewSparseSet[int](10)
	for i := 1; i <= n; i++ {
		ss.Put(uint(i), i)
	}

	assert.Equal(t, ss.n, 10_000)
	val, _ := ss.Get(999)
	assert.Equal(t, 999, val)
	val, _ = ss.Get(9999)
	assert.Equal(t, 9999, val)
	val, _ = ss.Get(11_111)
	assert.NotEqual(t, 11_111, val)
	val, _ = ss.Get(1)
	assert.Equal(t, 1, val)
	_, ok := ss.Get(0)
	assert.Equal(t, false, ok)
}
