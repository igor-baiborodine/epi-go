package primitivetypes

import (
	"math"
	"testing"
)

func TestComputeDivide(t *testing.T) {
	for _, test := range ComputeDivideTestData {
		if got := ComputeDivide(test.x, test.y); got != test.want {
			t.Errorf("ComputeDivide(%d, %d) = %d; want %d", test.x, test.y, got, test.want)
		}
	}
}

func BenchmarkComputeDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ComputeDivide(math.MaxUint64, math.MaxUint8)
	}
}

var ComputeDivideTestData = []struct {
	x    uint64
	y    uint64
	want uint64
}{
	{64, 1, 64},
	{64, 2, 32},
	{64, 3, 21},
	{64, 4, 16},
	{64, 5, 12},
	{65, 2, 32},
	{4, 2, 2},
	{9411, 4714, 1},
	{8186, 19, 430},
	{1313843, 515955, 2},
	{438, 1268, 0},
	{1441761, 7587904, 0},
	{28847325, 1260, 22894},
	{286, 8111194, 0},
	{161, 13, 12},
	{1379283057, 775, 1779720},
	{12, 647875067, 0},
	{15, 3000, 0},
	{1928, 23119, 0},
	{252764, 451235, 0},
}
