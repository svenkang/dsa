package ds

import (
	"testing"
)

func TestStackPush(t *testing.T) {
	stack := Stack[uint]{}
	var number uint
	number = 1
	stack.Push(number)
	if stack.size != 1 {
		t.Fatalf(`Expected: %d. Actual: %d`, 1, stack.size)
	}
	if stack.top.value != number {
		t.Fatalf(`Expected: %d. Actual: %d`, number, stack.top.value)
	}
	var number2 uint
	number2 = 2
	stack.Push(number2)
	if stack.size != 2 {
		t.Fatalf(`Expected: %d. Actual: %d`, 2, stack.size)
	}
	if stack.top.value != number2 {
		t.Fatalf(`Expected: %d. Actual: %d`, number2, stack.top.value)
	}
}

func TestStackEmptyLen(t *testing.T) {
	stack := Stack[uint]{}
	if stack.Len() != 0 {
		t.Fatalf(`Expected: %d. Actual %d`, 0, stack.Len())
	}
}

func TestStackLen(t *testing.T) {
	stack := Stack[uint]{}
	stack.Push(1)
	if stack.Len() != 1 {
		t.Fatalf(`Expected %d, Actual: %d`, 1, stack.Len())
	}
	stack.Push(1)
	if stack.Len() != 2 {
		t.Fatalf(`Expected %d, Actual: %d`, 2, stack.Len())
	}
}

func TestStackPop(t *testing.T) {
	stack := Stack[uint]{}
	stack.Push(1)
	stack.Push(2)
	p1, err := stack.Pop()
	if p1 != 2 {
		t.Fatalf(`Expected %d, Actual %d`, 2, p1)
	}
	if stack.Len() != 1 {
		t.Fatalf(`Expected %d, Actual %d`, 1, stack.Len())
	}
	if err != nil {
		t.Fatalf(`Expected %v, Actual %v`, nil, err)
	}
	p2, err2 := stack.Pop()
	if p2 != 1 {
		t.Fatalf(`Expected %d, Actual %d`, 1, p1)
	}
	if stack.Len() != 0 {
		t.Fatalf(`Expected %d, Actual %d`, 0, stack.Len())
	}
	if err2 != nil {
		t.Fatalf(`Expected %v, Actual %v`, nil, err2)
	}
	_, err3 := stack.Pop()
	if err3 == nil {
		t.Fatalf(`Expected %v, Actual %v`, err3, nil)
	}
}

func TestStackPeek(t *testing.T) {
	stack := Stack[uint]{}
	stack.Push(1)
	v, err := stack.Peek()
	if v != 1 {
		t.Fatalf("Expected: %v, Actual: %v", 1, v)
	}
	if err != nil {
		t.Fatalf("Expected: %v, Actual: %v", nil, err)
	}
	stack.Push(2)
	v2, err2 := stack.Peek()
	if v2 != 2 {
		t.Fatalf("Expected: %v, Actual: %v", 2, v)
	}
	if err2 != nil {
		t.Fatalf("Expected: %v, Actual: %v", nil, err2)
	}
	if stack.Len() != 2 {
		t.Fatalf("Expected: %v, Actual: %v", 2, stack.Len())
	}
	stack.Pop()
	stack.Pop()
	_, err3 := stack.Pop()
	if err3 == nil {
		t.Fatalf("Expected %v, Actual: %v", EmptyStackErr, nil)
	}
	if stack.Len() != 0 {
		t.Fatalf("Expected %v, Actual: %v", 0, stack.Len())
	}
}
