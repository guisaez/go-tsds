## SafeMap

In Go, maps are not safe for concurrent access by default; modifying a map from multiple goroutines simultaneously can lead to race conditions and undefined behavior.

SafeMap is a data structure designed to provide safe concurrent access to a Map. It provides synchronization mechanisms to ensure that concurrent access to the underlying map is safe.

### Usage 
```go
safemap := NewSafeMap[int, string]()
```

This creates a new SafeMap. A SafeMap will need to know the key-value types beforehand. The key has to be of `comparable` and the value can take `any` type.

#### SET A KEY-PAIR
```go
safemap.Set(0, "zero")
safemap.Set(1, "one")
```
This sets / updates a key-value pair in the SafeMap.

#### GET A VALUE BY KEY
```go
one, ok := safemap.Get(1) // "one", true

three, ok := safemap.Get(3) // "", false
```
It will retrieve the value associated with a key from the SafeMap along with a boolean indicating if the value currently exists in the map. If the value does not exists, it will return the zero value for the type of the maps's value.

#### DELETE A KEY-PAIR
```go
safemap.Delete(1)
```
Deletes a key-value pair from the SafeMap.

#### LENGTH
```go
safemap.Len()
```
Returns the number of key value pairs stored in the safemap.

#### FOR-EACH
```go
custom_fn := fun(key int, val string) {
    ...
}

safemap.ForEach(custom_fn)
```
Evaluates a function fn(K,V) for each key-value pair in the SafeMap.

