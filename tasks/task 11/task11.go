package main

import "fmt"

func main() {
	a := []int{3, 1, 6, 2, 1, 7, 8, 2}
	b := []int{10, 8, 1, 7, 2, 1}

	fmt.Println(intersection(a, b)) // O(n)
}

func intersection(a []int, b []int) []int {
	var result []int
	mapA := make(map[int]int)
	mapB := make(map[int]int)
	for _, v := range a {
		mapA[v]++
	}
	for _, v := range b {
		mapB[v]++
	}
	for number, timesRepeat := range mapA {
		for i := 0; (i < mapB[number]) && (timesRepeat-1 >= i); i++ {
			result = append(result, number)
		}
	}
	fmt.Println(mapA)
	fmt.Println(mapB)
	return result
}
