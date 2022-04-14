package palindromelinkedlist

/*
Given the head of a singly linked list, return true if it is a palindrome.

Example 1:
Input: head = [1,2,2,1]
Output: true

Example 2:
Input: head = [1,2]
Output: false

Constraints:
The number of nodes in the list is in the range [1, 105].
0 <= Node.val <= 9


Follow up: Could you do it in O(n) time and O(1) space?
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

// Solution:
// Specifically, the steps we need to do are:
// 1. Find the end of the first half.
// 2. Reverse the second half.
// 3. Determine whether or not there is a palindrome.
// 4. Restore the list.
// 5. Return the result.

func isPalindrome(head *ListNode) bool {
	middle := getMiddle(head)
	secondHalfStart := reverseList(middle.Next)

	result := true
	p1 := head
	p2 := secondHalfStart

	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}

		p1 = p1.Next
		p2 = p2.Next
	}

	// restore the list
	middle.Next = reverseList(secondHalfStart)

	return result
}

func getMiddle(head *ListNode) *ListNode {
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	var next *ListNode
	curr := head

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
