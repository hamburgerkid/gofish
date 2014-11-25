package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	z := x
	y := z
	for {
		z = z - ((z*z*z - x) / (3 * z * z))
		if y == z {
			break
		} else {
			y = z
		}
	}
	return z
}

func main() {
	x := complex128(2)
	fmt.Println(Cbrt(x))
	fmt.Println(cmplx.Pow(Cbrt(x), 3))
}
