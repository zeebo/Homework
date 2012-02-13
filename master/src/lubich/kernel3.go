package main

import (
	"math"
)

//make a cache of the first 10000 values

var memo [10000]float64

func load(n int64) {
	var fl float64 = 1.0
	for i := int64(1); i <= n; i++ {
		fl *= float64(2*i-1) / float64(2*i)
		memo[i-1] = fl
	}
}

//taylor coeffiencts of laplace transform of kernel
func omega(n int64, h float64) float64 {
	return math.Sqrt(math.Pi*h) * memo[n]
}
