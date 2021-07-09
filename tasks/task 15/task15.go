package main

import "fmt"

func main() {
	a := 1
	b := 2

	fmt.Printf("a = %v, b = %v\n", a, b)
	a, b = b, a
	fmt.Printf("a = %v, b = %v", a, b)
}
