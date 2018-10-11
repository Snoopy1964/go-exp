package main

import (
	"fmt"
	"os"
)

func main() {

	name := os.Getenv("USERNAME")
	course := "Docker Deep Dive"

	fmt.Println("\nHi", name, "you are currently watching", course)

	// c1 := changeCourse(&course)
	changeCourse(&course)

	fmt.Println("\nyou are now currently watching", course)
	// fmt.Println("\nyou are now currently watching", c1)

}

func changeCourse(course *string) {

	*course = "First Look: Native Docker Clustering"
	fmt.Println("\nTrying to change your course to", *course)
	//return *course

}
