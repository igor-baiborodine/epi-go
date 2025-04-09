package arrays

import (
	"math"
	"testing"
)

func TestBuyAndSellStockTwice(t *testing.T) {
	for _, test := range BuyAndSellStockTwiceTestData {
		if got := BuyAndSellStockTwice(test.s); got != test.want {
			t.Errorf("BuyAndSellStockTwice(%v) = %f; want %f", test.s, got, test.want)
		}
	}
}

func BenchmarkBuyAndSellStockTwice(b *testing.B) {
	s := GenerateSliceFloat64(math.MaxInt16, 1, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BuyAndSellStockTwice(s)
	}
	b.ReportAllocs()
}

var BuyAndSellStockTwiceTestData = []struct {
	s    []float64
	want float64
}{
	{[]float64{0.3, 0.1, 0.4, 0.1}, 0.3},                                                                  // First transaction buy at 0.3 and sell first at 0.3, and second transaction buy at 0.1 and sell at 0.4
	{[]float64{0.2, 0.7, 0.9, 0.4, 0.1, 0.1, 0.5, 0.4, 0.9}, 1.5},                                         // First transaction buy at 0.2 and sell first at 0.9, and second transaction buy at 0.1 and sell at 0.9
	{[]float64{0.5, 0.6, 0.5, 1.1, 0.4, 0.2, 0.2, 0.2, 0.9, 0.6, 1.1}, 1.5},                               // First transaction buy at 0.5 and sell first at 1.1, and second transaction buy at 0.2 and sell at 1.1
	{[]float64{1.5, 0.6, 1.4, 0.1, 1.2, 1.1, 0.4, 0.9, 1.1, 1.7, 0.7, 0.4, 1.7, 1.3, 0.5, 0.7, 1.0}, 2.9}, // First transaction buy at 0.1 and sell first at 1.7, and second transaction buy at 0.4 and sell at 1.7
	{[]float64{2.9, 2.1, 1.9, 2.2, 1.5, 1.6, 1.2, 0.2, 2.3, 1.9, 1.5, 0.8, 0.9, 2.8, 2.7, 1.5, 1.1, 2.4, 0.1, 2.4, 2.2, 1.0, 2.1, 2.1, 1.0, 0.8, 2.2, 2.8, 0.9}, 5.3},                                                                                                                                                                                                          // First transaction buy at 0.2 and sell first at 2.8, and second transaction buy at 0.1 and sell at 2.8
	{[]float64{0.9, 2.1, 3.0, 1.2, 2.1, 1.8, 1.3, 1.5, 2.5, 1.9, 0.7, 2.7, 3.0, 2.7, 1.4, 2.5, 2.3, 1.1, 0.3, 2.0, 1.9, 3.0, 1.3, 2.4, 0.8, 1.9, 0.6, 1.6, 0.2, 2.9}, 5.4},                                                                                                                                                                                                     // First transaction buy at 0.3 and sell first at 3.0, and second transaction buy at 0.2 and sell at 2.9
	{[]float64{6.9, 4.8, 3.1, 1.6, 4.8, 0.6, 0.8, 1.3, 1.3, 0.9, 1.6, 3.2, 0.4, 2.7, 0.4, 5.1, 1.2, 2.1, 5.9, 4.3, 0.3, 4.7, 5.8, 1.8, 1.2, 1.7, 0.9, 4.1, 0.4, 3.8, 3.0, 0.5, 3.6, 3.1, 6.9, 4.4, 4.5, 2.7, 4.9, 2.0, 3.0, 5.3, 2.2, 4.9, 6.8, 6.0, 1.7, 4.5, 0.3, 6.0, 6.7, 3.4, 0.2, 1.7, 5.0, 1.8, 3.2, 3.6, 4.8, 0.2, 4.4, 2.6, 0.8, 2.6, 0.6, 5.9, 0.2, 0.9, 4.2}, 13.0}, // First transaction buy at 0.3 and sell first at 6.9, and second transaction buy at 0.3 and sell at 6.7
}
