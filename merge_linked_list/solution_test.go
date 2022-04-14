package mergelinkedlist

import (
	"fmt"
	"testing"
)

func createListFromArray(arr []int) *ListNode {
	if arr == nil {
		return nil
	}

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

func TestMergeTwoLists(t *testing.T) {
	type test struct {
		list1  *ListNode
		list2  *ListNode
		output *ListNode
	}

	tests := []test{
		{
			list1:  createListFromArray([]int{1, 2, 4}),
			list2:  createListFromArray([]int{1, 3, 4}),
			output: createListFromArray([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			list1:  nil,
			list2:  nil,
			output: nil,
		},
		{
			list1:  nil,
			list2:  createListFromArray([]int{0}),
			output: createListFromArray([]int{0}),
		},
	}

	for idx, test := range tests {
		result := mergeTwoLists(test.list1, test.list2)
		if !isEqual(result, test.output) {
			t.Errorf("wrong result for test #%d", idx)
		}
	}
}

func isEqual(list1 *ListNode, list2 *ListNode) bool {
	if list1 == nil && list2 == nil {
		return true
	}

	curr1 := list1
	curr2 := list2

	for curr1 != nil && curr2 != nil {
		if curr1.Val != curr2.Val {
			return false
		}
		curr1 = curr1.Next
		curr2 = curr2.Next
	}

	return true
}
