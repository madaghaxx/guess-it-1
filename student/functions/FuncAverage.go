package functions

func Average(n []int) float64 {
	var res float64
	for _, num := range n {
		res += float64(num)
	}
	return res / float64(len(n))
}
