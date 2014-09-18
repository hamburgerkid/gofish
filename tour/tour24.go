package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	a := 1.0
	c := 0
	z := float64(1)
	for {
		c += 1
		if z = z - (z*z-x)/(2*x); a == z {
			break
		} else {
			a = z
		}
	}
	fmt.Println(c)
	return a
}

func main() {
	i := float64(2)
	fmt.Println(Sqrt(i))
	fmt.Println(math.Sqrt(i))
}
