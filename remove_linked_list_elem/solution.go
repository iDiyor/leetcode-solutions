package removelinkedlistelem

/*
Given the head of a linked list and an integer val,
remove all the nodes of the linked list that has Node.val == val,
and return the new head.

Example 1:
Input: head = [1,2,6,3,4,5,6], val = 6
Output: [1,2,3,4,5]

Example 2:
Input: head = [], val = 1
Output: []

Example 3:
Input: head = [7,7,7,7], val = 7
Output: []

Constraints:
The number of nodes in the list is in the range [0, 104].
1 <= Node.val <= 50
0 <= val <= 50


*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	curr := head

	valCount, nodeCount := 0, 0

	for curr != nil {
		if curr.Val == val {
			valCount++
		}
		nodeCount++
		curr = curr.Next
	}

	if valCount == nodeCount {
		return nil
	}

	// skip vals from head
	for head.Val == val {
		head = head.Next
	}

	curr = head

	for curr != nil {
		if curr.Val == val {
			if prev != nil {
				prev.Next = curr.Next
			}
		} else {
			prev = curr
		}
		curr = curr.Next
	}

	return head
}
