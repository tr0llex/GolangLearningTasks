package main

import "fmt"

func RemoveIndex(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	a := []int{1, 5, 7, 3, 2}
	fmt.Println(a)
	a = RemoveIndex(a, 2)
	fmt.Println(a)
}
