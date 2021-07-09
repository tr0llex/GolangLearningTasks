package main

import (
	"fmt"
	"sync"
	"time"
)

type syncMap struct {
	sync.Mutex
	dict map[int]int
}

func worker(data *syncMap, workerId int) {
	for key := 0; key < 10; key++ {

		data.Lock()
		data.dict[key] = workerId
		data.Unlock()

		fmt.Printf("Воркер %v написал по ключу %v значение %v\n", workerId, key, workerId)
	}
}

func main() {
	data := syncMap{
		Mutex: sync.Mutex{},
		dict:  make(map[int]int),
	}

	for i := 0; i < 10; i++ {
		go worker(&data, i)
	}

	time.Sleep(2 * time.Second)

	fmt.Println()
	data.Lock()
	for key, value := range data.dict {
		fmt.Printf("ключ: %v\tзначение: %v\n", key, value)
	}
	data.Unlock()
}
