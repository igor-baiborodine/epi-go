package primitivetypes

import (
	"math"
	"testing"
)

type reverseDigitsFn func(int) int

func testReverseDigitsFn(t *testing.T, fn reverseDigitsFn) {
	for _, test := range reverseDigitsTestData {
		if got := fn(test.in); got != test.want {
			t.Errorf("%s(%#.16x) = %#.16x; want %#.16x", fnName(fn), test.in, got, test.want)
		}
	}
}

func TestReverseDigitsViaString(t *testing.T) { testReverseDigitsFn(t, ReverseDigitsViaString) }
func TestReverseDigits(t *testing.T)          { testReverseDigitsFn(t, ReverseDigits) }

func benchReverseDigitsFn(b *testing.B, fn reverseDigitsFn) {
	for i := 0; i < b.N; i++ {
		fn(math.MaxInt)
	}
}

func BenchmarkReverseDigitsViaString(b *testing.B) { benchReverseDigitsFn(b, ReverseDigitsViaString) }
func BenchmarkReverseDigits(b *testing.B)          { benchReverseDigitsFn(b, ReverseDigits) }

var reverseDigitsTestData = []struct {
	in   int
	want int
}{
	{1683683571, 1753863861},
	{1799113645, 5463119971},
	{2138559785, 5879558312},
	{-1856396381, -1836936581},
	{1296932912, 2192396921},
	{-778610391, -193016877},
	{-1203840386, -6830483021},
	{1963072368, 8632703691},
	{-363773848, -848377363},
	{-1299988089, -9808899921},
	{388359, 953883},
	{-388359, -953883},
	{9097884, 4887909},
	{-9097884, -4887909},
	{110326, 623011},
	{-110326, -623011},
	{2465, 5642},
	{-2465, -5642},
	{19, 91},
	{-19, -91},
	{9, 9},
	{-9, -9},
	{11050549, 94505011},
	{-11050549, -94505011},
	{1, 1},
	{-1, -1},
	{18, 81},
	{-18, -81},
	{165, 561},
}
