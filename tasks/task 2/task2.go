package main

import (
	"fmt"
	"runtime"
	"time"
)

type indexValuePair struct {
	index int
	value int
}

func square(pair indexValuePair, c chan indexValuePair) {
	pair.value = pair.value * pair.value
	c <- pair
}

func main() {
	runtime.GOMAXPROCS(0)

	myArray := [...]int{2, 4, 6, 8, 10}
	c := make(chan indexValuePair)

	for i, v := range myArray {
		temp := indexValuePair{index: i, value: v}
		go square(temp, c)
	}

	for range myArray {
		temp := <-c
		myArray[temp.index] = temp.value
	}
	fmt.Println(myArray)

	time.Sleep(1 * time.Second)
}
