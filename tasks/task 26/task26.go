package main

import "fmt"

func reverse(a string) string {
	slice := []rune(a)
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return string(slice)
}

func main() {
	b := "❤💛💔"
	fmt.Println(reverse(b))
}
