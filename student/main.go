package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var num int
	arr := []int{}

	for i := 1; i <= 12500; i++ {
		fmt.Fscan(os.Stdin, &num)
		arr = append(arr, num)

		min := Average(arr) - Average(arr)
		max := Average(arr) + Average(arr)
		fmt.Println(min, max)
	}
}

func Average(n []int) float64 {
	var res float64
	for _, num := range n {
		res += float64(num)
	}
	return res / float64(len(n))
}

func Variance(n []int) float64 {
	average := Average(n)
	var result float64
	for _, number := range n {
		diff := float64(number) - average
		result += diff * diff
	}
	return float64((int(result) / (len(n))))
}

func Median(n []int) float64 {
	sorted := (BubbleSort(n))
	var res float64
	if len(sorted)%2 == 0 {
		res = float64((sorted[len(sorted)/2] + sorted[len(sorted)/2+1]) / 2)
	}
	if len(sorted)%2 != 0 {
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

func StandardDeviation(n []int) float64 {
	return math.Sqrt(Variance(n))
}
