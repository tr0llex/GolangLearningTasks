package main

import "fmt"

var justString string

func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ {

		v += "A"
	}
	return v
}

func someFunc() {
	v := createHugeString(1 << 10)
	fmt.Printf("%p\n", &v)
	justString = v[:100]
	fmt.Printf("%p\n", &justString)
}

func main() {
	someFunc()
}
