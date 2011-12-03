package main

import (
	"flag"
	"fmt"
	"math"
)

//define some command line arguments
var (
	k   = flag.Float64("k", 0.001, "stepsize")
	top = flag.Float64("top", 1, "top value")
)

//define some variables for the scheme
var (
	mesh, u []float64
	i, j    int64
	sum     float64
)

//taylor coeffiencts of laplace transform of kernel
func beta(n int64) float64 {
	return math.Pow(1.0/(*k+1), float64(n+1))
}

//forcing function
func f(t float64) float64 {
	return t * math.Exp(-1*t)
}

func main() {
	flag.Parse()

	//calculate the number of mesh points
	n := int64(math.Ceil(*top / *k))

	//create the mesh and solution arrays
	mesh = make([]float64, n)
	u = make([]float64, n)

	//calculate the mesh
	for i = 0; i < n; i++ {
		mesh[i] = *top * float64(i) / float64(n)
	}

	//run the scheme
	for i = 1; i < n; i++ {
		sum = f(mesh[i])
		for j = 1; j < i; j++ {
			sum -= beta(i-j) * u[j]
		}
		u[i] = sum / beta(0)
	}

	//print the results
	for i = 1; i < n; i++ {
		fmt.Printf("%.08f\t", mesh[i])
		fmt.Printf("%.08f\n", u[i])
	}
}
