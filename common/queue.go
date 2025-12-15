package common

import "iter"

type Queue[T any] struct {
	firstElement *queueElement[T]
	lastElement  *queueElement[T]
	Length       int
}

type queueElement[T any] struct {
	value       T
	nextElement *queueElement[T]
}

func (queue *Queue[T]) Push(value T) {

	element := queueElement[T]{
		value:       value,
		nextElement: nil,
	}
	if queue.firstElement == nil {
		queue.firstElement = &element
	}
	if queue.lastElement != nil {
		queue.lastElement.nextElement = &element
	}
	queue.Length++
	queue.lastElement = &element
}
func (queue *Queue[T]) Pop() (value T) {

	var elementValue T = queue.Peek(0)
	if queue.firstElement != nil {
		queue.firstElement = queue.firstElement.nextElement
		queue.Length--
	}
	return elementValue
}
func (queue *Queue[T]) Peek(offset int) (value T) {
	var elementValue T

	element := queue.firstElement
	for range offset {
		if element == nil {
			break
		}
		element = element.nextElement
	}
	if element != nil {
		elementValue = element.value
	}
	return elementValue
}
func (queue *Queue[T]) Items() iter.Seq[T] {
	return func(yield func(T) bool) {
		element := queue.firstElement
		for range queue.Length {
			if element == nil {
				return
			}
			if !yield(element.value) {
				return
			}

			element = element.nextElement
		}
	}

}
