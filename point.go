package geometry

type Point struct {
	X float64
	Y float64
}

func NewPoint(points [2]float64) Point {
	return Point{X: points[0], Y: points[1]}
}

func (p Point) Coords() [][]float64 {
	return [][]float64{{p.X, p.Y}}
}

func (p Point) Bounds() []float64 {
	return []float64{p.X, p.Y}
}

func (p Point) Area() float64 {
	return 0
}

// returns the length
func (p Point) Length() float64 {
	return 0
}

func (p Point) GeoType() int {
	return 0
}

func (p Point) Center() []float64 {
	return p.Coords()[0]
}

func (p Point) Valid() bool {
	return true
}
