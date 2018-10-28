package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {

	start := time.Now()
	//fmt.Println("Starttime: ", start)

	runtime.GOMAXPROCS(4)

	f, _ := os.Create("./log.txt")
	f.Close()

	logCh := make(chan string, 50)

	mutexI := make(chan bool, 1)
	mutexJ := make(chan bool, 1)

	go func() {
		for {
			msg, ok := <-logCh
			if ok {
				f, _ := os.OpenFile("./log.txt", os.O_APPEND, os.ModeAppend)
				logTime := time.Now().Format(time.RFC3339)
				f.WriteString(logTime + " - " + msg)
				f.Close()
			} else {
				break
			}
		}
	}()

	for i := 1; i < 5; i++ {
		mutexI <- true
		for j := 1; j < 5; j++ {
			mutexJ <- true
			go func() {
				msg := fmt.Sprintf("%d + %d = %d\n", i, j, i+j)
				logCh <- msg
				fmt.Print(msg)
				<-mutexJ

			}()
			//fmt.Printf("Loop: %d + %d = %d\n", i, j, i+j)
		}
		<-mutexI
	}

	elapsed := time.Since(start)

	fmt.Scanln()
	fmt.Println("Starttime: ", start)
	fmt.Println("runtime: ", elapsed)
}
