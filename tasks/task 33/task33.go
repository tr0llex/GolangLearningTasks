package main

import (
	"fmt"
	"math/rand"
	"time"
)

func evenChecker(randCh <-chan int, evenCh chan<- int) {
	defer close(evenCh)
	for i := range randCh {
		if i%2 == 0 {
			evenCh <- i
		}
	}
}

func printer(evenCh <-chan int) {
	for i := range evenCh {
		fmt.Println(i)
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	randCh := make(chan int)
	evenCh := make(chan int)
	go evenChecker(randCh, evenCh)
	go printer(evenCh)
	for i := 0; i < 20; i++ {
		randCh <- rand.Intn(100)
	}
	close(randCh)
	time.Sleep(time.Millisecond * 10)
}
