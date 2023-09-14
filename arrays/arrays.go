package main

import "fmt"

func main() {
	// Array's type = length + type BOTH
	// Values are all zero-valued by default, which is so nice

	var a [5]int
	fmt.Println(a)

	a[4] = 100
	fmt.Println(a) // Zero-indexed just like in C

	fmt.Println(len(a))

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	// Can drop length - length is inferred here
	var c = [...]int{1, 2, 3, 4, 5} // Declare + initialize an array
	fmt.Println(c)

	x := [5]int{3: 200} // Initialize 5-len array with element index 3 = 200
	fmt.Println(x)

	y := [4]string{"a b", "c d", "e f", "g h"}
	fmt.Println(y)
	fmt.Printf("%q\n", y)
}
