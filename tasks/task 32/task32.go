package main

import (
	"fmt"
	"time"
)

func Sleep(d time.Duration) {
	t := time.NewTimer(d)
	<-t.C
}

func main() {
	fmt.Println("3 сентября")
	Sleep(time.Second * 3)
	fmt.Println("И снова 3 сентября")
}
