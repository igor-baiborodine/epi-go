package arrays

const (
	red = iota
	white
	blue
)

// DutchflagWithSubArrays reorders slice so that all elements
// less than s[i] appear first, followed by elements equal to s[i],
// followed by elements greater than s[i].
// The time complexity is O(n) and space complexity is O(n).
func DutchflagWithSubArrays(s []int, p int) []int {
	var r, w, b []int

	for _, v := range s {
		if v < s[p] {
			r = append(r, v)
		} else if v == s[p] {
			w = append(w, v)
		} else {
			b = append(b, v)
		}
	}
	s = append(r, w...)
	return append(s, b...)
}
