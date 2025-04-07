package arrays

// CanReachEnd determines if it is possible to reach the last position of a slice
// by starting at the first position and advancing based on the values in the slice.
//
// The method uses a greedy approach to calculate the furthest index reachable at each step,
// iterating over the slice as long as the current position is within the reachable range.
//
// Parameters:
//   - s: A slice of integers where each element represents the maximum number of indices
//     that can be advanced from that position.
//
// Returns:
// - A boolean: true if the last index can be reached, false otherwise.
//
// Time Complexity: O(n), where n is the length of the slice s, since the method iterates
// through the slice at most once.
// Space Complexity: O(1), as the method uses only constant extra space.
func CanReachEnd(s []int) bool {
	if len(s) == 0 {
		return false
	}
	if len(s) == 1 {
		return true
	}
	curFurthestIdx := 0
	lastIdx := len(s) - 1

	for i := 0; i <= curFurthestIdx && curFurthestIdx < lastIdx; i++ {
		if curFurthestIdx < s[i]+i {
			curFurthestIdx = s[i] + i
		}
	}
	return curFurthestIdx >= lastIdx
}
