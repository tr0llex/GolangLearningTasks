package main

import "fmt"

func set(array []string) []string {
	table := make(map[string]bool)
	for _, v := range array {
		table[v] = true
	}
	var result []string
	for key := range table {
		result = append(result, key)
	}
	return result
}

func main() {
	array := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Printf("Array: %v\n", array)
	fmt.Printf("Set:   %v", set(array))
}
