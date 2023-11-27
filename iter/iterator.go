package iter

// Iterator is an interface for iterating over a collection.
type Iterator[T any] interface {
	HasNext() bool
	Next() T
	Remove()
}
