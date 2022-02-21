package utils

import "math"

func ExtractDecimals(num float64, decimals int) int {
	return int(num * math.Pow(10, float64(decimals)))
}

func AddDecimals(num int, decimals int) float64 {
	var dec = math.Pow(10, float64(decimals))
	var numberWithDecimals = float64(num) / dec
	return numberWithDecimals
}
