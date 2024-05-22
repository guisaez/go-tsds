package safequeue

import (
	"sync"
)

type SafeQueue[T comparable] struct {
	mu sync.RWMutex
	cond sync.Cond
	// element of the queue (added to the channel)
	elements []T
	// queue size
	size uint64
}

// Initializes the SafeQueue
func NewSafeQueue[T comparable]() *SafeQueue[T] {
	queue := &SafeQueue[T]{}
	queue.cond = *sync.NewCond(&queue.mu)
	return queue
}

// Checks if the queue is empty.
func (sq *SafeQueue[T]) IsEmpty() bool {
	sq.mu.Lock()
	defer sq.mu.Unlock()
	return len(sq.elements) == 0
}

// Returns the current size of the queue.
func (sq *SafeQueue[T]) Size() uint64 {
	sq.mu.Lock()
	defer sq.mu.Unlock()
	return sq.size
}

// Adds an element to the queue and signals any waiting goroutines.
func (sq *SafeQueue[T]) Enqueue(element T) {
	sq.mu.Lock()
	defer sq.mu.Unlock()
	sq.elements = append(sq.elements, element)
	sq.size++
	sq.cond.Signal()
}

// Removes and returns an element from the queue. If the queue is empty, it waits for an element to be enqueued.
func (sq *SafeQueue[T]) DequeueWait() T {
	sq.mu.Lock()
	defer sq.mu.Unlock()
	for len(sq.elements) == 0 {
		sq.cond.Wait()
	}
	element := sq.elements[0]
	sq.elements = sq.elements[1:]
	sq.size--
	return element
}

// Removes and returns an element from the queue. 
func (sq *SafeQueue[T]) Dequeue() (T, bool) {
	sq.mu.Lock()
	defer sq.mu.Unlock()
	if len(sq.elements) == 0{
		var z T
		return z, false
	}
	element := sq.elements[0]
	sq.elements = sq.elements[1:]
	sq.size--
	return element, true
}

// Returns the first element without removing it. Waits if the queue is empty.
func (sq *SafeQueue[T]) PeekWait() T {
	sq.mu.Lock()
	defer sq.mu.Unlock()
	for len(sq.elements) == 0 {
		sq.cond.Wait()
	}
	return sq.elements[0]
}

// Returns the first element without removing it.
func (sq *SafeQueue[T]) Peek() (T, bool) {
	sq.mu.Lock()
	defer sq.mu.Unlock()
	if len(sq.elements) == 0 {
		var z T
		return z, false
	}
	return sq.elements[0], true
}