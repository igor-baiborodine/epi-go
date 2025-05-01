package strings

import (
	"testing"
)

func TestRabinKarp(t *testing.T) {
	for _, test := range rabinKarpTestData {
		if got := RabinKarp(test.s, test.sub); got != test.want {
			t.Errorf("RabinKarp(%s, %s) = %d; want %d", test.s, test.sub, got, test.want)
		}
	}
}

var rabinKarpTestData = []struct {
	s    string
	sub  string
	want int
}{
	{"A", "", 0},
	{"A", "A", 0},
	{"A", "B", -1},
	{"", "A", -1},
	{"", "AAA", -1},
	{"A", "AAA", -1},
	{"AA", "AAA", -1},
	{"AAA", "AAA", 0},
	{"BAAAA", "AAA", 1},
	{"GACGCCA", "CGC", 2},
	{"GATACCCATCGAGTCGGATCGAGT", "GAG", 10},
	{"FOOBARWIDGET", "WIDGETS", -1},
	{"ADSADA", "", 0},
	{"BAAABAAAA", "AAA", 1},
	{"BAABBAABAAABS", "AAA", 8},
	{"BAABBAABAAABS", "AAAA", -1},
	{"FOOBAR", "BAR", 3},
	{"babababbaabaabbbbb", "babba", 4},
	{"babababbaabaabbbbb", "abbbbbbbab", -1},
	{"babababbaabaabbbbb", "abaabbba", -1},
	{"babababbaabaabbbbb", "aba", 1},
	{"babababbaabaabbbbb", "aabbaaabb", -1},
	{"babababbaabaabbbbb", "bbaaaaaba", -1},
	{"babababbaabaabbbbb", "bba", 6},
	{"babababbaabaabbbbb", "abaaabaaab", -1},
	{"babababbaabaabbbbb", "ababa", 1},
	{"babababbaabaabbbbb", "babb", 4},
	{"babababbaabaabbbbb", "aaabaaaab", -1},
	{"babababbaabaabbbbb", "bb", 6},
}
