// deque is a generalization of stacks and queues

package deque

import "errors"

var ErrEmptyDeque = errors.New("queue is empty")

type Deque[T any] struct {
	items []T
}

func New[T any]() *Deque[T] {
	d := &Deque[T]{
		items: make([]T, 0), // Empty slice of type T
	}

	return d
}

func (d *Deque[T]) AppendLeft(value T) {
	d.items = append([]T{value}, d.items...)
}

func (d *Deque[T]) Pop() T {
	removed := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]
	return removed
}
