package arrays

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestComputeMaxProfit(t *testing.T) {
	for _, test := range ComputeMaxProfitTestData {
		if got := ComputeMaxProfit(test.s); got != test.want {
			t.Errorf("ComputeMaxProfit(%v) = %f; want %f", test.s, got, test.want)
		}
	}
}

func BenchmarkComputeMaxProfit(b *testing.B) {
	s := GenerateSliceFloat64(math.MaxInt16, 1, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ComputeMaxProfit(s)
	}
	b.ReportAllocs()
}

func GenerateSliceFloat64(size int, min, max float64) []float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := make([]float64, size)

	for i := 0; i < size; i++ {
		s[i] = min + r.Float64()*(max-min)
	}
	return s
}

var ComputeMaxProfitTestData = []struct {
	s    []float64
	want float64
}{
	{[]float64{}, 0},
	{[]float64{100}, 0},
	{[]float64{0.2, 0.2, 0.2}, 0.0},           // Buy at 0.2 and sell at 0.2
	{[]float64{0.3, 0.2, 0.3}, 0.1},           // Buy at 0.2 and sell at 0.3
	{[]float64{0.5, 0.3, 0.4, 0.1, 0.2}, 0.1}, // Buy at 0.3 and sell at 0.4
	{[]float64{0.7, 0.6, 0.5, 1.1, 0.5, 0.9, 0.4, 0.6, 1.1, 0.6, 0.9}, 0.7},                                              // Buy at 0.4 and sell at 1.1
	{[]float64{0.3, 0.3, 0.1}, 0.0},                                                                                      // Buy at 0.3 and sell at 0.3
	{[]float64{0.4, 0.5, 0.3, 0.1, 1.9, 0.3, 1.0, 0.7, 0.1, 1.1, 1.7, 1.5, 0.7, 1.2, 1.8, 0.4, 0.6, 0.8, 1.5, 1.4}, 1.8}, // Buy at 0.1 and sell at 1.9
	{[]float64{0.2, 2.9, 1.3, 3.3, 1.1, 1.0, 0.3, 2.4, 2.8, 0.6, 2.4, 3.1, 1.3, 2.7, 0.6, 1.2, 1.2, 1.6, 1.0, 1.5, 2.8, 2.5, 3.1, 1.9, 1.2, 1.2, 1.1, 3.3, 2.3, 0.8, 2.8, 2.2, 0.1, 1.0, 2.2}, 3.1}, // Buy at 0.2 and sell at 3.3
}
