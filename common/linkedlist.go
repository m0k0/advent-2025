package common

import "iter"

type LinkedList[T any] struct {
	first *linkedListEntry[T]
	last  *linkedListEntry[T]
}
type linkedListEntry[T any] struct {
	list  *LinkedList[T]
	prev  *linkedListEntry[T]
	next  *linkedListEntry[T]
	Value T
}

func (list *LinkedList[T]) Add(value T) {

	entry := linkedListEntry[T]{
		list:  list,
		Value: value,
	}

	if list.first == nil {
		list.first = &entry
	}
	if list.last != nil {
		list.last.next = &entry
		entry.prev = list.last
	}
	list.last = &entry
}

func (entry *linkedListEntry[T]) Remove() {
	if entry.prev != nil {
		entry.prev.next = entry.next
	} else {
		entry.list.first = entry
	}
	if entry.next != nil {
		entry.next.prev = entry.prev
	} else {
		entry.list.last = entry
	}
}

func (list *LinkedList[T]) Entries() iter.Seq[*linkedListEntry[T]] {
	return func(yield func(*linkedListEntry[T]) bool) {

		entry := list.first
		for {
			if entry == nil {
				return
			}
			if !yield(entry) {
				return
			}
			entry = entry.next
		}
	}
}
func (list *LinkedList[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {

		for entry := range list.Entries() {
			if !yield(entry.Value) {
				return
			}
		}
	}
}
