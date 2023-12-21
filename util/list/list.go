package util

type (
	node struct {
		value interface{}
		next  *node
	}

	List struct {
		head   *node
		length int
	}
)

// Len()
func (this List) Len() int {
	return this.length
}

// Append()
func (this *List) Append(value interface{}) {
	node := &node{value, nil}

	if this.head == nil {
		this.head = node
	} else {
		current = this.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
	this.length++
}

// Prepend()
func (this *List) Prepend(value interface{}) {
	node := &node{value, this.head}
	this.head = node
	this.length++
}
