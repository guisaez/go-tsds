package safemap

import (
	"math/rand"
	"testing"
)

func TestNewSafeMap(t *testing.T) {
	sm := NewSafeMap[int, int]()

	if length := sm.Len(); length != 0 {
		t.Errorf("expected length 0, got %d", length)
	}
}

func TestSetAndGet(t *testing.T) {
	sm := NewSafeMap[int, string]()

	sm.Set(1, "one")
	sm.Set(2, "two")

	if length := sm.Len(); length != 2 {
		t.Errorf("expected length 2, got %d", length)
	}

	if val, ok := sm.Get(1); val != "one" || !ok {
		t.Errorf("expected val to be one, got %v, expected ok to be true, got %v", val, ok)
	}

	if _, ok := sm.Get(3); ok {
		t.Errorf("expected false, got %v", ok)
	}
}

func TestDelete(t *testing.T) {
	sm := NewSafeMap[int, string]()

	sm.Set(1, "one")
	sm.Set(2, "two")

	sm.Delete(2)

	if length := sm.Len(); length != 1 {
		t.Errorf("expected length 1, got %d", length)
	}

	sm.Delete(3)
}

func TestForEach(t *testing.T) {
	sm := NewSafeMap[int, string]()

	sm.Set(1, "one")
	sm.Set(2, "two")
	sm.Set(3, "three")

	sum := 0

	sm.ForEach(func(k int, val string) {
		sum += k
	})

	if sum != 6 {
		t.Errorf("expected sum to be 6, got %d", sum)
	}
}

func BenchmarkTestConcurrentGet(b *testing.B) {
	sm := NewSafeMap[uint64, uint64]()

	for i := range 10000000 {
		sm.Set(uint64(i),uint64(i))
	}

	b.SetParallelism(100)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB){
		for pb.Next() {
			r := rand.Uint64() % 1000000
			_,_ = sm.Get(r)
		}
	})
}

var t uint64
func BenchmarkTestConcurrentSet(b *testing.B) {
	sm := NewSafeMap[uint64, uint64]()
	b.SetParallelism(100)
	b.ResetTimer()

	b.RunParallel(func (pb*testing.PB) {
		for pb.Next() {
			r := rand.Uint64() % 10000
			t, _ = sm.Get(r)
			sm.Set(r, r)
			_, _ = sm.Get(r)
		}
	})
}