package arrays

import (
	"reflect"
	"testing"
)

type plusoneFn func(s []uint8) []uint8

func testPlusoneFn(t *testing.T, fn plusoneFn) {
	for _, test := range plusoneTestData {
		in := make([]uint8, len(test.in))
		copy(in, test.in)

		if got := fn(in); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s(%d) = %d; want %d", fnName(fn), test.in, got, test.want)
		}
	}
}

func TestPlusone(t *testing.T) { testPlusoneFn(t, PlusOne) }

func benchPlusoneFn(b *testing.B, fn plusoneFn) {
	for i := 0; i < b.N; i++ {
		fn([]uint8{1, 7, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9})
	}
}

func BenchmarkPlusOne(b *testing.B) { benchPlusoneFn(b, PlusOne) }

var plusoneTestData = []struct {
	in   []uint8
	want []uint8
}{
	{[]uint8{1, 2, 9}, []uint8{1, 3, 0}},
	{[]uint8{9, 9, 9, 9}, []uint8{1, 0, 0, 0, 0}},
	{[]uint8{9, 9}, []uint8{1, 0, 0}},
	{[]uint8{1}, []uint8{2}},
	{[]uint8{0}, []uint8{1}},
	{[]uint8{2, 2}, []uint8{2, 3}},
}
