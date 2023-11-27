package iter

// Iterable is an interface for collections that can be iterated over.
type Iterable[T any] interface {
	// Iterator returns an iterator over elements of type T.
	Iterator() Iterator[T]
}
