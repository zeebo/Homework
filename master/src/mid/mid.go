package main

import (
	"flag"
	"fmt"
)

var (
	n, r, i     int64
	top, sum    float64
	mesh, x, km []float64
)

func main() {
	//parse some command line flags
	flag.Float64Var(&top, "top", 1, "top value")
	flag.Int64Var(&n, "n", 10000, "N")
	flag.Parse()

	//make arrays of the size we need
	mesh = make([]float64, n)
	x = make([]float64, n)
	km = make([]float64, n)

	//find the stepsize
	h := top / float64(n)

	//fill in the mesh at half steps
	for i = 0; i < n; i++ {
		mesh[i] = top*float64(i)/float64(n) + (h / 2)
	}

	//calculate the kernel at the mid points
	for i = 0; i < n; i++ {
		km[i] = K(mesh[i])
	}

	//run the scheme
	for r = 0; r < n; r++ {
		sum = f(mesh[r]+(h/2)) / h
		for i = 0; i < r; i++ {
			sum -= km[r-i] * x[i]
		}
		x[r] = sum / km[0]
	}

	//print the results
	for i = 0; i < n; i++ {
		fmt.Printf("%.08f\t", mesh[i])
		fmt.Printf("%.08f\n", x[i])
	}
}
