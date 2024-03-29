package main

import "math"

func load(n int64) {}

//taylor coeffiencts of laplace transform of kernel
func omega(n int64, h float64) float64 {
	return .5 * h * (math.Pow(1-h, float64(-1-n)) + math.Pow(1+h, float64(-1-n)))
}
