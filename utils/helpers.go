package utils

import "math"

func Sum(num ...float64) float64 {
	var sum float64
	for _, v := range num {
		sum += v
	}
	return sum
}

func Min(num ...float64) float64 {
	min := math.Inf(1)
	for _, v := range num {
		if v < min {
			min = v
		}
	}
	return min
}

func Max(num ...float64) float64 {
	max := math.Inf(-1)
	for _, v := range num {
		if v > max {
			max = v
		}
	}
	return max
}
