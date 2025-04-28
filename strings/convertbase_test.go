package strings

import (
	"testing"
)

func TestConvertBase(t *testing.T) {
	for _, test := range convertBaseTestData {
		if got := ConvertBase(test.s, test.fb, test.tb); got != test.want {
			t.Errorf("ConvertBase(%s, %d, %d) = %s; want %s", test.s, test.fb, test.tb, got, test.s)
		}
	}
}

var convertBaseTestData = []struct {
	s    string
	fb   int
	tb   int
	want string
}{
	{"615", 7, 13, "1A7"},
	{"7345", 10, 7, "30262"},
	{"-7345", 10, 7, "-30262"},
	{"42", 13, 5, "204"},
	{"-42", 13, 5, "-204"},
	{"1512", 13, 10, "3057"},
	{"-1512", 13, 10, "-3057"},
	{"4B5C", 14, 5, "410324"},
	{"-4B5C", 14, 5, "-410324"},
	{"210124203402", 6, 2, "101110111111111001100011011010"},
	{"-210124203402", 6, 2, "-101110111111111001100011011010"},
	{"158C3", 13, 13, "158C3"},
	{"-158C3", 13, 13, "-158C3"},
	{"143643223", 7, 10, "9523510"},
	{"-143643223", 7, 10, "-9523510"},
	{"A40", 16, 15, "B9E"},
	{"-A40", 16, 15, "-B9E"},
	{"2470755", 9, 8, "5104126"},
	{"-2470755", 9, 8, "-5104126"},
}
