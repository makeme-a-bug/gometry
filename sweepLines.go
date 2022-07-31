package geometry

func det(a, b []float64) float64 {
	return (a[0] * b[1]) - (a[1] * b[0])
}

func sort_left_to_right(edges [][][]float64, labels []int) ([][][]float64, []int) {
	for i := 0; i < len(edges); i++ {
		for j := i + 1; j < len(edges); j++ {
			if edges[i][0][0] > edges[j][0][0] {
				edges[i], edges[j] = edges[j], edges[i]
				labels[i], labels[j] = labels[j], labels[i]
			}
		}
	}

	return edges, labels
}

func endPoint_sort_left_to_right(edges [][]float64, labels, direction []int) ([][]float64, []int, []int) {

	for i := 0; i < len(edges); i++ {
		for j := i + 1; j < len(edges); j++ {
			if edges[i][0] > edges[j][0] {
				edges[i], edges[j] = edges[j], edges[i]
				if len(labels) > 0 {
					labels[i], labels[j] = labels[j], labels[i]
				}
				if len(direction) > 0 {
					direction[i], direction[j] = direction[j], direction[i]
				}

			}
		}
	}

	return edges, labels, direction
}

// returns false if not intersecting else ture
// edges are segments
func SweepLinesAlgorithm(coords [][]float64) bool {

	var edges [][][]float64
	var labels []int
	var endpoints [][]float64
	var endpoints_labels []int
	var endpoints_direction []int // 1 for left, -1 for right

	for i := 0; i < len(coords)-1; i++ {
		edge := [][]float64{{coords[i][0], coords[i][1]}, {coords[i+1][0], coords[i+1][1]}}
		edge, _, _ = endPoint_sort_left_to_right(edge, []int{}, []int{})
		edges = append(edges, edge)
		labels = append(labels, i)
		endpoints = append(endpoints, edge...)
		endpoints_labels = append(endpoints_labels, []int{i, i}...)
		endpoints_direction = append(endpoints_direction, []int{1, -1}...)

	}

	edges, labels = sort_left_to_right(edges, labels)
	endpoints, endpoints_labels, endpoints_direction = endPoint_sort_left_to_right(endpoints, endpoints_labels, endpoints_direction)

	var vertical_positions []int

	for i := 0; i < len(edges)*2; i++ {
		label := endpoints_labels[i]

		if endpoints_direction[i] == 1 {
			vertical_positions = insert(vertical_positions, label, edges)
			pred, succ := get_neighbours(vertical_positions, label)
			if pred != -1 {
				if DoIntersect(edges[label][0], edges[label][1], edges[pred][0], edges[pred][1]) {
					return true
				}
			}
			if succ != -1 {
				if DoIntersect(edges[label][0], edges[label][1], edges[succ][0], edges[succ][1]) {
					return true
				}
			}

		} else {

			pred, succ := get_neighbours(vertical_positions, label)
			if pred != -1 && succ != -1 {
				if DoIntersect(edges[pred][0], edges[pred][1], edges[succ][0], edges[succ][1]) {
					return true
				}
			}

			vertical_positions = remove(vertical_positions, label)

		}

	}
	return false

}

func get_neighbours(vP []int, label int) (int, int) {
	pred := -1
	succ := -1
	for i := 0; i < len(vP); i++ {
		if vP[i] == label {
			if i > 0 {
				pred = vP[i-1]
			}
			if i < len(vP)-1 {
				succ = vP[i+1]
			}
		}
	}

	return pred, succ
}

func insert(vP []int, label int, segment [][][]float64) []int {
	position := 0
	for i := 0; i < len(vP); i++ {
		point := segment[label][0]
		position = i + 1
		if point[1] >= segment[vP[i]][0][1] {
			break
		}
	}
	vP = append(vP[:position], append([]int{label}, vP[position:]...)...)
	return vP

}

func remove(vP []int, label int) []int {
	for i := 0; i < len(vP); i++ {
		if vP[i] == label {
			vP = append(vP[:i], vP[i+1:]...)
			break
		}
	}
	return vP
}
