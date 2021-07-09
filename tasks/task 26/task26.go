package main

import "fmt"

func reverse(a string) string {
	slice := []rune(a)
	var temp []rune
	for i := len(slice) - 1; i >= 0; i-- {
		temp = append(temp, slice[i])
	}
	return string(temp)
}

func main() {
	b := "❤💛💔"
	fmt.Println(reverse(b))

}
