package arrays

func AdvanceByOffsets(s []int) bool {
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
