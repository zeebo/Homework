package main

import "math"

//define the kernel
func K(t float64) float64 {
	return math.Pow(t, -0.5)
}
