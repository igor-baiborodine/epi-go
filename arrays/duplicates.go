package arrays

// DeleteDuplicatesSubarray removes repeated elements from a sorted slice.
// The time complexity is O(n) and space complexity is O(n).
func DeleteDuplicatesSubarray(s []int) []int {
	if len(s) == 0 {
		return s
	}
	set := append([]int{}, s[0])

	for i := 1; i < len(s); i++ {
		if s[i] != set[len(set)-1] {
			set = append(set, s[i])
		}
	}
	return set
}

func DeleteDuplicatesShift(s []int) []int {
	if len(s) == 0 {
		return s
	}
	var w int
	//fmt.Println("[1] w, s:", w, s)

	for i := 1; i < len(s); i++ {
		//fmt.Println("[2] w, i, s:", w, i, s)
		if s[w] != s[i] {
			if w != i {
				s = append(s[:w+1], s[i:]...)
				w++
				i = w
				//fmt.Println("[3] w, i, s:", w, i, s)
			} else {
				w++
				//fmt.Println("[4] w, i, s:", w, i, s)
			}
		}
	}
	if s[w] == s[len(s)-1] {
		//fmt.Println("[5] w, len(s)-1:", w, len(s)-1)
		s = s[:w+1]
	}
	//fmt.Println("[6] s:", s)
	return s
}
