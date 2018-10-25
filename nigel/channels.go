package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	go func() {
		defer waitGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("hello")
	}()

	go func() {
		defer waitGrp.Done()
		fmt.Println("Pluralsight")

	}()

	waitGrp.Wait()
}
