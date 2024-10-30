package functions

func Variance(n []int) float64 {
	average := Average(n)
	var result float64
	for _, number := range n {
		diff := float64(number) - average
		result += diff * diff
	}
	return float64((int(result) / (len(n))))
}
