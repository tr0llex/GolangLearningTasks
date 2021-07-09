package main

import (
	"fmt"
	"runtime"
	"time"
)

const workersNum = 4

func startWorker(workerNum int, in <-chan interface{}) {
	for input := range in {
		fmt.Printf("Воркер: %v, Данные: %v\n", workerNum, input)
		runtime.Gosched()
	}
	fmt.Printf("Воркер №%v завершил работу\n", workerNum)
}

func main() {
	runtime.GOMAXPROCS(0)
	worketInput := make(chan interface{})

	for i := 0; i < workersNum; i++ {
		go startWorker(i, worketInput)
	}

	data := []interface{}{
		"Текст", "Дом", "Каракули", "Абрикос", 2347373, "Апельсин", "Машина", "Картина", "Дерево",
		"Галактика", "Монитор", "Стол", "Лампа", "Берёза", "Футболка", 43.23,
	}

	for _, v := range data {
		worketInput <- v
	}
	close(worketInput)

	time.Sleep(time.Millisecond)
}
