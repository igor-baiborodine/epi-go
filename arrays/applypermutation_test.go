package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestApplyPermutations(t *testing.T) {
	for _, test := range ApplyPermutationTestData {
		fmt.Println("Values in test.s before permutation:", test.s)
		ApplyPermutation(test.p, test.s)

		if !reflect.DeepEqual(test.s, test.want) {
			t.Errorf("ApplyPermutation(%d, %v) = %d; want %d", test.p, nil, test.s, test.want)
		}
	}
}

var ApplyPermutationTestData = []struct {
	p    []int
	s    []int
	want []int
}{
	{[]int{}, []int{}, []int{}},
	{[]int{0}, []int{9}, []int{9}},
	{[]int{0, 1}, []int{0, 1}, []int{0, 1}},
	{[]int{1, 0, 2}, []int{2, 0, 1}, []int{0, 2, 1}},
	{[]int{3, 4, 0, 2, 1, 5}, []int{2, 1, 5, 0, 3, 4}, []int{5, 3, 0, 2, 1, 4}},
	{[]int{3, 2, 1, 0, 4}, []int{2, 1, 0, 3, 4}, []int{3, 0, 1, 2, 4}},
	{[]int{0, 4, 5, 1, 3, 2}, []int{3, 0, 5, 2, 4, 1}, []int{3, 2, 1, 4, 0, 5}},
	{[]int{8, 5, 2, 6, 1, 9, 4, 3, 7, 0}, []int{4, 5, 6, 0, 7, 3, 9, 8, 2, 1}, []int{1, 7, 6, 8, 9, 5, 0, 2, 4, 3}},
	{[]int{1, 5, 0, 2, 3, 4, 6}, []int{2, 5, 4, 3, 1, 6, 0}, []int{4, 2, 3, 1, 6, 5, 0}},
	{[]int{2, 8, 4, 7, 11, 10, 3, 1, 5, 6, 0, 9}, []int{11, 9, 4, 1, 10, 2, 8, 6, 0, 3, 5, 7}, []int{5, 6, 11, 8, 4, 0, 3, 1, 9, 7, 2, 10}},
}
