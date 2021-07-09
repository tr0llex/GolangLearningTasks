package main

import "fmt"

func main() {
	b := make(chan int)
	a := []interface{}{123, "Hello", 678, b, false}
	for _, v := range a {
		fmt.Println(v)
		switch v.(type) {
		case int:
			fmt.Printf("Это int\n\n")
		case string:
			fmt.Printf("Это string\n\n")
		case chan int:
			fmt.Printf("Это channel (int)\n\n")
		case bool:
			fmt.Printf("Это bool\n\n")
		}
	}
}
