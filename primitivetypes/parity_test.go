package primitivetypes

import (
	"math"
	"reflect"
	"runtime"
	"testing"
)

const (
	even = iota
	odd
)

type parityFn func(x uint64) uint16

func testParityFn(t *testing.T, fn parityFn) {
	for _, test := range parityTestData {
		if got := fn(test.in); got != test.want {
			t.Errorf("%s(%.64b) = %d; want %d", fnName(fn), test.in, got, test.want)
		}
	}
}

func fnName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func TestParityBruteForce(t *testing.T)         { testParityFn(t, ParityBruteForce) }
func TestParityBruteForceImproved(t *testing.T) { testParityFn(t, ParityBruteForceImproved) }
func TestParity(t *testing.T)                   { testParityFn(t, Parity) }
func TestParityBitsOnesCount(t *testing.T)      { testParityFn(t, ParityBitsOnesCount) }

func benchParityFn(b *testing.B, fn parityFn) {
	for i := 0; i < b.N; i++ {
		fn(uint64(math.Pow(2, 64))) // only most significant bit is set to 1
	}
}

func BenchmarkParityBruteForce(b *testing.B)         { benchParityFn(b, ParityBruteForce) }
func BenchmarkParityBruteForceImproved(b *testing.B) { benchParityFn(b, ParityBruteForceImproved) }
func BenchmarkParity(b *testing.B)                   { benchParityFn(b, Parity) }
func BenchmarkParityBitsOnesCount(b *testing.B)      { benchParityFn(b, ParityBitsOnesCount) }

var parityTestData = []struct {
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
