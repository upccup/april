package main

import (
	"fmt"
	"testing"
)

func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer()
	messageChan := make(chan string, 20)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		producerID := fmt.Sprintf("producer-%d", i)
		p := NewStrProducer(producerID, messageChan)
		p.produce()

		c := NewStrConsumer(messageChan)
		c.consume()
	}
}
