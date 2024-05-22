## SafeQueue

A queue is a data structure that follows a specific pattern everytime an element is added or removed. This patters is know as First-In-First-Out (FIFO). It means that the first element added to the queue will be the first element to be taken out. A SafeQueue follows the same behavior, it is just thread safe.

SafeQueues allos mutliple theards or goroutines to access and manipulate the queue simultaneously and safely. This implementation in particular uses lock to ensure all queue operations are atomic and goroutines are synchronized.

### Usage 
```go
queue := NewSafeQueue[int]()
```

### Enqueue
```go
queue.Enqueue(10)
queue.Enqueue(100)
```

### Dequeue
```go
val, ok := queue.Dequeue() // 10, true
```

### DequeueAndWait
```go
val := queue.DequeueWait()
```

### Peek
```go
val, ok := queue.Peek()
```

### PeekWait
```go
val := queue.PeekWait()
```

### IsEmpty
```
ok := queue.IsEmpty()
```

### Size 
```go
size := queue.Size()
```