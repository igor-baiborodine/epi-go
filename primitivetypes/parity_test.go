package primitivetypes

import (
	"math"
	"testing"
)

const (
	even = iota
	odd
)

type parityFn func(x uint64) uint16

func testParityFn(t *testing.T, fn parityFn, fnName string) {
	for _, test := range testData {
		if got := fn(test.in); got != test.want {
			t.Errorf("%s(%.64b) = %d; want %d", fnName, test.in, got, test.want)
		}
	}
}

func benchParityFn(b *testing.B, fn parityFn) {
	for i := 0; i < b.N; i++ {
		fn(math.MaxUint64)
	}
}

func TestParityBruteForce(t *testing.T) { testParityFn(t, ParityBruteForce, "ParityBruteForce") }

func BenchmarkParity(b *testing.B) { benchParityFn(b, ParityBruteForce) }

var testData = []struct {
	in   uint64
	want uint16
}{
	{568184680, odd},
	{56, odd},
	{5443530242318502, odd},
	{29469493683482, odd},
	{55042960196257, even},
	{31757824034208103, even},
	{31564787, odd},
	{168618482290043, even},
	{5794530425053671, even},
	{161247358015, odd},
	{855, odd},
	{4184, even},
	{30617518, even},
	{3646550576974, even},
	{8132, even},
	{24652236723, odd},
	{518056968971782, even},
	{363771602569, even},
	{3741538515502, even},
	{math.MaxUint64, even},
}
