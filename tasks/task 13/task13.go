package main

import (
	"fmt"
	"sync"
)

// Программа завершится дедлоком, т.к. WaitGroup нужно передавать по указателю, она одна и не должна копироваться
// Wait - блокирующая операция, поэтому возникает дедлок

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
