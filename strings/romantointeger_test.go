package strings

import (
	"testing"
)

func TestRomanToInteger(t *testing.T) {
	for _, test := range romanToIntegerTestData {
		if got := RomanToInteger(test.r); got != test.want {
			t.Errorf("RomanToInteger(%s) = %d; want %d", test.r, got, test.want)
		}
	}
}

var romanToIntegerTestData = []struct {
	r    string
	want int
}{
	{"I", 1},
	{"II", 2},
	{"III", 3},
	{"IV", 4},
	{"V", 5},
	{"VI", 6},
	{"VII", 7},
	{"VIII", 8},
	{"IX", 9},
	{"X", 10},
	{"XI", 11},
	{"XII", 12},
	{"XIII", 13},
	{"XIV", 14},
	{"XV", 15},
	{"XVI", 16},
	{"XVII", 17},
	{"XVIII", 18},
	{"XIX", 19},
	{"XX", 20},
	{"L", 50},
	{"LI", 51},
	{"LII", 52},
	{"LIII", 53},
	{"LIV", 54},
	{"LV", 55},
	{"LVI", 56},
	{"LVII", 57},
	{"LVIII", 58},
	{"LIX", 59},
	{"LX", 60},
	{"LXI", 61},
	{"LXII", 62},
	{"LXIII", 63},
	{"LXIV", 64},
	{"LXV", 65},
	{"LXVI", 66},
	{"LXVII", 67},
	{"LXVIII", 68},
}
