package main

import "fmt"

// Программа выведет 0, т.к. внутри ветвления своя область видимости => "n := 1" - "n" уже новая переменная внутри if

func main() {
	n := 0
	if true {
		n := 1
		n++
	}
	fmt.Println(n)
}
