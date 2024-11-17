// deque is a generalization of stacks and queues

package deque

import "errors"

var ErrEmptyDeque = errors.New("queue is empty")

type Deque[T any] struct {
	Items []T
}

func New[T any]() *Deque[T] {
	d := &Deque[T]{
		Items: make([]T, 0), // Empty slice of type T
	}

	return d
}

func (d *Deque[T]) AppendLeft(value T) {
	d.Items = append([]T{value}, d.Items...)
}

func (d *Deque[T]) Pop() T {
	removed := d.Items[len(d.Items)-1]
	d.Items = d.Items[:len(d.Items)-1]
	return removed
}
