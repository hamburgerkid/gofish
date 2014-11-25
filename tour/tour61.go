package main

import (
	"bytes"
	"io"
	"os"
	"strings"
)

var alphabet_upper = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var alphabet_upper_len = len(alphabet_upper)
var alphabet_lower = []byte("abcdefghijklmnopqrstuvwxyz")
var alphabet_lower_len = len(alphabet_lower)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}
	return n, err
}

func rot13(b byte) byte {
	pos := bytes.IndexByte(alphabet_lower, b)
	if pos != -1 {
		return alphabet_lower[(pos+13)%alphabet_lower_len]
	}
	pos = bytes.IndexByte(alphabet_upper, b)
	if pos != -1 {
		return alphabet_upper[(pos+13)%alphabet_upper_len]
	}
	return b
}

func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
