package main

import (
	"fmt"
)

func main() {

	type courseMeta struct {
		Author string
		Titel  string
		Level  string
		Rating float64
	}
	DockerForDummies := new(courseMeta)
	DockerForDummies.Author = "Nigel Poulton"
	fmt.Println(*DockerForDummies)

	DockerDeepDive := courseMeta{
		Author: "Nigel Poulton",
		Titel:  "Docker Deep Dive",
		Level:  "Intermediate",
		Rating: 5,
	}
	fmt.Println(DockerDeepDive)
}
