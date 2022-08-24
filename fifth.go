package main

import (
	"bytes"
	"fmt"
)

func main() {
	var s1 bytes.Buffer

	s1.WriteString("Fifth")
	s1.WriteString(" Concatenation")
	s1.WriteString(" Method")
	fmt.Println(s1.String())
}
