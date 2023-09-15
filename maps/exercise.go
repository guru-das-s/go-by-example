// https://go.dev/tour/moretypes/23

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	hist := make(map[string]int)
	for _, v := range words {
		hist[v] = hist[v] + 1
	}
	return hist
}

func main() {
	wc.Test(WordCount)
}
