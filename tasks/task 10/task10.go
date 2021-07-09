package main

import "fmt"

func main() {
	initial := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[int][]float32)

	for _, num := range initial {
		decade := int(num/10) * 10
		groups[decade] = append(groups[decade], num)
	}

	fmt.Println(groups)
}
