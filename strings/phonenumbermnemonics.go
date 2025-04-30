package strings

type Key int

var keyMappings = map[int][]rune{
	0: []rune("0"),
	1: []rune("1"),
	2: []rune("ABC"),
	3: []rune("DEF"),
	4: []rune("GHI"),
	5: []rune("JKL"),
	6: []rune("MNO"),
	7: []rune("PQRS"),
	8: []rune("TUV"),
	9: []rune("WXYZ"),
}

// ComputeMnemonics generates all possible letter combinations (mnemonics) for a given phone number
// where each digit maps to a set of letters as defined in keyMappings.
//
// Parameters:
// - n: A string representing the phone number, containing digits from 0-9
//
// Returns:
// - A slice of strings containing all possible letter combinations
//
// Time Complexity: O(4^n) where n is the length of the phone number, as each digit can map to up to 4 letters
// Space Complexity: O(n) for recursion stack and O(4^n) for storing all combinations
func ComputeMnemonics(n string) []string {
	var m []string
	pm := make([]rune, len(n))
	computePartialMnemonics(n, 0, pm, &m)
	return m
}

// computePartialMnemonics recursively generates all possible letter combinations for a phone number
// by processing one digit at a time and building partial combinations.
//
// Parameters:
// - n: The input phone number string
// - digitIdx: Current digit index being processed
// - partialMnemonics: Current partial combination being built
// - mnemonics: Pointer to slice storing all generated combinations
//
// Time Complexity: O(4^n) where n is the length of the phone number
// Space Complexity: O(n) for recursion stack
func computePartialMnemonics(n string, digitIdx int, partialMnemonics []rune, mnemonics *[]string) {
	if digitIdx == len(n) {
		// All digits are processed, so add partialMnemonic to mnemonics.
		*mnemonics = append(*mnemonics, string(partialMnemonics))
	} else {
		// Try all possible characters for this digitIdx.
		digit := int(n[digitIdx] - '0')
		letters := keyMappings[digit]
		for i := 0; i < len(letters); i++ {
			partialMnemonics[digitIdx] = letters[i]
			computePartialMnemonics(n, digitIdx+1, partialMnemonics, mnemonics)
		}
	}
}
