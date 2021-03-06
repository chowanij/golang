package main

import (
	"fmt"
)

const precision = 20;

func Sqrt(x float64) float64 {
	z := 1.0
	for i:= 0 ; i <= precision ; i++ {
		z -= (z*z -x) / (2*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
