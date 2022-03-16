package oddevenlinkedlist

import (
	"log"
	"testing"
)

func TestOddEvenList(t *testing.T) {
	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	expectedVals := []int{1, 3, 5, 2, 4}

	resultHead := oddEvenList(&head)

	curr := resultHead

	for _, val := range expectedVals {
		if curr.Val != val {
			t.Errorf("expected %d, got %d", val, curr.Val)
		}
		log.Printf("curr.Val - %d", curr.Val)
		curr = curr.Next
	}
}
