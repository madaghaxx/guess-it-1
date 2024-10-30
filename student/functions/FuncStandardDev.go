package functions
import "math"

func StandardDeviation(n []int) float64{
	return math.Sqrt(Variance(n))
}