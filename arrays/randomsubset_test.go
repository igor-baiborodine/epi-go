package arrays

import (
	"fmt"
	"testing"
)

func TestRandomSubset(t *testing.T) {
	for _, test := range randomSubsetTestData {
		got := RandomSubset(test.n, test.k)
		fmt.Println("n:,", test.n, ", k", test.k, ", got:", got)
	}
}

var randomSubsetTestData = []struct {
	n int
	k int
}{
	{5, 0},
	{5, 1},
	{5, 2},
	{5, 3},
	{5, 4},
	{5, 5},
	{10, 1},
	{10, 2},
	{10, 3},
	{10, 4},
	{10, 5},
	{10, 6},
	{10, 7},
	{10, 8},
	{10, 9},
	{10, 10},
}
