package primitivetypes

import "math"

type Rect struct {
	x      float64
	y      float64
	width  float64
	height float64
}

// RectangleIntersection returns intersect rectangle for given two rectangles that has intersection,
// otherwise an empty intersection.
// The time complexity is O(1).
func RectangleIntersection(r1 Rect, r2 Rect) Rect {
	if !isIntersect(r1, r2) {
		return Rect{0, 0, -1, -1}
	}
	return Rect{
		math.Max(r1.x, r2.x), math.Max(r1.y, r2.y),
		math.Min(r1.x+r1.width, r2.x+r2.width) - math.Max(r1.x, r2.x),
		math.Min(r1.y+r1.height, r2.y+r2.height) - math.Max(r1.y, r2.y)}
}

func isIntersect(r1 Rect, r2 Rect) bool {
	return r1.x <= r2.x+r2.width && r1.x+r1.width >= r2.x && r1.y <= r2.y+r2.height && r1.y+r1.height >= r2.y
}
