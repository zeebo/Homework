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

	//fill in the mesh
	for i = 0; i < n; i++ {
		mesh[i] = top * float64(i) / float64(n)
	}

	//calculate the kernel at the left endpoints
	for i = 0; i < n; i++ {
		km[i] = K(mesh[i])
	}

	h := top / float64(n)

	//run the scheme
	for r = 0; r < n-1; r++ {
		sum = f(mesh[r+1]) / h
		for i = 0; i < r; i++ {
			sum -= km[r+1-i] * x[i]
		}
		x[r] = sum / km[1]
	}

	//print the results
	for i = 0; i < n-1; i++ {
		fmt.Printf("%.08f\t", mesh[i])
		fmt.Printf("%.08f\n", x[i])
	}
}
