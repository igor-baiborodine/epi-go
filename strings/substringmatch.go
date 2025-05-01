package strings

// RabinKarp implements the Rabin-Karp string matching algorithm.
// Parameters:
//   - s: source string to search in
//   - t: target string to search for
//
// Returns:
//   - index of the first occurrence of target string in source string
//   - -1 if target string is not found
//   - 0 if target string is empty
//
// Time Complexity: O(n+m) on average, O(nm) worst case,
//
//	where n is length of source string and m is length of target string
//
// Space Complexity: O(1) as it uses only a constant amount of extra space
func RabinKarp(s, t string) int {
	if len(t) == 0 {
		return 0
	}
	if len(s) < len(t) {
		return -1
	}
	const base = 256
	const prime = 101
	// Calculate hash value for pattern t
	patternHash := 0
	for i := 0; i < len(t); i++ {
		patternHash = (patternHash*base + int(t[i])) % prime
	}
	// Calculate initial hash value for first window of text s
	windowHash := 0
	for i := 0; i < len(t); i++ {
		windowHash = (windowHash*base + int(s[i])) % prime
	}
	// Calculate value of base^(m-1) where m is pattern length
	h := 1
	for i := 0; i < len(t)-1; i++ {
		h = (h * base) % prime
	}
	// Slide pattern over text one by one
	for i := 0; i <= len(s)-len(t); i++ {
		if patternHash == windowHash {
			// Check character by character
			match := true
			for j := 0; j < len(t); j++ {
				if s[i+j] != t[j] {
					match = false
					break
				}
			}
			if match {
				return i
			}
		}
		// Calculate hash value for next window
		if i < len(s)-len(t) {
			windowHash = (base*(windowHash-int(s[i])*h) + int(s[i+len(t)])) % prime
			if windowHash < 0 {
				windowHash += prime
			}
		}
	}
	return -1
}
