package sms

import "log"

type Message struct{}

type Messenger interface {
	SendMessage(msg string) error
}

func NewMsg() *Message {
	return &Message{}
}

func (m *Message) SendMessage(msg string) error {
	log.Println("*************************************")
	log.Printf("Message: %s\n", msg)
	log.Println("*************************************")
	return nil
}
