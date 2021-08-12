// Написать калькулятор, решающий обратную бесскобочную запись: "5 7 + 3 -"

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type stack struct {
	buffer []int
}

func NewStack() *stack {
	return &stack{make([]int, 0)}
}

func (s *stack) Push(v int) {
	s.buffer = append(s.buffer, v)
}

func (s *stack) Pop() (int, bool) { // Возвращает bool=true если стек пустой
	bufLen := len(s.buffer)
	if bufLen == 0 {
		return 0, true
	}
	res := s.buffer[bufLen-1]
	s.buffer = s.buffer[:bufLen-1]
	return res, false
}

func main() {
	stack := NewStack()
	input := "5 7 + 3 -"
	calcSlice := strings.Split(input, " ")
	var a, b int
	var stackIsEmpty = false

	for _, v := range calcSlice {
		switch v {
		case "+":
			a, stackIsEmpty = stack.Pop()
			b, stackIsEmpty = stack.Pop()
			stack.Push(a + b)
		case "-":
			a, stackIsEmpty = stack.Pop()
			b, stackIsEmpty = stack.Pop()
			stack.Push(b - a)
		case "*":
			a, stackIsEmpty = stack.Pop()
			b, stackIsEmpty = stack.Pop()
			stack.Push(a * b)
		case "/":
			a, stackIsEmpty = stack.Pop()
			b, stackIsEmpty = stack.Pop()
			if a == 0 {
				fmt.Println("Error! Division by zero!")
				return
			} else {
				stack.Push(b / a)
			}
		default:
			number, _ := strconv.Atoi(v)
			stack.Push(number)
		}
		if stackIsEmpty {
			fmt.Println("Error! Incorrect input!")
			return
		}
	}
	result, _ := stack.Pop()
	fmt.Println(result)
}
