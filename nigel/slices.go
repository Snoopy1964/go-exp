package main

import (
	"fmt"
)

func main() {
	mySlice := make([]int, 1, 4)
	//	fmt.Println(mySlice[4])

	for i := 1; i < 17; i++ {
		mySlice = append(mySlice, 10*i)

		fmt.Printf("Length is: %d Capacity is: %d - ", len(mySlice), cap(mySlice))
		fmt.Println(mySlice)
	}
}
