package functions

func Median(n []int) float64 {
	sorted := (BubbleSort(n))
	var res float64
	if len(sorted)%2 == 0 {
		res = float64((sorted[len(sorted)/2] + sorted[len(sorted)/2+1]) / 2)
	}
	if len(sorted)%2!=0{
		res = float64(sorted[len(sorted)/2])
	}
	return res
}

func BubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
