package strings

// DecodeColId converts a string representing a column ID (e.g., "A", "Z", "AA")
// into its corresponding integer representation, where "A" maps to 1, "B" to 2,
// and so on, following a base-26 number system.
//
// Parameters:
//   - id: A string representing the column ID.
//
// Returns:
//   - An integer corresponding to the decoded column ID.
//
// Time Complexity: O(n), where n is the length of the input string id.
// Space Complexity: O(1), as no additional space is used apart from variables.
func DecodeColId(id string) int {
	var result int
	for _, c := range id {
		result = result*26 + int(c-'A'+1)
	}
	return result
}
