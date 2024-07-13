package ds

import "testing"

func TestLinkedListLen(t *testing.T) {
	var list = LinkedList[int]{}

	if list.Len() != 0 {
		t.Fatalf("expected list size to be empty on init but got %d", list.Len())
	}
}

func TestLinkedListPrepend(t *testing.T) {
	var list = LinkedList[int]{}

	list.Prepend(6)
	if list.Len() != 1 {
		t.Fatalf("expected list size to be 1 but got %d", list.Len())
	}
	if list.head.value != 6 {
		t.Fatalf("expected first item to be 6 but got %d", list.head.value)
	}

	list.Prepend(9)
	if list.Len() != 2 {
		t.Fatalf("expected list size to be 2 but got %d", list.Len())
	}
	if list.head.value != 9 {
		t.Fatalf("expected first item to be 9 but got %d", list.head.value)
	}
	if list.head.next.value != 6 {
		t.Fatalf("expected first item to be 6 but got %d", list.head.value)
	}
}

func TestLinkedListAppend(t *testing.T) {
	var list = LinkedList[int]{}
	list.Append(6)

	if list.Len() != 1 {
		t.Fatalf("expected list to be of size 1 but got %d", list.Len())
	}
	if list.head.value != 6 {
		t.Fatalf("expected the first item to be 6 but got %d", list.head.value)
	}

	list.Append(9)
	if list.Len() != 2 {
		t.Fatalf("expected list to be of size 2 but got %d", list.Len())
	}
	if list.head.value != 6 {
		t.Fatalf("expected the first item to be 6 but got %d", list.head.value)
	}
	if list.head.next.value != 9 {
		t.Fatalf("expected the first item to be 9 but got %d", list.head.next.value)
	}

	list.Append(4)
	if list.Len() != 3 {
		t.Fatalf("expected list to be of size 3 but got %d", list.Len())
	}
	if list.head.value != 6 {
		t.Fatalf("expected the first item to be 6 but got %d", list.head.value)
	}
	if list.head.next.value != 9 {
		t.Fatalf("expected the first item to be 9 but got %d", list.head.next.value)
	}
	if list.head.next.next.value != 4 {
		t.Fatalf("expected the first item to be 4 but got %d", list.head.next.next.value)
	}
}

func TestLinkedListRemove(t *testing.T) {
	var list = LinkedList[int]{}

	list.Append(6)
	list.Remove(6)
	if list.Len() != 0 {
		t.Fatalf("expected list size to be 0 after removal but got %d", list.Len())
	}
	if list.head != nil {
		t.Fatalf("expected head to be nil but got %v", list.head)
	}
	list.Append(6)
	list.Append(9)
	list.Remove(9)
	if list.Len() != 1 {
		t.Fatalf("expected list size to be 1 after removal but got %d", list.Len())
	}
	if list.head.value != 6 {
		t.Fatalf("expected head value to be 6 but got %d", list.head.value)
	}

	list.Append(9)
	list.Append(4)
	list.Prepend(4)
	// should have 4,6,9,4
	list.Remove(4)
	if list.Len() != 3 {
		t.Fatalf("expected list size to be 3 after removal but got %d", list.Len())
	}
	if list.head.value != 6 {
		t.Fatalf("expected head value to be 6 but got %d", list.head.value)
	}
	if list.head.next.value != 9 {
		t.Fatalf("expected head value to be 9 but got %d", list.head.next.value)
	}

	list.Remove(4)
	if list.Len() != 2 {
		t.Fatalf("expected list size to be 2 after removal but got %d", list.Len())
	}
	if list.head.value != 6 {
		t.Fatalf("expected head value to be 6 but got %d", list.head.value)
	}
	if list.head.next.value != 9 {
		t.Fatalf("expected head value to be 9 but got %d", list.head.next.value)
	}

	list.Remove(10)
	if list.Len() != 2 {
		t.Fatalf("expected list size to be 2 after removal but got %d", list.Len())
	}
	if list.head.value != 6 {
		t.Fatalf("expected head value to be 6 but got %d", list.head.value)
	}
	if list.head.next.value != 9 {
		t.Fatalf("expected head value to be 9 but got %d", list.head.next.value)
	}
}

