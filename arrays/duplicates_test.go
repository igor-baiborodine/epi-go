package arrays

import (
	"math"
	"reflect"
	"testing"
)

type deleteduplicatesFn func(s []int) []int

func testDeleteduplicatesFn(t *testing.T, fn deleteduplicatesFn) {
	for _, test := range deleteDuplicatesTestData {
		if got := fn(test.in); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s(%v)=%v; want %v", fnName(fn), test.in, got, test.want)
		}
	}
}

func TestDeleteduplicatesSubarray(t *testing.T) { testDeleteduplicatesFn(t, DeleteDuplicatesSubarray) }
func TestDeleteDuplicatesShift(t *testing.T)    { testDeleteduplicatesFn(t, DeleteDuplicatesShift) }

func benchDeleteDuplicatesFn(b *testing.B, fn deleteduplicatesFn) {
	for i := 0; i < b.N; i++ {
		fn(make([]int, math.MaxInt16))
	}
}

func BenchmarkDeleteDuplicatesSubarray(b *testing.B) {
	benchDeleteDuplicatesFn(b, DeleteDuplicatesSubarray)
}
func BenchmarkDeleteDuplicatesShift(b *testing.B) { benchDeleteDuplicatesFn(b, DeleteDuplicatesShift) }

var deleteDuplicatesTestData = []struct {
	in   []int
	want []int
}{
	{[]int{}, []int{}},
	{[]int{-1}, []int{-1}},
	{[]int{0}, []int{0}},
	{[]int{1, 1, 1}, []int{1}},
	{[]int{-2, -1, 1}, []int{-2, -1, 1}},
	{[]int{1, 1, 2, 2, 3, 3, 4}, []int{1, 2, 3, 4}},
	{[]int{-2, -2, 0}, []int{-2, 0}},
	{[]int{-11, -10, -8, -7, -6, -6, -6, -2, -1, -1, 0, 0, 0, 0, 3, 4, 4, 7, 7, 8, 9, 11}, []int{-11, -10, -8, -7, -6, -2, -1, 0, 3, 4, 7, 8, 9, 11}},
	{[]int{-12, -12, -12, -10, -5, -5, -5, -3, -3, -2, -1, -1, 0, 1, 1, 4, 4, 6, 7, 7, 9, 10, 12, 12}, []int{-12, -10, -5, -3, -2, -1, 0, 1, 4, 6, 7, 9, 10, 12}},
	{[]int{-3, -3, -2, -2, -2, -1, 1, 4}, []int{-3, -2, -1, 1, 4}},
	{[]int{-8, -8, -7, -5, -5, -4, 1, 2, 2, 2, 2, 4, 5, 5, 7, 7}, []int{-8, -7, -5, -4, 1, 2, 4, 5, 7}},
}
