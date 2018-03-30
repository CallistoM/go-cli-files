package main

import "math"

// roundInteger return float with copysign
func roundInteger(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

// fixedFloat takes precision and a number
func fixedFloat(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(roundInteger(num*output)) / output
}
