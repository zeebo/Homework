package main

import (
	"flag"
	"fmt"
)

//declare some variables
var (
	n, r, i     int64
	top, sum    float64
	mesh, c, km []float64
)

func main() {
	//parse some command line flags
	flag.Float64Var(&top, "top", 1, "top value")
	flag.Int64Var(&n, "n", 10000, "N")
	flag.Parse()

	//make arrays of the size we need
	mesh = make([]float64, n+1)
	c = make([]float64, n)
	km = make([]float64, n)

	//find the stepsize
	h := top / float64(n)

	//fill in the mesh
	for i = 0; i < n+1; i++ {
		mesh[i] = top * float64(i) / float64(n)
	}

	//calculate the integral of the kernel over the mesh
	for i = 0; i < n; i++ {
		km[i] = K(mesh[i+1]) - K(mesh[i])
	}

	//run the scheme
	for r = 0; r < n; r++ {
		sum = f(mesh[r+1])
		for i = 0; i < r; i++ {
			sum -= c[i] * km[r-i]
		}
		c[r] = sum / km[0]
	}

	//print the results
	for i = 0; i < n; i++ {
		fmt.Printf("%.16f\t", mesh[i]+(h/2))
		fmt.Printf("%.16f\n", c[i])
	}
}
