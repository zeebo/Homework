package main

import (
	"math"
	"math/big"
	"strconv"
)

//make a cache of the first 10000 values

var memo [10000]float64

func load(n int64) {
	r := big.NewRat(1, 1)
	o := big.NewRat(1, 1)
	var fl float64
	for i := int64(1); i <= n; i++ {
		o.SetFrac64(2*i-1, 2*i)
		r.Mul(r, o)

		fl, _ = strconv.ParseFloat(r.FloatString(16), 64)
		memo[i-1] = fl
	}
}

//taylor coeffiencts of laplace transform of kernel
func omega(n int64, h float64) float64 {
	return math.Sqrt(math.Pi*h) * memo[n]
}
