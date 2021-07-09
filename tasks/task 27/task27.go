package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "snow dog sun"
	afterSplit := strings.Split(text, " ")
	var result []string
	for i := len(afterSplit) - 1; i >= 0; i-- {
		result = append(result, afterSplit[i])
	}
	fmt.Println(result)
}
