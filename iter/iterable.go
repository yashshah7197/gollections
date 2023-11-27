package iter

// Iterable is an interface for collections that can be iterated over.
type Iterable[T any] interface {
	Iterator() Iterator[T]
}
