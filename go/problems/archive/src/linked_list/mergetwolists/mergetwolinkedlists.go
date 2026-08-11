package mergetwolists

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	pointerToList := new(ListNode)

	for ptr := pointerToList; l1 != nil || l2 != nil; ptr = ptr.Next {

		if l1 == nil {
			ptr.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
			continue
		}

		if l2 == nil {
			ptr.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
			continue
		}

		if l1.Val > l2.Val {
			ptr.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		} else {
			ptr.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		}

	}

	return pointerToList.Next
}
