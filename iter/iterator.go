package iter

// Iterator is an interface for iterating over a collection.
type Iterator[T any] interface {
	// HasNext returns true if the iteration has more elements.
	HasNext() bool
	// Next returns the next element in the iteration.
	Next() T
	// Remove removes from the underlying collection the last element returned by this iterator.
	Remove()
}
