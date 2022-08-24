package main

import (
	"fmt"
	"strings"
)

func main() {
	string := strings.Builder{}
	a := []byte{66, 117, 105, 108, 100}
	b := []byte{32}
	m := []byte{97}
	d := []byte{32}
	e := []byte{115, 116, 114, 105, 110, 103}

	string.Write(a)
	string.Write(b)
	string.Write(m)
	string.Write(d)
	string.Write(e)

	fmt.Println(string.String())
}
