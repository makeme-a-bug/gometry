package geometry

import (
	"math"
)

func Line_intersection(p1, q1, p2, q2 []float64) (float64, float64) {
	xdiff := []float64{p1[0] - q1[0], p2[0] - q2[0]}
	ydiff := []float64{p1[1] - q1[1], p2[1] - q2[1]}

	div := det(xdiff, ydiff)
	if div == 0 {
		return math.Inf(1), math.Inf(1)
	}
	d := []float64{det(p1, q1), det(p2, q2)}

	x := det(d, xdiff) / div
	if det(d, xdiff) == 0 {
		x = 0
	}
	y := det(d, ydiff) / div
	if det(d, ydiff) == 0 {
		y = 0
	}
	return x, y
}

func onSegment(p, q, r []float64) bool {
	if q[0] <= math.Max(p[0], r[0]) && q[0] >= math.Min(p[0], r[0]) && q[1] <= math.Max(p[1], r[1]) && q[1] >= math.Min(p[1], r[1]) {
		return true
	}
	return false
}

// To find orientation of ordered triplet (p, q, r).
// The function returns following values
// 0 --> p, q and r are collinear
// 1 --> Clockwise
// 2 --> Counterclockwise
func orientation(p, q, r []float64) int {
	val := int((q[1]-p[1])*(r[0]-q[0]) -
		(q[0]-p[0])*(r[1]-q[1]))

	if val == 0 {
		return 0 // collinear
	}
	if val > 0 {
		return 1 // clockwise
	}
	return 2 // counterclockwise
}

// The main function that returns true if line segment 'p1q1'
// and 'p2q2' intersect.
func DoIntersect(p1, q1, p2, q2 []float64) bool {
	// Find the four orientations needed for general and
	// special cases
	inter1, inter2 := Line_intersection(p1, q1, p2, q2)
	if inter1 != math.Inf(1) || inter2 != math.Inf(1) {
		for _, v := range [][]float64{p1, q1, p2, q2} {
			if v[0] == inter1 && v[1] == inter2 {
				return false
			}
		}

	}

	o1 := orientation(p1, q1, p2)
	o3 := orientation(p2, q2, p1)
	o2 := orientation(p1, q1, q2)
	o4 := orientation(p2, q2, q1)

	// General case
	if o1 != o2 && o3 != o4 {
		return true
	}

	cases_validated := 0
	// Special Cases
	// p1, q1 and p2 are collinear and p2 lies on segment p1q1
	if o1 == 0 && onSegment(p1, p2, q1) {
		cases_validated += 1
	}

	// p1, q1 and q2 are collinear and q2 lies on segment p1q1
	if o2 == 0 && onSegment(p1, q2, q1) {
		cases_validated += 1
	}

	// p2, q2 and p1 are collinear and p1 lies on segment p2q2
	if o3 == 0 && onSegment(p2, p1, q2) {
		cases_validated += 1
	}

	// p2, q2 and q1 are collinear and q1 lies on segment p2q2
	if o4 == 0 && onSegment(p2, q1, q2) {
		cases_validated += 1
	}

	if cases_validated > 1 {
		return true
	}

	return false // Doesn't fall in any of the above cases
}
