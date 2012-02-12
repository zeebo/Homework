package main

import "math"

//define the kernel
func K(t float64) float64 {
	return math.Exp(-1 * t)
}

//define the function f
func f(t float64) float64 {
	return t * math.Exp(-1*t)
}
