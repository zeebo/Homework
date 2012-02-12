package main

import (
	"flag"
	"fmt"
)

//define some variables for the scheme
var (
	n, r, i     int64
	sum, top, h float64
	mesh, u, om []float64
)

func main() {
	flag.Float64Var(&top, "top", 1, "top value")
	flag.Int64Var(&n, "n", 10000, "N")
	flag.Parse()

	//create the mesh and solution arrays
	mesh = make([]float64, n)
	u = make([]float64, n)
	om = make([]float64, n)
	h = top / float64(n)

	//calculate the mesh
	for i = 0; i < n; i++ {
		mesh[i] = top * float64(i) / float64(n)
	}

	//memoize the omega coefficients
	for i = 0; i < n; i++ {
		om[i] = omega(i, h)
	}

	//run the scheme
	for r = 0; r < n; r++ {
		sum = f(mesh[r])
		for i = 0; i < r; i++ {
			sum -= om[r-i] * u[i]
		}
		u[r] = sum / om[0]
	}

	//print the results
	for i = 0; i < n-1; i++ {
		fmt.Printf("%.16f\t", mesh[i])
		fmt.Printf("%.16f\n", u[i+1])
	}
}
