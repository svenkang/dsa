package ds

type (
	node struct {
		value interface{}
		prev  *node
	}
	Stack struct {
		top    *node
		length int
	}
)

// Len()
func (this *Stack) Len() int {
	return this.length
}

// Peek()
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

// Push()
func (this *Stack) Push(value interface{}) {
	n := &node{value, this.top}
	this.top = n
	this.length++
}

// Pop()
func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}
