package reverselist

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	var prev *ListNode = nil
	current := head

	for current != nil {

		nxt := current.Next // store next node as we break the relationship further

		current.Next = prev
		prev = current
		current = nxt // restore to the previous next node
	}

	return prev
}
