package main

import "math"

//taylor coeffiencts of laplace transform of kernel
func omega(n int64) float64 {
	return h * math.Pow(1+h, float64(-1-n))
}

//forcing function
func f(t float64) float64 {
	return t * math.Exp(-1*t)
}
