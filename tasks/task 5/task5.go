package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, workerNum int, out chan<- string) {
	for {
		time.Sleep(100 * time.Millisecond)
		timeNow := time.Now().String()
		fmt.Printf("+ Воркер №%v написал: %v\n", workerNum, timeNow)
		out <- timeNow
	}
}

func main() {
	N := 1
	workersNum := 1
	workTime := time.Duration(N) * time.Second
	ctx, _ := context.WithTimeout(context.Background(), workTime) //
	result := make(chan string, 1)

	for i := 1; i <= workersNum; i++ {
		go worker(ctx, i, result)
	}

	totalWritten := 0
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		case readResult := <-result:
			totalWritten++
			fmt.Println("- Прочитано", readResult)
		}
	}
	fmt.Println("Всего объектов: ", totalWritten)
	time.Sleep(time.Second)
}
