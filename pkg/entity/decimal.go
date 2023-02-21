package entity

import "math"

func FormatDecimal(value float64) float64 {
	return math.Round(value*100) / 100
}
