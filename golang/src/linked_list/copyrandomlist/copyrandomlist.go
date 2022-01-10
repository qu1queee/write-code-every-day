package copyrandomlist

//Definition for a Node
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {

	originalList := head

	nodeMap := map[*Node]*Node{}

	copyList := &Node{}

	// list that would be using to iterate
	fakeList := copyList

	for originalList != nil {
		// provisional Node
		n := &Node{Val: originalList.Val}
		nodeMap[originalList] = n
		fakeList.Next = n
		fakeList = fakeList.Next
		originalList = originalList.Next
	}

	// copy random pointers
	originalList, fakeList = head, copyList
	for originalList != nil {
		fakeList.Next.Random = nodeMap[originalList.Random]
		fakeList = fakeList.Next
		originalList = originalList.Next
	}

	return copyList.Next
}
