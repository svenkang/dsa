package ds

// FIFO 
type (
	node[T comparable] struct {
		value T
		next *node[T]
	}

	Queue[T comparable] struct {
		head *node[T]
		size uint
	}
)

func (q *Queue[T]) Len() uint {
	return q.size
}

func (q *Queue[T]) Push(v T) {
	n := &node[T]{ value: v, next: nil }
	if q.head == nil {
		q.head = n
	} else {
		tail := q.head
		for tail.next != nil {
			tail = tail.next
		}
		tail.next = n
	}
	q.size++
	return
}

func (q *Queue[T]) Pop() T {
	if q.head == nil {
		return *new(T)
	}
	current := q.head.value
	q.head = q.head.next
	q.size--
	return current
}

func (q *Queue[T]) Peek() T {
	return q.head.value
}
