package twosum2

import (
	"testing"

	"github.com/iDiyor/leetcode-solutions/utils"
)

type test struct {
	numbers []int
	target  int
	output  []int
}

func getTests() []test {
	return []test{
		{
			numbers: []int{2, 7, 11, 15},
			target:  9,
			output:  []int{1, 2},
		},
		{
			numbers: []int{2, 3, 4},
			target:  6,
			output:  []int{1, 3},
		},
		{
			numbers: []int{-1, 0},
			target:  -1,
			output:  []int{1, 2},
		},
	}
}

func TestTwoSum(t *testing.T) {
	for _, test := range getTests() {
		result := twoSum(test.numbers, test.target)
		if !utils.IsSliceEqual(test.output, result) {
			t.Errorf("expected %v, got %v", test.output, result)
		}
	}
}

func TestTwoSum2(t *testing.T) {
	for _, test := range getTests() {
		result := twoSum2(test.numbers, test.target)
		if !utils.IsSliceEqual(test.output, result) {
			t.Errorf("expected %v, got %v", test.output, result)
		}
	}
}
