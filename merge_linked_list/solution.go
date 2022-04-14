package mergelinkedlist

/*
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

Example 1:
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

Example 2:
Input: list1 = [], list2 = []
Output: []

Example 3:
Input: list1 = [], list2 = [0]
Output: [0]


Constraints:
The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both list1 and list2 are sorted in non-decreasing order.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func newListNode(val int) *ListNode {
	return &ListNode{
		Val: val,
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newHead := newListNode(-1)
	currPtr := newHead

	ptr1 := list1
	ptr2 := list2

	for ptr1 != nil && ptr2 != nil {
		if ptr1.Val == ptr2.Val {
			currPtr.Next = newListNode(ptr1.Val)
			currPtr = currPtr.Next
			currPtr.Next = newListNode(ptr2.Val)
			ptr1 = ptr1.Next
			ptr2 = ptr2.Next
		} else if ptr1.Val < ptr2.Val {
			currPtr.Next = newListNode(ptr1.Val)
			ptr1 = ptr1.Next
		} else {
			currPtr.Next = newListNode(ptr2.Val)
			ptr2 = ptr2.Next
		}

		currPtr = currPtr.Next
	}

	for ptr1 != nil {
		currPtr.Next = newListNode(ptr1.Val)
		currPtr = currPtr.Next
		ptr1 = ptr1.Next
	}

	for ptr2 != nil {
		currPtr.Next = newListNode(ptr2.Val)
		currPtr = currPtr.Next
		ptr2 = ptr2.Next
	}

	return newHead.Next
}
