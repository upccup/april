package main

import "testing"

func TestProducer(t *testing.T) {
	messageChan := make(chan string, 1)
	p := NewStrProducer("a", messageChan)
	p.produce()
	m := <-messageChan
	t.Log(m)
}
