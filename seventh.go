package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 strings.Builder
	s1.WriteString("Seventh")
	s1.WriteString(" Concatenation")
	s1.WriteString(" Method")
	fmt.Println(s1.String())
}
