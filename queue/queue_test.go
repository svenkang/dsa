package ds

import "testing"

func TestQueue(t *testing.T) {
	q := Queue[int]{}
	if q.Len() != 0 {
		t.Fatalf("expcted empty 0 but got %v", q.Len())
	}

	q.Push(1)
	if q.Len() != 1 {
		t.Fatalf("expcted 1 but got %v", q.Len())
	}
	if q.head.value != 1 {
		t.Fatalf("expcted 1 but got %v", q.head.value)
	}

	q.Push(2)
	if q.Len() != 2 {
		t.Fatalf("expcted 2 but got %v", q.Len())
	}
	if q.head.next.value != 2 {
		t.Fatalf("expcted 2 but got %v", q.head.next.value)
	}

	q.Push(3)
	if q.Len() != 3 {
		t.Fatalf("expcted 3 but got %v", q.Len())
	}
	if q.head.next.next.value != 3 {
		t.Fatalf("expcted 3 but got %v", q.head.next.next.value)
	}

	first := q.Pop()
	if q.Len() != 2 {
		t.Fatalf("expcted 2 but got %v", q.Len())
	}
	if first != 1 {
		t.Fatalf("expcted 1 but got %v", first)
	}
	if q.head.value != 2 {
		t.Fatalf("expcted 2 but got %v", q.head.value)
	}
	if q.head.next.value != 3 {
		t.Fatalf("expcted 3 but got %v", q.head.next.value)
	}

	peek := q.Peek()
	if q.Len() != 2 {
		t.Fatalf("expcted 2 but got %v", q.Len())
	}
	if peek != 2 {
		t.Fatalf("expcted 2 but got %v", peek)
	}
}
