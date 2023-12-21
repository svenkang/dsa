package util

import "testing"

func Test(t *testing.T) {
	s := Stack{}

	if s.Len() != 0 {
		t.Errorf("Length of the stack must have been 0")
	}
	if s.Peek() != nil {
		t.Errorf("Peek value of the stack must have been nil")
	}

	s.Push(1)
	if s.Len() != 1 {
		t.Errorf("Length of the stack must have been 1")
	}
	if s.Peek() != 1 {
		t.Errorf("Peek value of the stack must have been 1")
	}
	if s.Pop() != 1 {
		t.Errorf("Pop value of the stack must have been 1")
	}
	if s.Len() != 0 {
		t.Errorf("Length of the stack must have been 0")
	}

	s.Push(2)
	s.Push(3)
	if s.Len() != 2 {
		t.Errorf("Length of the stack must have been 2")
	}
	if s.Peek() != 3 {
		t.Errorf("Peek value of the stack must have been 3")
	}
}
