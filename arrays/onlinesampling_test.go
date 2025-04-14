package arrays

import (
	"fmt"
	"testing"

	"github.com/ipfs/boxo/routing/http/types/iter"
)

func TestOnlineRandomSampling(t *testing.T) {
	for _, test := range onlineRandomSamplingTestData {
		fmt.Print("k:", test.k, ", s:", test.s)
		got := OnlineRandomSampling(test.k, iter.FromSlice(test.s))
		fmt.Println(", got:", got)
	}
}

var onlineRandomSamplingTestData = []struct {
	k int
	s []int
}{
	{1, []int{1, 2}},
	{1, []int{1, 2, 3, 4}},
	{2, []int{1, 2, 3, 4}},
	{3, []int{1, 2, 3, 4}},
	{4, []int{1, 2, 3, 4}},
	{1, []int{400, 200, 300}},
	{2, []int{400, 200, 300}},
	{1, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	{2, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
}
