package strings

var romanToInt = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// RomanToInteger converts a Roman number string to an integer value.
// Parameters:
//   - s: A string representing a valid complex Roman number, e.g., 'LXIII'
//
// Returns:
//   - An integer representing the decimal value of the Roman number
//
// Time Complexity: O(n) where n is the length of the input string
// Space Complexity: O(1) as we only use a fixed-size map and constant extra space
func RomanToInteger(s string) int {
	r := []rune(s)
	sum := romanToInt[r[len(r)-1]]
	for i := len(r) - 2; i >= 0; i-- {
		intVal := romanToInt[r[i]]
		if intVal < romanToInt[r[i+1]] {
			sum -= intVal
		} else {
			sum += intVal
		}

	}
	return sum
}
