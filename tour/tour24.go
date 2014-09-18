package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func Sqrt(x float64) float64 {
	c := 0
	a := float64(1)
	z := float64(1)
	for {
		c += 1
		if z = z - (z*z-x)/(2*x); a == z {
			break
		} else {
			a = z
		}
	}
	fmt.Printf("num of calc:%d\n", c)
	return a
}

func main() {
	f := float64(2)
	if len(os.Args) > 1 {
		f, _ = strconv.ParseFloat(os.Args[1], 64)
	}
	fmt.Printf("this func:%v\n", Sqrt(f))
	fmt.Printf("math pkg :%v\n", math.Sqrt(f))
}
