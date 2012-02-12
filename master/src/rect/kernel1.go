package main

import "math"

//define the kernel
func K(t float64) float64 {
	return math.Exp(-1 * t)
}
