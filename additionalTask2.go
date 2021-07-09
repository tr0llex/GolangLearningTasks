// Дан массив из 10 строк, посчитать в каждой строке количество уникальных символов с помощью горутин

package main

import (
	"fmt"
	"runtime"
	"time"
)

const workersNum = 10

func countUniqChars(a string) int {
	temp := []rune(a)
	countTable := make(map[rune]int)
	for _, v := range temp {
		countTable[v]++
	}
	return len(countTable)
}

func startWorker(workerNum int, in <-chan string) {
	for input := range in {
		uniqSymbols := countUniqChars(input)
		fmt.Printf("Воркер: %v\nСтрока: %v\nКоличествово уникальных символов: %v\n\n",
			workerNum, input, uniqSymbols)
		runtime.Gosched()
	}
	fmt.Printf("Воркер №%v завершил работу\n\n", workerNum)
}

func main() {
	array := [10]string{
		"AAAAAAAAA",     // 1 unique char
		"ABC",           // 3 unique chars
		"ABBCC",         // 3 unique chars
		"ABCXYZ",        // 6 unique chars
		"ABABABABAB",    // 2 unique chars
		"ABCDEFGHIJKLM", // 13 unique chars
		"Alexey",        // 5 unique chars
		"GoLang",        // 6 unique chars
		"GO",            // 2 unique chars
		"Alexey Golang", // 11 unique chars
	}
	dataCh := make(chan string)
	for i := 0; i < workersNum; i++ {
		go startWorker(i, dataCh)
	}
	for _, v := range array {
		dataCh <- v
	}
	close(dataCh)

	time.Sleep(time.Millisecond)
}
