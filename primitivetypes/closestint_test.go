package primitivetypes

import (
	"math"
	"testing"
)

func TestClosestInt(t *testing.T) {
	for _, test := range closestIntTestData {
		if got, ok := ClosestInt(test.x); got != test.want || ok != test.ok {
			t.Errorf("ClosestInt(%d) = %d, %t; want %d, %t", test.x, got, ok, test.want, test.ok)
		}
	}
}

func BenchmarkClosestInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClosestInt(math.MaxUint64)
	}
}

var closestIntTestData = []struct {
	x    uint64
	want uint64
	ok   bool
}{
	{0, 0, false},                    // All 0
	{18446744073709551615, 0, false}, // All 1
	{48, 40, true},
	{14, 13, true},
	{39698800462691, 39698800462693, true},
	{626143, 626159, true},
	{130743209, 130743210, true},
	{47572837640579, 47572837640581, true},
	{477960486436286508, 477960486436286506, true},
	{681162947378, 681162947377, true},
	{165314, 165313, true},
	{1829739316, 1829739314, true},
	{108174474857473513, 108174474857473514, true},
	{7542726383620640, 7542726383620624, true},
	{12271042128255, 12271042128319, true},
	{581302105924, 581302105922, true},
	{15421717530450, 15421717530449, true},
	{6471, 6475, true},
	{15576770168, 15576770164, true},
	{1673, 1674, true},
	{483279878541583728, 483279878541583720, true},
}
