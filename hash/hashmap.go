package ds

// hashmap
// get -- search
// set -- insert
// del -- delete

// linked list
// insert
// search
// delete
type HashMap struct {
	size uint
	arr []LinkedList
}

type LinkedList struct {
	head *node
	size uint
}

type node struct {
	next *node
	value string
}

func (h *HashMap) New(size uint) {
    s := make([]LinkedList, size)
	h = &HashMap{ arr: s, size: size }
}

func (h *HashMap) Set(key string) {
	idx := hash(key, h.size)
	h.arr[idx].Insert(key)
}

func (h *HashMap) Search(key string) bool {
	idx := hash(key, h.size)
	return h.arr[idx].Search(key)
}

func (h *HashMap) Delete(key string) {
	idx := hash(key, h.size)
	h.arr[idx].Delete(key)
}

func hash(key string, size uint) int {
	sum := 0
	for _, char := range key {
		sum += int(char)
	}
	return sum % int(size)
}

func (l *LinkedList) Insert(v string) string {
	n := &node{value: v, next: nil}
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

func (l *LinkedList) Search(v string) bool {
	if l.head == nil {
		return false
	}

	if l.head.value == v {
		return true
	} else {
		ptr := l.head
		for ptr.next != nil {
			if ptr.next.value == v {
				return true
			}
			ptr = ptr.next
		}
	}
	return false
}

func (l *LinkedList) Delete(v string) {
	if l.head == nil {
		return
	}

	if l.head.value == v {
		l.head = nil
		l.size--
	} else {
		ptr := l.head
		for ptr.next != nil {
			if ptr.next.value == v {
				ptr.next = ptr.next.next
				l.size--
				return
			}
			ptr = ptr.next
		}
	}
}
