package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	a := strings.Fields(s)
	m := make(map[string]int, len(a))
	for _, v := range a {
		m[v] = m[v] + 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
