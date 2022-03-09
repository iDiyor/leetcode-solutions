package linkedlistcycle

import "testing"

func TestHasCycle(t *testing.T) {
	type test struct {
		head   *ListNode
		output bool
	}

	node := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val:  0,
			Next: nil,
		},
	}
	node.Next.Next = node

	tests := []test{
		{
			head: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val:  -4,
							Next: node,
						},
					},
				},
			},
			output: true,
		},
		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: node,
				},
			},
			output: true,
		},
		{
			head: &ListNode{
				Val:  1,
				Next: nil,
			},
			output: false,
		},
	}

	for _, test := range tests {
		result := hasCycle(test.head)
		if result != test.output {
			t.Errorf("expected %v, got %v", test.output, result)
		}
	}
}
