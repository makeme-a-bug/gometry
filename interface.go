package geometry

type GeoObject interface {
	Coords() [][]float64
	Bounds() []float64
	Area() float64
	Length() float64
	GeoType() int
	Center() []float64
	Valid() bool
}

// geoType is an integer that represents the type of geometry object
// 0: Point
// 1: polygon
// 2: linearRing
// 3: linearStrings
// 4: MultiLine
// 5: MultiPolygon
