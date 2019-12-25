package model 

type Event interface {
	PublishEvent(subject string, message struct) error 
}