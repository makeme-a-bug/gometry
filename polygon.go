package geometry

import (
	"fmt"

	"github.com/makeme-a-bug/gometry/utils"
)

// **** remaining function ****
// 1. create a constructor for the polygon(constructor should have to parameters (ordered_point []Points , holes []points-> can be none))
// 3. the validation funciton needs to be more concise

// polygon points should be ordered in a way that is a polygon the start cordinate and end corddinate should be the same
type Polygon struct {
	Points []Point
}

// contructor for the polygon checks validity and informs if the polygon is not valid
// returns the polygon,validity(bool)
func NewPolygon(points [][]float64) (Polygon, bool) {
	var polygon Polygon
	polygon.assignPoints(points)
	if polygon.Points[0] != polygon.Points[len(polygon.Points)-1] {
		polygon.Points = append(polygon.Points, polygon.Points[0])
	}
	return polygon, polygon.Valid()
}

// returns the area of the polygon
func (p Polygon) Area() float64 {
	return area_of_polygon(p.Coords())
}

// returns the type of polygon
func (p Polygon) GeoType() int {
	return 1
}

// returns the length
func (p Polygon) Length() float64 {
	length := float64(0)
	points := p.Points
	for i := 0; i < len(points)-1; i++ {
		length += Distance(points[i], points[i+1])
	}
	return length
}

// returns the center of the polygon
func (p Polygon) Center() []float64 {
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

// returns the coords in the form of a [][]float64
func (p Polygon) Coords() [][]float64 {
	var coords [][]float64
	for _, point := range p.Points {
		coords = append(coords, []float64{point.X, point.Y})
	}
	return coords
}

// returns the maxx , maxy , minx , miny in the form of a []float64
func (p Polygon) Bounds() []float64 {
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

// assigns new points to the polygon and removes the old one its only for internal usage e.g private
func (p *Polygon) assignPoints(points [][]float64) {
	var new_points []Point
	for _, p := range points {
		new_points = append(new_points, Point{X: p[0], Y: p[1]})
	}
	p.Points = new_points
}

// orders the points and assigns them to the polygon
func (p *Polygon) order_points() {
	temp := PolySort(*p)
	p.assignPoints(temp)
}

// checks the validity of the polygon
func (p Polygon) Valid() bool {
	// checks the validity of the polygon

	if p.Points[0] != p.Points[len(p.Points)-1] {
		fmt.Println("The first and last point should be the same")
		return false
	}

	// check if the number of points is greater than 3
	if len(p.Points) < 4 {
		return false
	}

	// // check if the points are in the form of a polygon
	// for i := 0; i < len(p.Points); i++ {
	// 	fmt.Println(p.Points[i].X, p.Points[(i+1)%len(p.Points)].X, p.Points[i].Y, p.Points[(i+1)%len(p.Points)].Y)
	// 	if p.Points[i].X == p.Points[(i+1)%len(p.Points)].X && p.Points[i].Y == p.Points[(i+1)%len(p.Points)].Y {
	// 		fmt.Println("The points are not in the form of a polygon")
	// 		return false
	// 	}
	// }

	// check if the edges intersect
	// fmt.Println(SweepLinesAlgorithm(p))
	if SweepLinesAlgorithm(p.Coords()) {
		return false
	}

	return true
}
