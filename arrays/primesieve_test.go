package arrays

import (
	"reflect"
	"testing"
)

func TestGeneratePrimes(t *testing.T) {
	for _, test := range generatePrimesTestData {
		if got := GeneratePrimes(test.n); !reflect.DeepEqual(got, test.want) {
			t.Errorf("GeneratePrimes(%v) = %v, want %v", test.n, got, test.want)
		}
	}
}

func BenchmarkGeneratePrimes(b *testing.B) {
	GeneratePrimes(10000000)
}

var generatePrimesTestData = []struct {
	n    int
	want []int
}{
	{1, []int{}},
	{2, []int{2}},
	{3, []int{2, 3}},
	{4, []int{2, 3}},
	{5, []int{2, 3, 5}},
	{6, []int{2, 3, 5}},
	{7, []int{2, 3, 5, 7}},
	{8, []int{2, 3, 5, 7}},
	{9, []int{2, 3, 5, 7}},
	{10, []int{2, 3, 5, 7}},
	{11, []int{2, 3, 5, 7, 11}},
	{12, []int{2, 3, 5, 7, 11}},
	{13, []int{2, 3, 5, 7, 11, 13}},
	{14, []int{2, 3, 5, 7, 11, 13}},
	{15, []int{2, 3, 5, 7, 11, 13}},
	{16, []int{2, 3, 5, 7, 11, 13}},
	{17, []int{2, 3, 5, 7, 11, 13, 17}},
	{18, []int{2, 3, 5, 7, 11, 13, 17}},
	{19, []int{2, 3, 5, 7, 11, 13, 17, 19}},
	{100, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},
}
