package main

import (
	"fmt"
)

func SetBit(number int64, bitPos int, bitValue int) int64 {
	if int((number&(1<<bitPos))>>bitPos) != bitValue {
		if bitValue == 1 {
			number |= 1 << bitPos
		} else {
			number &^= 1 << bitPos
		}
	}
	return number
}

func main() {
	var number int64
	var bitPos, bitValue int

	fmt.Println("Введи через пробел: число, номер бита [0;63] и значение, которое хочешь поместить в бит")
	_, err := fmt.Scanf("%d %d %d", &number, &bitPos, &bitValue)
	if err != nil {
		fmt.Println("Input error")
		return
	}

	number = SetBit(number, bitPos, bitValue)
	fmt.Printf("bin(%v) = %b", number, number)
}
