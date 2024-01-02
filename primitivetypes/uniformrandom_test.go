package primitivetypes

import (
	"fmt"
	"math"
	"testing"
)

func TestUniformRandom(t *testing.T) {
	for _, test := range UniformRandomTestData {
		var seq []uint64
		for i := 0; i < 10; i++ {
			seq = append(seq, UniformRandom(test.min, test.max))
		}
		fmt.Printf("Generated sequence: %v\n", seq)
	}
}

func BenchmarkUniformRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UniformRandom(math.MaxInt16, math.MaxInt64)
	}
}

var UniformRandomTestData = []struct {
	min uint64
	max uint64
}{
	{0, 10},
	{10, 25},
	{1, 25},
	{999, 1000},
	{0, 1},
}
