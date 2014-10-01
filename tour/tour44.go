package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	f1, f2, f3 := 0, 0, 0
	return func(i int) int {
		if i == 0 {
			// nothing
		} else if i == 1 {
			f3 = i
		} else {
			f1 = f2
			f2 = f3
			f3 = f1 + f2
		}
		return f3
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
