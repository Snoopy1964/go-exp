package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("c:\\temp\\test1.txt")

	if err != nil {
		fmt.Println("error returned was: ", err)
	}
	fmt.Println("error returned was: ", err)
}
