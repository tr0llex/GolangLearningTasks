package main

import "fmt"

func reverse(a string) string {
	slice := []rune(a)
	var temp []rune
	for i := len(slice) - 1; i >= 1; i -= 2 {
		temp = append(temp, slice[i], slice[i-1])
	}
	if len(slice)%2 == 1 {
		temp = append(temp, slice[0])
	}
	return string(temp)
}

func main() {
	b := "❤💛💔"
	fmt.Println(reverse(b))

}
