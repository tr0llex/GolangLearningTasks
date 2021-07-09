package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 5, 2, 6, 6, 2, 8, 9, 0} // sorted: [0 1 2 2 5 6 6 8 9]
	fmt.Println(sort.SearchInts(a, 1))
}
