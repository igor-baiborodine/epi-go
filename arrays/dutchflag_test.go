package arrays

import (
	"math"
	"math/rand"
	"reflect"
	"runtime"
	"testing"
)

type dutchflagFn func(s []int, p int) []int

func testDutchflagFn(t *testing.T, fn dutchflagFn) {
	for _, test := range dutchflagTestData {
		if got := fn(test.in, test.p); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s(%d, %d) = %d; want %d", fnName(fn), test.in, test.p, got, test.want)
		}
	}
}

func fnName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func TestDutchflagWithSubArrays(t *testing.T) { testDutchflagFn(t, DutchflagWithSubArrays) }

func benchDutchflagFn(b *testing.B, fn dutchflagFn) {
	b.StopTimer()
	s := make([]int, math.MaxInt16)

	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(3)
	}

	for i := 0; i < b.N; i++ {
		p := rand.Intn(len(s))

		b.StartTimer()
		fn(s, p)
		b.StopTimer()
	}
}

func BenchmarkDutchflagWithSubArrays(b *testing.B) { benchDutchflagFn(b, DutchflagWithSubArrays) }

var dutchflagTestData = []struct {
	in   []int
	p    int
	want []int
}{
	{[]int{1}, 0, []int{1}},
	{[]int{0, 1, 1, 2}, 2, []int{0, 1, 1, 2}},
	{[]int{1, 1, 0, 2}, 1, []int{0, 1, 1, 2}},
	{[]int{1, 1, 0, 2, 0}, 0, []int{0, 0, 1, 1, 2}},
	{[]int{1, 1, 0, 2, 0}, 1, []int{0, 0, 1, 1, 2}},
	{[]int{1, 1, 0, 2, 0}, 2, []int{0, 0, 1, 1, 2}},
	{[]int{1, 1, 0, 2, 0}, 3, []int{1, 1, 0, 0, 2}},
	{[]int{1, 1, 0, 2, 0}, 4, []int{0, 0, 1, 1, 2}},
}
