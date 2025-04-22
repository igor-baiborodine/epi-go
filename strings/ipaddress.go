package strings

import (
	"math"
	"strconv"
	"strings"
)

// ValidIpAddress takes a string s representing a numeric string
// and generates all possible valid IPv4 addresses by dividing the string
// into four segments. An IPv4 address is valid if each segment is
// an integer between 0 and 255 (inclusive) and does not have leading zeros.
//
// The function returns a slice of strings containing the valid IPv4 addresses
// formed from the input string.
//
// Time Complexity: O(n^3)
//   - There are three nested loops iterating through the string,
//     where n is the length of the string s.
//   - In the innermost loop, string slicing, leading zero checks,
//     and conversion to integers have constant time complexity.
//
// Space Complexity: O(m * n)
// - m is the number of valid IPv4 addresses generated.
// - Each IPv4 address can consume O(n) space in the worst case.
func ValidIpAddress(s string) []string {
	var r, p = []string{}, make([]string, 4)

first:
	for i := 1; i < minInt(4, len(s)); i++ {
		p[0] = s[:i]
		if !validPart(p[0]) {
			continue first
		}
	second:
		for j := 1; j < minInt(4, len(s)-i); j++ {
			p[1] = s[i : i+j]
			if !validPart(p[1]) {
				continue second
			}
			for k := 1; k < minInt(4, len(s)-i-j); k++ {
				p[2], p[3] = s[i+j:i+j+k], s[i+j+k:]
				if validPart(p[2]) && validPart(p[3]) {
					r = append(r, strings.Join(p, "."))
				}
			}
		}
	}
	return r
}

func minInt(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func validPart(s string) bool {
	// "00", "000'", "01", etc. are not valid, but '0' is valid
	if len(s) == 1 {
		return true
	}
	if s[0] == '0' {
		return false
	}
	i, err := strconv.Atoi(s)
	return err == nil && i <= 255
}
