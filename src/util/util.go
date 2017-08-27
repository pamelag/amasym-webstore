package util

import "math"

func Round(val float64) (newVal float64) {
	
	var round float64
	
	pow := math.Pow(10, float64(2))
	digit := pow * val
	
	_, div := math.Modf(digit)
	
	if div >= 0.5 {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	
	newVal = round / pow
	return
}
