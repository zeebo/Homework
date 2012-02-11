package main

import (
	"flag"
	"fmt"
	"math"
)

//declare some variables
var (
	n, r, i     int64
	top, sum    float64
	mesh, c, km []float64
)

//define the integral of the kernel
func K(t float64) float64 {
	return -1 * math.Exp(-1*t)
}

//define the function f
func f(t float64) float64 {
	return t * math.Exp(-1*t)
}

func main() {
	//parse some command line flags
	flag.Float64Var(&top, "top", 1, "top value")
	flag.Int64Var(&n, "n", 10000, "N")
	flag.Parse()

	//make arrays of the size we need
	mesh = make([]float64, n)
	c = make([]float64, n)
	km = make([]float64, n)

	//fill in the mesh
	for i = 0; i < n; i++ {
		mesh[i] = top * float64(i) / float64(n)
	}

	//calculate the integral of the kernel over the mesh
	for i = 0; i < n-1; i++ {
		km[i] = K(mesh[i+1]) - K(mesh[i])
	}

	//run the scheme
	for r = 0; r < n; r++ {
		sum = f(mesh[r])
		for i = 1; i < r; i++ {
			sum -= c[r-i] * km[i]
		}
		c[r] = sum / km[0]
	}

	//print the results
	for i = 0; i < n; i++ {
		fmt.Printf("%.08f\t", mesh[i])
		fmt.Printf("%.08f\n", c[i])
	}
}
