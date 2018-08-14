package player

import (
	"math"
)

const (
	defaultCoefForUCB = math.Sqrt2
	ucbMax            = 999999
)

// Calculate UCB (Upper Confidence Bound)
func calcUcb(sum, num, totalNum int, coef float64) float64 {
	if num == 0 || totalNum == 0 {
		return ucbMax
	}

	// average + confidence
	return float64(sum)/float64(num) +
		coef*(math.Sqrt(math.Log(float64(totalNum)))/float64(num))
}

func calcUcbDefault(sum, num, totalNum int) float64 {
	return calcUcb(sum, num, totalNum, defaultCoefForUCB)
}
