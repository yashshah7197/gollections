package list

import (
	"github.com/yashshah7197/gollections"
	"github.com/yashshah7197/gollections/iter"
)

// List represents an ordered collection of elements.
type List[T any] interface {
	// Add appends an element to the end of this list.
	Add(T) bool
	// AddAt inserts the specified element at the specified position in this list.
	AddAt(int, T)
	// AddAll appends all the elements in the specified collection to the end of this list,
	// in the order that they are returned by the specified collection's iterator.
	AddAll(gollections.Collection[T]) bool
	// Clear removes all the elements from this list.
	Clear()
	// Contains returns true if this list contains the specified element.
	Contains(T) bool
	// ContainsAll returns true if this list contains all the elements of the specified collection.
	ContainsAll(gollections.Collection[T]) bool
	// Equals returns true if this list is equal to the specified collection.
	Equals(gollections.Collection[T]) bool
	// Get returns the element at the specified position in this list.
	Get(int) T
	// IndexOf returns the index of the first occurrence of the specified element in this list,
	// or -1 if this list does not contain the element.
	IndexOf(T) int
	// IsEmpty returns true if this list contains no elements.
	IsEmpty() bool
	// Iterator returns an iterator over the elements in this list in proper sequence.
	Iterator() iter.Iterator[T]
	// LastIndexOf returns the index of the last occurrence of the specified element in this list,
	// or -1 if this list does not contain the element.
	LastIndexOf(T) int
	// Remove removes the first occurrence of the specified element from this list,
	// if it is present.
	Remove(T) bool
	// RemoveAt removes the element at the specified position in this list.
	RemoveAt(int) T
	// RemoveAll removes all the elements in this list that are also contained in the specified
	// collection.
	RemoveAll(gollections.Collection[T]) bool
	// RetainAll retains only the elements in this list that are contained in the specified
	// collection.
	RetainAll(gollections.Collection[T]) bool
	// Set replaces the element at the specified position in this list with the specified element.
	Set(int, T) T
	// Size returns the number of elements in this list.
	Size() int
	// SubList returns a view of the portion of this list between the specified fromIndex,
	// inclusive, and toIndex, exclusive.
	SubList(int, int) List[T]
	// ToSlice returns a slice containing all the elements in this list in proper sequence.
	ToSlice() []T
}
