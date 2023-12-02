package gollections

import "github.com/yashshah7197/gollections/iter"

// Collection represents a group of elements.
type Collection[T any] interface {
	// Add adds an element to this collection.
	Add(T) bool
	// AddAll adds all elements from another collection to this collection.
	AddAll(Collection[T]) bool
	// Clear removes all elements from this collection.
	Clear()
	// Contains returns true if this collection contains the specified element.
	Contains(T) bool
	// ContainsAll returns true if this collection contains all the elements from the
	// specified collection.
	ContainsAll(Collection[T]) bool
	// Equals returns true if this collection is equal to the specified collection.
	Equals(Collection[T]) bool
	// IsEmpty returns true if this collection contains no elements.
	IsEmpty() bool
	// Iterator returns an iterator over the elements in this collection.
	Iterator() iter.Iterator[T]
	// Remove removes a single instance of the specified element from this collection, if present.
	Remove(T) bool
	// RemoveAll removes all this collection's elements that are also contained in the specified
	// collection.
	RemoveAll(Collection[T]) bool
	// RetainAll retains only the elements in this collection that are contained in the specified
	// collection.
	RetainAll(Collection[T]) bool
	// Size returns the number of elements in this collection.
	Size() int
	// ToSlice returns a slice containing all the elements in this collection.
	ToSlice() []T
}
