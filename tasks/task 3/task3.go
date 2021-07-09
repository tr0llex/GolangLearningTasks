package main

import (
	"fmt"
	"runtime"
	"time"
)

func squareNum(num int, c chan int) {
	c <- num * num
}

func main() {
	runtime.GOMAXPROCS(0)

	myArray := [...]int{2, 4, 6, 8, 10}
	c := make(chan int)

	for _, v := range myArray {
		go squareNum(v, c)
	}

	var result int
	for range myArray {
		result += <-c
	}

	fmt.Println(result)

	time.Sleep(1 * time.Second)
}
