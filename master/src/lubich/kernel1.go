package main

import "math"

func load(n int64) {}

//taylor coeffiencts of laplace transform of kernel
func omega(n int64, h float64) float64 {
	return h * math.Pow(1+h, float64(-1-n))
}
