// https://go.dev/tour/moretypes/26

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var a, b int = 0, 1

	return func() int {
		update_fn := func() {
			a, b = b, b+a
		}
		defer update_fn()
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
