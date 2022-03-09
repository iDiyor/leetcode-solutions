package removenthnodefromendoflist

import (
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
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

	resultHead := removeNthFromEnd(&head, 2)
	idx := 1
	for idx <= 5 && resultHead != nil {
		if resultHead.Val != idx {
			t.Errorf("expected %d, got %d", idx, resultHead.Val)
		}
		resultHead = resultHead.Next
		idx++
	}
}
