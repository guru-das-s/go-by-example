package main

import (
	"fmt"
)

func main() {
	// Switch-cases have an implicit break:
	// If a case matches, other cases are not evaluated

	// Basic calculator program

	var a, b int
	var op string

	fmt.Println("Enter two numbers, and an operation")
	fmt.Scan(&a, &b, &op)
	fmt.Println("Entered:", a, b, op)

	switch op {
	case "+":
		fmt.Println("Result:", a+b)
	case "-":
		fmt.Println("Result:", a-b)
	case "*":
		fmt.Println("Result:", a*b)
	case "/":
		fmt.Println("Result:", a/b)
	default:
		fmt.Println("Unsupported operation")
	}
}
