package problems

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1) // 缓存是 {1=1}
	cache.Put(2, 2) // 缓存是 {1=1, 2=2}
	if cache.Get(1) != 1 {
		t.Fatal("get 1")
	}
	cache.Put(3, 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	if get := cache.Get(2); get != -1 {
		t.Fatalf("get %d, want %d", get, -1)
	}
	cache.Put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	if get := cache.Get(1); get != -1 {
		t.Fatalf("get %d, want %d", get, -1)
	}
	if get := cache.Get(3); get != 3 {
		t.Fatalf("get %d, want %d", get, 3)
	}
	if get := cache.Get(4); get != 4 {
		t.Fatalf("get %d, want %d", get, 4)
	}
}
