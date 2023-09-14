package main

import "fmt"

func main() {
	// Find sum of first 100 numbers using a for loop
	// Unlike C, no parantheses required
	var sum int
	for i := 1; i <= 100; i++ {
		sum += i
	}

	fmt.Println("sum = ", sum)
}
