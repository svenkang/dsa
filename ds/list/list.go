package ds

type (
	node[T comparable] struct {
		value T
		next *node[T]
	}

	LinkedList[T comparable] struct {
		head *node[T]
		size uint
	}
)

func (l LinkedList[T]) Len() uint {
	return l.size
}

func (l *LinkedList[T]) Prepend(v T) T {
	l.head = &node[T]{ value: v, next: l.head }
	l.size++
	return v
}

func (l *LinkedList[T]) Append(v T) T {
	n := &node[T]{value: v, next: nil }
	if l.head == nil {
		l.head = n
	} else {
		ptr := l.head
		for ptr.next != nil {
			ptr = ptr.next
		}
		ptr.next = n
	}
	l.size++
	return v
}


func (l *LinkedList[T]) Remove(v T) {
	if l.head == nil {
		return
	}

	if l.head.value == v {
		l.head = l.head.next
		l.size--
		return
	}

	ptr := l.head
	for ptr.next != nil && ptr.next.value != v {
		ptr = ptr.next
	}
	if ptr.next != nil {
		ptr.next = ptr.next.next
		l.size--
	}
	return
}
