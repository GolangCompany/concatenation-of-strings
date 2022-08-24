package main

import "fmt"

func main() {

	s1 := "Fourth"
	s2 := " Concatenation"
	s3 := " Method"

	output := fmt.Sprintf("%s%s%s", s1,
		s2, s3)

	fmt.Println(output)
}
