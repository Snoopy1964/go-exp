package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

const msgUrl = "amqp://guest:guest@localhost:5672"

var (
	rabbitConn          *amqp.Connection
	rabbitCloseReceiver chan *amqp.Error
)

func rabbitConnector(url string) (conn *amqp.Connection, err error) {
	var rabbitErr *amqp.Error

	for {
		rabbitErr = <-rabbitCloseReceiver
		if rabbitErr != nil {
			log.Printf("Connection closed, try to reconnect to %s\n", url)

			rabbitConn, err = connectToRabbitMQ(url)
			if err != nil {
				failOnError(err, "no connection to RabbitMQ possible")
			}
			rabbitCloseReceiver = make(chan *amqp.Error)
			rabbitConn.NotifyClose(rabbitCloseReceiver)
		}
	}

}

func connectToRabbitMQ(url string) (conn *amqp.Connection, err error) {
	t0 := time.Now()

	for {
		conn, err = amqp.Dial(msgUrl)
		if err == nil {
			break
		}

		time.Sleep(1000 * time.Millisecond)
		log.Printf("Trying to connect to RabbitMQ at %s\n", url)
		if time.Now().Sub(t0) > 30*time.Second {
			log.Printf("Timeout for connecting reached\n ..... last error: %v\n", err.Error())
			break
		}
	}
	return conn, err
}

func main() {

	conn, err := connectToRabbitMQ(msgUrl)

	if err != nil {
		log.Printf("err: %s", err.Error())
		panic(fmt.Sprintf("........ err: %s", err.Error()))
	}
	log.Printf("Connected: %v", conn != nil)

	receiver := make(chan *amqp.Error)
	conn.NotifyClose(receiver)
	//	go rabbitConnector(msgUrl)
	for r := range receiver {
		fmt.Printf("Connection CLosed received: %v\n", r)
		conn, err := connectToRabbitMQ(msgUrl)
		if err != nil {
			log.Printf("err: %s", err.Error())
			panic(fmt.Sprintf("........ err: %s", err.Error()))
		}
		log.Printf("Connected: %v", conn != nil)
	}

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
