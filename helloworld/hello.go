package main

import (
	"fmt"

	"go-exp/helloworld/greeting"
)

// Does not work, but in Training it works -> IDE ????
// import "greeting"

func rename2Frog(r greeting.Renamable) {
	r.Rename("Frog")
}

func main() {

	// var s = greeting.Salutation{Name: "Bob", Greeting: "hallo!"}

	var si []int
	si = make([]int, 3)

	si[0] = 1
	si[1] = 10
	si[2] = 100
	// si[3] = 1000

	// short form: s := []int{1, 10, 500, 250}

	salutations := greeting.Salutations{
		{"Bob", "hallo"},
		{"Ralf", "Hi"},
		{"Aileen", "Ho, ho, ho"},
		{"Marie", "What's up???"},
		{"Mary", "What's up?"},
	}

	// salutations = append(slice, greeting.Salutation{"Frank", "Hi"})
	// slice = append(slice, slice...)
	// deleting slices
	// ??? does not work, why????
	// salutations = append(salutations[:1], salutations[2:])
	/*
		salutations[0].Rename("John")
		rename2Frog(&salutations[4])
	*/

	// fmt.Fprintf(&salutations[0], "The count is %d", 10)

	c1 := make(chan greeting.Salutation)
	c2 := make(chan greeting.Salutation)

	// some method to put a Salutation TYpe into channel
	go salutations.ChannelGreeter(c1)
	go salutations.ChannelGreeter(c2)

	for {
		select {
		case s, ok := <-c1:
			if ok {
				fmt.Println(s.Name, ":1")
			} else {
				return
			}

		case s, ok := <-c2:
			if ok {
				fmt.Println(s.Name, ":2")
			} else {
				return
			}
		default:
			fmt.Println("Waiting...")

		}
	}
	// done := make(chan bool, 2)
	// go func() {
	// 	salutations.Greet(greeting.CreatePrintFunction("<c>"), true, 5)
	// 	done <- true
	// 	time.Sleep(100 * time.Millisecond)
	// 	done <- true
	// 	fmt.Println("Done!")
	// }()
	// salutations.Greet(greeting.CreatePrintFunction("!!!"), true, 5)
	// // greeting.Greet(s, greeting.Print, true)
	// // time.Sleep(100 * time.Millisecond)
	// <-done
}
