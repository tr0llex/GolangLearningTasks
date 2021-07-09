package main

import (
	"context"
	"fmt"
	"time"
)

type SumCalculation struct {
	operand1 int
	operand2 int
	result   int
}

func contextUsage(ctx context.Context, dataCh chan SumCalculation) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина останавливается через контекст")
			return
		case calculation := <-dataCh:
			calculation.result = calculation.operand1 + calculation.operand2
			dataCh <- calculation
		}
	}
}

// Еще можно передавать в горутину один канал, делать по нему range и закрыть его
func channelUsage(cancelCh chan bool, dataCh chan int) {
	for {
		select {
		case <-cancelCh:
			fmt.Println("Горутина останавливается через канал\n")
			close(dataCh)
			return
		case b := <-dataCh:
			fmt.Println(b)
		}
	}
}

func main() {
	// Остановка горутины через канал отмены
	cancelCh := make(chan bool)
	dataCh := make(chan int)
	go channelUsage(cancelCh, dataCh)

	for i := 0; i < 5; i++ {
		dataCh <- i
		time.Sleep(time.Millisecond * 200)
	}
	cancelCh <- true

	// Остановка горутины через контекст
	sumCh := make(chan SumCalculation)
	ctx, finish := context.WithCancel(context.Background())
	go contextUsage(ctx, sumCh)
	calculation := SumCalculation{operand1: 111, operand2: 222}
	sumCh <- calculation
	calculation = <-sumCh
	finish()
	time.Sleep(time.Second)
	fmt.Printf("%v + %v = %v\n", calculation.operand1, calculation.operand2, calculation.result)
}
