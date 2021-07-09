package main

import "fmt"

func update(p *int) {
	b := 2
	p = &b
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p) // выведет 1
	update(p)       // в функции update создастся копия переменной, т.к. мы передаем не адрес &p
	fmt.Println(*p) // выведет 1
}
