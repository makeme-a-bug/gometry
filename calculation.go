package geometry

import "math"

func Distance(g1, g2 GeoObject) float64 {
	coords1 := g1.Center()
	coords2 := g2.Center()
	return math.Sqrt(math.Pow(coords1[0]-coords2[0], 2) + math.Pow(coords1[1]-coords2[1], 2))
}

func Union(gs ...GeoObject) float64 {
	return 12
}

func Intersection(gs ...GeoObject) float64 {
	return 22
}

// Returns the squared distance and angle between two points.
func squaredPolar(point, centre []float64) (float64, float64) {
	dx := math.Atan2(point[1]-centre[1], point[0]-centre[0])
	dy := math.Pow(point[0]-centre[0], 2) + math.Pow(point[1]-centre[1], 2)
	return dx, dy
}

// Returns the ordered points in the anti-clockwise direction of the polygon in the form of a [][]float64
func PolySort(polygon Polygon) [][]float64 {

	centre := polygon.Center()
	var points [][]float64
	for _, x := range polygon.Coords() {
		dx, dy := squaredPolar(x, centre)
		temp := []float64{x[0], x[1], dx, dy}
		points = append(points, temp)
	}

	// Sort by polar angle and distance, centered at this centre of mass.

	points = sortPoints(points)

	// Throw away the temporary polar coordinates
	for i := 0; i < len(points); i++ {
		points[i] = points[i][:2]
	}

	return points
}

// sorts the points according to the polar angle and distance
func sortPoints(points [][]float64) [][]float64 {
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			if points[i][2]-points[j][2] != 0 {
				if points[i][2] > points[j][2] {
					points[i], points[j] = points[j], points[i]
				}
			} else {
				if points[i][3] > points[j][3] {
					points[i], points[j] = points[j], points[i]
				}
			}
		}
	}

	return points
}

// Returns true if the geoObject intersects with other geoObject
func Intersects(g1, g2 GeoObject) bool {
	points := g1.Coords()
	poly := g2.Coords()

	intercets := false
	for _, point := range points {
		j := len(poly) - 1
		for i := 0; i < len(poly); i++ {
			if (point[0] == poly[i][0]) && (point[1] == poly[i][1]) {
				// point is a corner
				return true
			}

			if (poly[i][1] > point[1]) != (poly[j][1] > point[1]) {
				slope := (point[0]-poly[i][0])*(poly[j][1]-poly[i][1]) - (poly[j][0]-poly[i][0])*(point[1]-poly[i][1])
				if slope == 0 {
					// point is on boundary
					return true
				}
				if (slope < 0) != (poly[j][1] < poly[i][1]) {
					intercets = !intercets
				}
			}

			j = i
		}
	}

	return intercets
}
