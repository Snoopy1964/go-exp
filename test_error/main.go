package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	var err0 error
	var err1 amqp.Error
	var err2 *amqp.Error

	fmt.Printf("error  .......: %T, %v\n", err0, err0)
	fmt.Printf("amqp.Error ...: %T, %v\n", err1, err1)
	fmt.Printf("*amqp.Error ..: %T, %v\n", err2, err2)
}
