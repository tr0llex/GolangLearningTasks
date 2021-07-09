package main

import "fmt"

func readArray(a []int, dataCh chan int) {
	for _, v := range a {
		dataCh <- v
	}
	close(dataCh)
}

func main() {
	myArray := [...]int{5, 1, 4, 2, 3}
	dataCh := make(chan int)
	go readArray(myArray[:], dataCh)

	for i := range dataCh {
		fmt.Println(i)
	}
}
