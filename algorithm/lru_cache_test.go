package algorithm

import "testing"

type Assert struct {
	T        *testing.T
	CaseName string
}

func TestLRUCache(t *testing.T) {
	cache := NewLRUCache[int, int](2)

	cache.Put(1, 1)
	cache.Put(2, 2)

}
