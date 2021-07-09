package main

import (
	"fmt"
)

func worker(multiplyCh chan<- int, numCh <-chan int) {
	for number := range numCh {
		multiplyCh <- number * 2
	}
}

func main() {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	numCh := make(chan int)
	multiplyCh := make(chan int, len(numbers))
	go worker(multiplyCh, numCh)

	for _, value := range numbers {
		numCh <- value
	}
	close(numCh)

	for i := range numbers {
		fmt.Println(numbers[i], "* 2 =", <-multiplyCh)
	}
}
