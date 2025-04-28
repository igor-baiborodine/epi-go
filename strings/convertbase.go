package strings

import (
	"fmt"
	"strings"
	"unicode"
)

// ConvertBase converts a number represented as a string s in a specified fromBase
// to its equivalent representation in a different toBase. Negative numbers
// are handled by checking for the '-' prefix. It supports bases from 2 to 36 and
// throws an error for invalid characters.
//
// Steps:
// 1. Parse the input string s to its decimal (base 10) integer value using fromBase.
// 2. Convert the decimal value to the desired toBase.
//
// Parameters:
// - s: The number as a string in the fromBase.
// - fromBase: The base of the input number s (e.g., 2 for binary, 10 for decimal, etc.).
// - toBase: The desired base for the output (e.g., 16 for hexadecimal, 8 for octal, etc.).
//
// Returns:
// - A string representing the number converted to the toBase.
//
// Panics:
// - If the input contains invalid characters for the specified fromBase.
//
// Time Complexity:
//   - O(n + m), where n is the length of the input string s (for conversion
//     from fromBase to a decimal number) and m is the number of digits in the
//     output base representation (for conversion from decimal to toBase).
//
// Space Complexity:
//   - O(m), where m is the number of digits in the output base representation (due to
//     the use of the `strings.Builder` for constructing the result).
func ConvertBase(s string, fromBase, toBase int) string {
	var result int
	var neg bool
	i := 0

	if strings.HasPrefix(s, "-") {
		neg = true
		i = 1
	}
	for ; i < len(s); i++ {
		result *= fromBase

		if unicode.IsDigit(rune(s[i])) {
			result += int(s[i] - '0')
		} else if unicode.IsLetter(rune(s[i])) {
			result += int(unicode.ToUpper(rune(s[i]))-'A') + 10
		} else {
			panic(fmt.Sprintf("Invalid character '%c' in input", s[i]))
		}
	}
	if result == 0 {
		return "0"
	}

	var sb strings.Builder
	for result > 0 {
		remainder := result % toBase
		result /= toBase

		if remainder < 10 {
			sb.WriteByte(byte('0' + remainder))
		} else {
			sb.WriteByte(byte('A' + remainder - 10))
		}
	}
	converted := ReverseString(sb.String())

	if neg {
		return "-" + converted
	}
	return converted
}
