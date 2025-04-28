package strings

import (
	_ "math"
	"testing"
)

func TestIntToString(t *testing.T) {
	for _, test := range intStringConvTestData {
		if got := IntToString(test.n); got != test.s {
			t.Errorf("IntToStirng(%d) = %s; want %s", test.n, got, test.s)
		}
	}
}

func TestStringToInt(t *testing.T) {
	for _, test := range intStringConvTestData {
		if got := StringToInt(test.s); got != test.n {
			t.Errorf("StringToInt(\"%s\") = %d; want %d", test.s, got, test.n)
		}
	}
}
func BenchmarkIntToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntToString(9223372036854775807) // math.MaxInt64
	}
}

func BenchmarkStringToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringToInt("9223372036854775807")
	}
}

var intStringConvTestData = []struct {
	n int
	s string
}{
	{4, "4"},
	{4, "4"},
	{-4, "-4"},
	{4176473, "4176473"},
	{4176473, "4176473"},
	{-4176473, "-4176473"},
	{1435211240, "1435211240"},
	{1435211240, "1435211240"},
	{-1435211240, "-1435211240"},
	{7964, "7964"},
	{7964, "7964"},
	{-7964, "-7964"},
	{507038918, "507038918"},
	{507038918, "507038918"},
	{-507038918, "-507038918"},
	{5708618, "5708618"},
	{5708618, "5708618"},
	{-5708618, "-5708618"},
	{239135, "239135"},
	{239135, "239135"},
	{-239135, "-239135"},
}
