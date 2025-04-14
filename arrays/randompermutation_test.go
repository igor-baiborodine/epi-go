package arrays

import (
	"fmt"
	"testing"
)

func TestComputeRandomPermutation(t *testing.T) {
	for _, test := range computeRandomPermutationTestData {
		got := ComputeRandomPermutation(test.n)
		fmt.Println("n:,", test.n, ", got:", got)
	}
}

var computeRandomPermutationTestData = []struct {
	n int
}{
	{1},
	{2},
	{3},
	{4},
	{5},
	{6},
}
