package main

import "fmt"

func main() {
	type Human struct {
		firstName string
		lastName  string
		age       int
	}

	type Action struct {
		person     Human
		actionType string
	}

	alex := Human{
		firstName: "Alex",
		lastName:  "Samoylov",
		age:       21,
	}
	workingAlex := Action{
		person:     alex,
		actionType: "работает",
	}

	fmt.Println(workingAlex)
}
