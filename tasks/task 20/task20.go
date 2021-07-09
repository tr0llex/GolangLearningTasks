package main

import "fmt"

func main() {
	slice := []string{"a", "a"}
	// передаем в анонимную функцию слайс по значению => он не меняется вне функции
	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
}
