package main

import "golang.org/x/tour/pic"

// https://go.dev/tour/moretypes/18

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)

	for i := range pic {
		pic[i] = make([]uint8, dx)
		for j := range pic[i] {
			pic[i][j] = uint8(i) * uint8(j)
		}
	}
	return pic
}

func Pic2(dx, dy int) [][]uint8 {
	ss := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		s := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			s[x] = uint8(x * y)
		}
		ss[y] = s
	}
	return ss
}

func main() {
	pic.Show(Pic2)
}
