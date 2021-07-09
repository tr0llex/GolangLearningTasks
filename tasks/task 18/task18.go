package main

import "fmt"

func someAction(v []int8, b int8) {
	fmt.Printf("Адрес v[0] до модификации: %p\n", &v[0])
	fmt.Printf("Адрес слайса до добавления v : %p\n\n", &v)

	v[0] = 100
	v = append(v, b)

	fmt.Printf("Адрес v[0] после модификации: %p\n", &v[0])
	fmt.Printf("Адрес слайса v после добавления: %p\n\n", &v)
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)

	fmt.Printf("Конечный адрес a[0] в main : %p\n", &a[0])
	fmt.Printf("Конечный адрес слайса a в main: %p\n\n", &a)

	fmt.Println(a)

}

// Выведет 100 2 3 4 5, т.к. мы
