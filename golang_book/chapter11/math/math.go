package math

// Finds the average of a series of numbers
func Average(xs []float64) float64 {
	total := float64(0)
	for _, value := range xs {
		total += value
	}
	return total / float64(len(xs))
}

// Finds the average of a series of numbers

// Finds the minimum value from the slice xs
func Min(xs []float64) float64 {
	min := xs[0]
	for _, value := range xs {
		if value < min {
			min = value
		}
	}
	return min
}

// Finds the maximum value from the slice xs
func Max(xs []float64) float64 {
	max := xs[0]
	for _, value := range xs {
		if value > max {
			max = value
		}
	}
	return max
}
