package main

import (
	"fmt"
)

func main() {

	max := max(13, 45, 87, 15, 27, 18)

	fmt.Println(max)
}

func max(nrs ...int) int {
	max := nrs[0]
	for _, n := range nrs {
		if n > max {
			max = n
		}
	}
	return max
}
