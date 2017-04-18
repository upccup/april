package main

import "fmt"

type consumer interface {
	consume()
}

type StrConsumer struct {
	messageChan <-chan string
}

func NewStrConsumer(messageChan <-chan string) *StrConsumer {
	return &StrConsumer{messageChan}
}

func (c *StrConsumer) consume() {
	message := <-c.messageChan
	fmt.Println(message)
	return
}
