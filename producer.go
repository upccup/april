package main

type producer interface {
	produce()
}

type StrProducer struct {
	ID          string
	messageChan chan<- string
}

func NewStrProducer(ID string, messageChan chan<- string) *StrProducer {
	return &StrProducer{
		ID:          ID,
		messageChan: messageChan,
	}
}

func (p *StrProducer) produce() {
	p.messageChan <- p.ID
	return
}
