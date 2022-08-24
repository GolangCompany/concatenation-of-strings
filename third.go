package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"Third", " Concatenation", " Method"}
	t := strings.Join(s, "")
	fmt.Println(t)
}
