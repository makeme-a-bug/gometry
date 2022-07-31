package geometry

import "github.com/makeme-a-bug/gometry/utils"

type LinearRing struct {
	Points []Point
}

// contructor for the polygon checks validity and informs if the polygon is not valid
// returns the polygon,validity(bool)
func NewLinearRing(points [][]float64) (LinearRing, bool) {
	var linearRing LinearRing
	linearRing.assignPoints(points)
	if linearRing.Points[0] != linearRing.Points[len(linearRing.Points)-1] {
		linearRing.Points = append(linearRing.Points, linearRing.Points[0])
	}

	return linearRing, linearRing.Valid()
}

// returns the area
func (p LinearRing) Area() float64 {
	return 0
}

// returns the type
func (p LinearRing) GeoType() int {
	return 2
}

func (p LinearRing) Center() []float64 {
	var x_list []float64
	var y_list []float64
	for _, point := range p.Points {
		x_list = append(x_list, point.X)
		y_list = append(y_list, point.Y)
	}
	totalVertices := float64(len(p.Points))
	x := utils.Sum(x_list...) / totalVertices
	y := utils.Sum(y_list...) / totalVertices
	return []float64{x, y}
}

func (p *LinearRing) assignPoints(points [][]float64) {
	var new_points []Point
	for _, p := range points {
		new_points = append(new_points, Point{X: p[0], Y: p[1]})
	}
	p.Points = new_points
}

// returns the length
func (p LinearRing) Length() float64 {
	length := float64(0)
	points := p.Points
	for i := 0; i < len(points)-1; i++ {
		length += Distance(points[i], points[i+1])
	}
	return length
}

// returns the coords in the form of a [][]float64
func (p LinearRing) Coords() [][]float64 {
	var coords [][]float64
	for _, point := range p.Points {
		coords = append(coords, []float64{point.X, point.Y})
	}
	return coords
}

// returns the maxx , maxy , minx , miny in the form of a []float64
func (p LinearRing) Bounds() []float64 {
	var x_list []float64
	var y_list []float64
	for _, point := range p.Points {
		x_list = append(x_list, point.X)
		y_list = append(y_list, point.Y)
	}
	minx := utils.Min(x_list...)
	miny := utils.Min(y_list...)
	maxx := utils.Max(x_list...)
	maxy := utils.Max(y_list...)
	return []float64{minx, miny, maxx, maxy}
}

// checks the validity
func (p LinearRing) Valid() bool {
	// check if the number of points is greater than 4
	if len(p.Points) < 4 {
		return false
	}

	// check if the edges intersect
	if SweepLinesAlgorithm(p.Coords()) {
		return false
	}

	return true
}
