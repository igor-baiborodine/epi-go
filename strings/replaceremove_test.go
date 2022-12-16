package strings

import (
	"math"
	"reflect"
	"runtime"
	"testing"
)

type replaceRemoveFn func(n int, s []string) []string

func fnName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func testReplaceRemoveFn(t *testing.T, fn replaceRemoveFn) {
	for _, test := range replaceRemoveTest {
		if got := fn(test.n, test.in); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s(%v, %v)=%v; want %v", fnName(fn), test.n, test.in, got, test.want)
		}
	}
}

func TestReplaceRemove(t *testing.T) { testReplaceRemoveFn(t, ReplaceRemove) }

func benchmarkReplaceRemoveFn(b *testing.B, fn replaceRemoveFn) {
	b.StopTimer()
	n := math.MaxInt16
	s := make([]string, 2*n)

	for i := 0; i < n; i++ {
		if i < n/2 {
			s[i] = "a"
		} else {
			s[i] = "b"
		}
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		fn(n, s)
	}
}

func BenchmarkReplaceRemove(b *testing.B) { benchmarkReplaceRemoveFn(b, ReplaceRemove) }

var replaceRemoveTest = []struct {
	n    int
	in   []string
	want []string
}{
	{0, []string{}, []string{}},
	{1, []string{"a", "", "", ""}, []string{"d", "d"}},
	{2, []string{"a", "b", "", ""}, []string{"d", "d"}},
	{2, []string{"b", "a", "", ""}, []string{"d", "d"}},
	{3, []string{"a", "b", "a", ""}, []string{"d", "d", "d", "d"}},
	{3, []string{"a", "a", "b", ""}, []string{"d", "d", "d", "d"}},
	{3, []string{"b", "a", "b", ""}, []string{"d", "d"}},
	{3, []string{"b", "b", "a", ""}, []string{"d", "d"}},
	{3, []string{"a", "b", "c", ""}, []string{"d", "d", "c"}},
	{3, []string{"a", "b", "c", ""}, []string{"d", "d", "c"}},
	{6, []string{"b", "c", "b", "a", "a", "a", "", "", "", "", "", ""}, []string{"c", "d", "d", "d", "d", "d", "d"}},
	{6, []string{"b", "c", "c", "d", "c", "b", "", "", "", "", "", ""}, []string{"c", "c", "d", "c"}},
	{14, []string{"c", "d", "b", "a", "a", "d", "c", "c", "a", "c", "d", "d", "a", "d", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}, []string{"c", "d", "d", "d", "d", "d", "d", "c", "c", "d", "d", "c", "d", "d", "d", "d", "d"}},
	{27, []string{"a", "a", "d", "c", "a", "a", "d", "d", "b", "b", "d", "a", "b", "a", "c", "b", "b", "d", "c", "b", "b", "d", "b", "d", "c", "a", "a", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}, []string{"d", "d", "d", "d", "d", "c", "d", "d", "d", "d", "d", "d", "d", "d", "d", "d", "d", "c", "d", "c", "d", "d", "c", "d", "d", "d", "d"}},
}
