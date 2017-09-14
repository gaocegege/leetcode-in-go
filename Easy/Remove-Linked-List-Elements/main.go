package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	aheadNode := &ListNode{
		Val:  -1,
		Next: head,
	}

	firstPointer := aheadNode
	secondPointer := head

	for secondPointer != nil {
		if secondPointer.Val == val {
			firstPointer.Next = secondPointer.Next
			secondPointer = secondPointer.Next
		} else {
			firstPointer = secondPointer
			secondPointer = secondPointer.Next
		}
	}

	return aheadNode.Next
}
