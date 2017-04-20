package math

// Returns minimum value from given slice
func Min(xs []float64) float64 {
	min := xs[0]
	for _, v := range xs {
		if v < min {
			min = v
		}
	}
	return min
}

// Returns maximum value from given slive
func Max(xs []float64) float64 {
	max := 0.0
	for _, v := range xs {
		if v > max {
			max = v
		}
	}
	return max
}

// Returns average of slice
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}
