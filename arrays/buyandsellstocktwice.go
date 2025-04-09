package arrays

import (
	"math"

	"github.com/shopspring/decimal"
)

// BuyAndSellStockTwice calculates the maximum profit that can be obtained by buying
// and selling a stock at most twice. The function uses a forward and backward pass
// to compute the cumulative maximum profit for every day in the array.
//
// Parameters:
//   - s: A slice of float64 values where each value represents the stock price on a given day.
//
// Returns:
//   - A float64 value representing the maximum profit possible with at most two transactions.
//
// Time Complexity: O(n), where n is the length of the n slice,
// due to single forward and backward traversals of the array.
// Space Complexity: O(n), for storing the intermediate profits in the firstBuySellProfits slice.
func BuyAndSellStockTwice(s []float64) float64 {
	var firstBuySellProfits []decimal.Decimal
	maxTotalProfit := decimal.NewFromFloat(0.0)
	minPriceSoFar := decimal.NewFromFloat(math.MaxFloat64)

	// Forward phase. For each day, we record maximum profit if we sell on that day
	for _, f := range s {
		price := decimal.NewFromFloat(f)
		minPriceSoFar = decimal.Min(minPriceSoFar, price)
		maxTotalProfit = decimal.Max(maxTotalProfit, price.Sub(minPriceSoFar))
		firstBuySellProfits = append(firstBuySellProfits, maxTotalProfit)
	}

	// Backward phase. For each day, find the maximum profit if we make the second buy on that day.
	maxPriceSoFar := decimal.NewFromFloat(0.0)
	for i := len(s) - 1; i > 0; i-- {
		price := decimal.NewFromFloat(s[i])
		maxPriceSoFar = decimal.Max(maxPriceSoFar, price)
		maxTotalProfit =
			decimal.Max(maxTotalProfit, maxPriceSoFar.Sub(price).Add(firstBuySellProfits[i-1]))
	}
	return maxTotalProfit.InexactFloat64()
}
