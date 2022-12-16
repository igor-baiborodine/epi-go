package strings

import (
	"math"
	"strconv"
	"strings"
)

func ValidIpAddress(s string) []string {
	var r, p = []string{}, make([]string, 4)

first:
	for i := 1; i < min(4, len(s)); i++ {
		p[0] = s[:i]
		if !validPart(p[0]) {
			continue first
		}
	second:
		for j := 1; j < min(4, len(s)-i); j++ {
			p[1] = s[i : i+j]
			if !validPart(p[1]) {
				continue second
			}
			for k := 1; k < min(4, len(s)-i-j); k++ {
				p[2], p[3] = s[i+j:i+j+k], s[i+j+k:]
				if validPart(p[2]) && validPart(p[3]) {
					r = append(r, strings.Join(p, "."))
				}
			}
		}
	}
	return r
}

func min(x, y int) int {
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
