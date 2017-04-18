package main

import "testing"

func TestConsumer(t *testing.T) {
	messageChan := make(chan string, 20)
	messageChan <- "a"
	messageChan <- "b"
	messageChan <- "c"
	messageChan <- "d"

	c := NewStrConsumer(messageChan)

	c.consume()

	c.consume()

	c.consume()

	c.consume()
}
