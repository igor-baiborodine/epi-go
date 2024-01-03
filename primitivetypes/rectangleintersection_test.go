package primitivetypes

import (
	"testing"
)

func TestRectangleIntersection(t *testing.T) {
	for _, test := range RectangleIntersectionTestData {
		if got := RectangleIntersection(test.r1, test.r2); got != test.want {
			t.Errorf("RectangleIntersection(%+v, %+v) = %+v; want %+v", test.r1, test.r2, got, test.want)
		}
	}
}

func BenchmarkRectangleIntersection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RectangleIntersection(Rect{76, 9, 12, 14}, Rect{20, 1, 62, 60})
	}
}

var RectangleIntersectionTestData = []struct {
	r1   Rect
	r2   Rect
	want Rect
}{
	{Rect{76, 9, 12, 14}, Rect{20, 1, 62, 60}, Rect{76, 9, 6, 14}},
	{Rect{54, 66, 66, 24}, Rect{27, 97, 68, 95}, Rect{0, 0, -1, -1}},
	{Rect{82, 9, 20, 84}, Rect{84, 87, 14, 82}, Rect{84, 87, 14, 6}},
	{Rect{1, 11, 38, 93}, Rect{78, 6, 33, 31}, Rect{0, 0, -1, -1}},
	{Rect{68, 9, 43, 28}, Rect{37, 47, 61, 60}, Rect{0, 0, -1, -1}},
	{Rect{43, 6, 89, 6}, Rect{66, 48, 65, 13}, Rect{0, 0, -1, -1}},
	{Rect{90, 32, 17, 75}, Rect{43, 95, 35, 95}, Rect{0, 0, -1, -1}},
	{Rect{39, 52, 74, 66}, Rect{76, 53, 24, 11}, Rect{76, 53, 24, 11}},
	{Rect{83, 64, 54, 42}, Rect{72, 93, 32, 1}, Rect{83, 93, 21, 1}},
	{Rect{4, 70, 53, 89}, Rect{57, 14, 70, 34}, Rect{0, 0, -1, -1}},
	{Rect{84, 72, 46, 37}, Rect{99, 91, 36, 29}, Rect{99, 91, 31, 18}},
	{Rect{36, 65, 69, 5}, Rect{99, 83, 99, 61}, Rect{0, 0, -1, -1}},
	{Rect{31, 2, 87, 62}, Rect{68, 78, 75, 46}, Rect{0, 0, -1, -1}},
	{Rect{18, 4, 79, 63}, Rect{37, 79, 86, 24}, Rect{0, 0, -1, -1}},
	{Rect{17, 9, 13, 10}, Rect{37, 64, 62, 12}, Rect{0, 0, -1, -1}},
	{Rect{82, 44, 65, 47}, Rect{83, 91, 61, 40}, Rect{83, 91, 61, 0}},
	{Rect{28, 95, 77, 30}, Rect{29, 47, 62, 75}, Rect{29, 95, 62, 27}},
	{Rect{91, 79, 57, 31}, Rect{63, 63, 96, 64}, Rect{91, 79, 57, 31}},
	{Rect{92, 83, 44, 91}, Rect{29, 8, 25, 72}, Rect{0, 0, -1, -1}},
	{Rect{26, 12, 76, 62}, Rect{68, 59, 77, 28}, Rect{68, 59, 34, 15}},
}
