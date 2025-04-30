package strings

func DecodeColId(id string) int {
	var result int
	for _, c := range id {
		result = result*26 + int(c-'A'+1)
	}
	return result
}
