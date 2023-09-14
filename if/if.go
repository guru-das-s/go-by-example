package main

import "fmt"

func main() {
	// Parantheses are not required for if conditions
	// But curly brackets ARE.
	var num int = 7
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// You can even declare variables before the condition
	// itself. These will have scope ONLY across all branches
	// of the if condition.
	if n := 10; n > 11 {
		fmt.Println(n, "is greater than 11")
	} else {
		fmt.Println(n, "is lesser than 11")
	}

	// Next line will cause an error!
	// fmt.Println(n)
}
