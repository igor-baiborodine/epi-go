package primitivetypes

import (
	"math"
	"testing"
)

func TestPalindromeNumber(t *testing.T) {
	for _, test := range PalindromeNumberTestData {
		if got := PalindromeNumber(test.x); got != test.want {
			t.Errorf("PalindromeNumber(%d) = %t; want %t", test.x, got, test.want)
		}
	}
}

func BenchmarkPalindromeNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PalindromeNumber(math.MaxInt64)
	}
}

var PalindromeNumberTestData = []struct {
	x    int
	want bool
}{
	{0, true},
	{-1, false},
	{1, true},
	{2982623, false},
	{-2982623, false},
	{14455441, true},
	{1445441, true},
}
