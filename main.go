package main

import (
	"flag"
	"fmt"
	"time"
)

var producerNum int
var consumerNum int

func init() {
	flag.IntVar(&consumerNum, "consume", 1, "number of consumer")
	flag.IntVar(&producerNum, "produce", 1, "number of producer")
	flag.Parse()
}

func main() {
	messageChan := make(chan string, 20)
	for i := 1; i <= producerNum; i++ {
		producerID := fmt.Sprintf("producer-%d", i)
		p := NewStrProducer(producerID, messageChan)
		go launchProducer(p)
	}

	for i := 1; i <= consumerNum; i++ {
		c := NewStrConsumer(messageChan)
		go launchConsumer(c)
	}

	time.Sleep(10 * time.Second)
}

func launchProducer(p producer) {
	for {
		p.produce()
		time.Sleep(1 * time.Second)
	}
}

func launchConsumer(c consumer) {
	for {
		c.consume()
	}
}
