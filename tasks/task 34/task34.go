package main

import "fmt"

func uniqueCheck(a string) bool {
	temp := []rune(a)
	countTable := make(map[rune]int)
	for _, v := range temp {
		countTable[v]++
	}
	for _, v := range countTable {
		if v != 1 {
			return false
		}
	}
	return true
}

func main() {
	a := "ABCDEF"
	b := "ABCDCDCDEF"
	fmt.Println(uniqueCheck(a))
	fmt.Println(uniqueCheck(b))
}
