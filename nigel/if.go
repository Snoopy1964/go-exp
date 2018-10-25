package main

import (
	"fmt"
)

func main() {

	if firstRank, secondRank := 39, 614; firstRank < secondRank {
		fmt.Println("\nFirst course is doing better " +
			"than second course")
	} else if firstRank > secondRank {
		fmt.Println("Oh my God.... you are sooo good!")
	} else {
		fmt.Println("Hmmm, seems to be that we are on the same rank,.... or something weird is going on")
	}
}
