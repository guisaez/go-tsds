package safequeue

import (
	"math/rand"
	"testing"
	"time"
)

func TestNewSafeQueueAndIsEmpty(t *testing.T) {

	queue := NewSafeQueue[int]()

	if empty := queue.IsEmpty(); empty != true {
		t.Error("expected the queue to be empty")
	}
}

func TestEnqueueSizeEmpty(t *testing.T) {
	queue := NewSafeQueue[int]()

	queue.Enqueue(10)
	queue.Enqueue(30)

	if size := queue.Size(); size != 2 {
		t.Errorf("expected size to be 2, got %d", size)
	}

	if empty := queue.IsEmpty(); empty != false {
		t.Errorf("expected IsEmpty to return false, got %v", empty)
	}
}

func TestDequeue(t *testing.T) {
	queue := NewSafeQueue[int]()

	if val, ok := queue.Dequeue(); ok != false  {
		t.Errorf("expected val = 0, ok = false, got val = %d, ok = %v", val, ok)
	}
	
	queue.Enqueue(10)
	queue.Enqueue(100)
	queue.Enqueue(1231)

	val, ok := queue.Dequeue()
	if val != 10 || !ok {
		t.Errorf("expected val = 10, ok = true, got val = %d, ok = %v", val, ok)
	}

	if size := queue.Size(); size != 2 {
		t.Errorf("expected size to be 2, got %d", size)
	}
}

func TestPeek(t *testing.T) {
	queue := NewSafeQueue[int]()

	if val, ok := queue.Peek(); ok != false  {
		t.Errorf("expected val = 0, ok = false, got val = %d, ok = %v", val, ok)
	}

	queue.Enqueue(1231)
	queue.Enqueue(100)
	queue.Enqueue(10)

	val, ok := queue.Peek()
	if val != 1231 || ok != true {
		t.Errorf("expected val = 1231, ok = true, got val = %d, ok = %v", val, ok)
	} 

	if size := queue.Size(); size != 3 {
		t.Errorf("expected size to be 3, got %d", size)
	}
}

func TestDequeueWait(t *testing.T) {
	queue := NewSafeQueue[int]()

	dq := func(){
		val := queue.DequeueWait()
		if val != 10 {
			t.Errorf("expected value = 10, got %d", val)
		}
	}

	eq := func(){
		queue.Enqueue(10)
	}

	go dq()
	time.Sleep(time.Second * 1)
	go eq()
}

func TestPeekWait(t *testing.T) {
	queue := NewSafeQueue[int]()

	dq := func(){
		val := queue.PeekWait()
		if val != 10 {
			t.Errorf("expected value = 10, got %d", val)
		}
	}

	eq := func(){
		queue.Enqueue(10)
	}

	go dq()
	time.Sleep(time.Second * 1)
	go eq()
}

func BenchmarkTestConcurrentEnqueue(b *testing.B){

	queue := NewSafeQueue[uint64]()

	b.SetParallelism(100)
	b.ResetTimer()

	b.RunParallel(func (pb *testing.PB) {
		for pb.Next(){
			r := rand.Uint64() % 10000
			queue.Enqueue(r)
		}
	})
}

func BenchmarkTestConcurrentDequeue(b *testing.B){
	queue := NewSafeQueue[uint64]()

	b.SetParallelism(100)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB){
		for pb.Next(){
			r := rand.Uint64() % 100000
			queue.Enqueue(r)
			queue.DequeueWait()
		}
	})
}