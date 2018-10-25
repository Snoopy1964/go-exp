package main

import (
	"fmt"
)

func main() {
	courseInProg := []string{"Docker Deep Dive", "Docker Clustering", "Docker and Kubernetes"}

	for i, c := range courseInProg {
		fmt.Println(i, " : ", c)
	}
}
