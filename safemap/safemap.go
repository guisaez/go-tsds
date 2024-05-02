package safemap

import "sync"

type SafeMap[K comparable, V any] struct {
	mx sync.RWMutex
	map_ map[K]V
}

// Creates a new instance of SafeMap with the specified K, V types.
// Returns a pointer to the new SafeMap instance.
func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		map_: make(map[K]V),
	}
}

// Returns the number of key-value pairs stored in the SafeMap
func (sm *SafeMap[K, V]) Len() int {
	sm.mx.Lock()
	defer sm.mx.Unlock()
	return len(sm.map_)
}

// Adds / updates a key-value pair in the SafeMap
func (sm *SafeMap[K, V]) Set(k K, v V) {
	sm.mx.Lock()
	defer sm.mx.Unlock()
	sm.map_[k] = v
}

// Retrieves the value associated with the specified key from the SafeMap.
// Returns the value associated with the key and a boolean indicating whether the key exists in the SafeMap.
func (sm *SafeMap[K, V]) Get(k K) (V, bool) {
	sm.mx.RLock()
	defer sm.mx.RUnlock()
	val, ok := sm.map_[k]
	return val, ok
}

// Deletes a key-value pair from the SafeMap
func (sm *SafeMap[K, V]) Delete(k K) {
	sm.mx.Lock()
	defer sm.mx.Unlock()
	delete(sm.map_, k)
}

// Executes a function fn(K,V) for each key-value pair in the SafeMap.
func (sm *SafeMap[K, V]) ForEach(fn func(key K, val V)) {
	sm.mx.RLock()
	defer sm.mx.RUnlock()
	for k, v := range sm.map_ {
		fn(k, v)
	}
}