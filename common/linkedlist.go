package common

import "iter"

type LinkedList[T any] struct {
	first *linkedListEntry[T]
	last  *linkedListEntry[T]
}
type linkedListEntry[T any] struct {
	list  *LinkedList[T]
	next  *linkedListEntry[T]
	value T
}

func (list *LinkedList[T]) Add(value T) {

	entry := linkedListEntry[T]{
		list:  list,
		value: value,
	}

	if list.first == nil {
		list.first = &entry
	}
	if list.last != nil {
		list.last.next = &entry
	}
	list.last = &entry
}
func (list *LinkedList[T]) Iterate() iter.Seq[T] {
	return func(yield func(T) bool) {

		entry := list.first
		for {
			if entry == nil {
				return
			}
			if !yield(entry.value) {
				return
			}
			entry = entry.next
		}
	}
}
