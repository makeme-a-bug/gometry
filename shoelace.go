package geometry

import "math"

func area_of_polygon(points [][]float64) float64 {
	var area float64
	segments := segment(points)
	sum := float64(0)
	for i := 0; i < len(segments); i++ {
		sum += segments[i][0][0]*segments[i][1][1] - segments[i][1][0]*segments[i][0][1]
	}
	area = sum / 2

	return math.Abs(area)
}

func segment(points [][]float64) [][][]float64 {
	var segments [][][]float64
	for i := 0; i < len(points)-1; i++ {
		edge := [][]float64{{points[i][0], points[i][1]}, {points[i+1][0], points[i+1][1]}}
		segments = append(segments, edge)
	}

	return segments
}
