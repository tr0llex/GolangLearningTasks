package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)
	a.SetString("1267650600228229401496703205376", 10)
	b.SetString("1732349399771770598503296794624", 10)
	fmt.Println("A =", a)
	fmt.Println("B =", b)
	fmt.Printf("Частное: %v\n", big.NewInt(0).Div(a, b))
	fmt.Printf("Произведение: %v\n", big.NewInt(0).Mul(a, b))
	fmt.Printf("Сумма: %v\n", big.NewInt(0).Add(a, b))
	fmt.Printf("Разность: %v\n", big.NewInt(0).Sub(a, b))
}
