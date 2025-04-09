package arrays

import (
	"math"

	"github.com/shopspring/decimal"
)

// ComputeMaxProfit calculates the maximum profit that can be made by buying and
// selling a stock once, given the list of stock prices s.
// The function iterates through the prices, maintaining the minimum price seen
// so far and the maximum profit possible. Decimal-based operations are used to
// ensure precision when performing arithmetic calculations.
//
// Parameters:
//   - s: A slice of float64 values where each value represents the stock price on a given day.
//
// Returns:
//   - A float64 value representing the maximum profit that can be achieved.
//
// Time Complexity: O(n), where n is the length of the slice s, as the function
// iterates through the slice once.
// Space Complexity: O(1), constant space is used as no additional data structures
// are employed.
func ComputeMaxProfit(s []float64) float64 {
	minPrice := decimal.NewFromFloat(math.MaxFloat64)
	maxProfit := decimal.NewFromFloat(0.0)

	for _, f := range s {
		price := decimal.NewFromFloat(f)
		maxProfit = decimal.Max(maxProfit, price.Sub(minPrice))
		minPrice = decimal.Min(minPrice, price)
	}
	return maxProfit.InexactFloat64()
}
