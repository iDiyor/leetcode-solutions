package palindromelinkedlist

import (
	"fmt"
	"testing"
)

func createListFromArray(arr []int) *ListNode {
	var head *ListNode
	var curr *ListNode
	for _, val := range arr {
		if head == nil {
			head = &ListNode{
				Val:  val,
				Next: nil,
			}
			curr = head
		} else {
			curr.Next = &ListNode{
				Val: val,
			}
			curr = curr.Next
		}
	}
	return head
}

func (n *ListNode) Print() {
	fmt.Printf("\n####### START #######\n")
	result := "["

	curr := n

	for curr != nil {
		result += fmt.Sprintf("%d,", curr.Val)
		curr = curr.Next
	}

	result += "]"

	fmt.Printf("list -> %s\n", result)
	fmt.Printf("####### END #######")
}

func TestIsPalindrome(t *testing.T) {
	type test struct {
		Head   *ListNode
		Output bool
	}

	tests := []test{
		{
			Head:   createListFromArray([]int{1, 2, 2, 1}),
			Output: true,
		},
		{
			Head:   createListFromArray([]int{1, 2}),
			Output: false,
		},
	}

	for _, test := range tests {
		result := isPalindrome(test.Head)
		if result != test.Output {
			t.Errorf("expected %v, got %v", test.Output, result)
		}
	}
}
