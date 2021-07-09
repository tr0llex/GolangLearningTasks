package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []string{"Виктор", "Alexey", "Семён", "Алексей", "Юлия", "Анастасия", "Anastasia"}
	sort.Strings(a)
	fmt.Println(a)
}
