package strings

import (
	"testing"
)

func TestDecodeColId(t *testing.T) {
	for _, test := range decodeColIdTestData {
		if got := DecodeColId(test.s); got != test.want {
			t.Errorf("DecodeColId(%s) = %d; want %d", test.s, got, test.want)
		}
	}
}

var decodeColIdTestData = []struct {
	s    string
	want int
}{
	{"A", 1},
	{"D", 4},
	{"AA", 27},
	{"WQPMD", 10820398},
	{"YW", 673},
	{"WZWV", 422444},
	{"PPPO", 292463},
	{"ESJRGPI", 1775216309},
	{"CWRVJC", 46486079},
	{"OPEDF", 7139346},
	{"BDQZ", 38324},
	{"HN", 222},
	{"UIV", 14452},
	{"NEC", 9597},
	{"MYZC", 246067},
	{"BQZJTQQ", 831886059},
	{"JCGFNY", 120312165},
	{"GQQFGJ", 91241264},
	{"VX", 596},
	{"ULKQQ", 9815303},
	{"LQNESE", 150595047},
	{"GUSMJ", 3581120},
	{"QWIEVC", 212655979},
	{"BDLJ", 38178},
	{"XOMPW", 11240291},
	{"ULFFZ", 9811646},
	{"DJFU", 77241},
}
