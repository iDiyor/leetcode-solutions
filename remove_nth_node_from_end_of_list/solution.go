package removenthnodefromendoflist

/*
Given the head of a linked list, remove the nth node from the end of the list and return its head.

Example 1:
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

Example 2:
Input: head = [1], n = 1
Output: []

Example 3:
Input: head = [1,2], n = 1
Output: [1]

Constraints:
The number of nodes in the list is sz.
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz


Follow up: Could you do this in one pass?
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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil && n == 1 {
		return nil
	}

	afterNode := head
	beforeNode := head

	length := getLength(head)

	if n == length {
		return head.Next
	}

	currIdx := 0

	for currIdx < length-n-1 {
		beforeNode = beforeNode.Next
		currIdx++
	}

	currIdx = 0

	for currIdx < length-n+1 {
		afterNode = afterNode.Next
		currIdx++
	}

	beforeNode.Next = afterNode
	return head
}

func getLength(head *ListNode) int {
	var count int
	curr := head
	for curr != nil {
		count++
		curr = curr.Next
	}

	return count
}
