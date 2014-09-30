package main

// https://code.google.com/p/go-tour/source/browse/pic/pic.go
import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for y := range pic {
		pic[y] = make([]uint8, dx)
		for x := range pic[y] {
			pic[y][x] = uint8(y ^ x)
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
