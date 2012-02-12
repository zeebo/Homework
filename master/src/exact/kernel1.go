package main

import "math"

//define the integral of the kernel
func K(t float64) float64 {
	return -1 * math.Exp(-1*t)
}
