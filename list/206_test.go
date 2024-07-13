package ds

import "testing"

/*
Given the head of a singly linked list, reverse the list, and return the reversed list.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var revHead *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = revHead
		revHead = head
		head = tmp
	}
	return revHead
}

func TestReverseLinkedList(t *testing.T) {
	var list *ListNode
	rev := ReverseList(list)
	// 0 item
	if rev != nil {
		t.Fatalf("expected nil but got %v", rev)
	}

	// 1 item
	list2 := &ListNode{ Val: 5, Next: nil }
	rev2 := ReverseList(list2)
	if rev2.Val != 5 {
		t.Fatalf("expected 5 but got %d", rev2.Val)
	}

	// 2 item
	secondNode := &ListNode{ Val: 10, Next: nil }
	list3 := &ListNode{ Val: 5, Next: secondNode }
	rev3 := ReverseList(list3)
	if rev3.Val != 10 {
		t.Fatalf("expected 10 but got %d", rev3.Val)
	}
	if rev3.Next.Val != 5 {
		t.Fatalf("expected 5 but got %d", rev3.Val)
	}
	if rev3.Next.Next != nil {
		t.Fatalf("expected nil but got %d", rev3.Val)
	}
}
