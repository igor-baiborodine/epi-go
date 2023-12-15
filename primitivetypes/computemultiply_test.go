package primitivetypes

import (
	"math"
	"testing"
)

func TestComputeMultiply(t *testing.T) {
	for _, test := range ComputeMultiplyTestData {
		if got := ComputeMultiply(test.x, test.y); got != test.want {
			t.Errorf("ComputeMultiply(%d, %d) = %d; want %d", test.x, test.y, got, test.want)
		}
	}
}

func BenchmarkComputeMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ComputeMultiply(math.MaxUint8, math.MaxUint64)
	}
}

var ComputeMultiplyTestData = []struct {
	x    uint64
	y    uint64
	want uint64
}{
	{0, 0, 0},
	{0, 1, 0},
	{0, 65533, 0},
	{1, 65533, 65533},
	{345, 1, 345},
	{345, 0, 0},
	{57536, 2187, 125831232},
	{4639, 45265, 209984335},
	{54059, 27160, 1468242440},
	{27192, 45176, 1228425792},
	{52577, 27367, 1438874759},
	{49906, 12644, 631011464},
	{61443, 16564, 1017741852},
	{48076, 38151, 1834147476},
	{8155, 46880, 382306400},
	{39185, 13717, 537500645},
	{61599, 25652, 1580137548},
	{26336, 10917, 287510112},
	{50781, 18221, 925280601},
	{11092, 54111, 600199212},
}
