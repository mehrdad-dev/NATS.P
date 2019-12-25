package event

import (
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"natsp/model"
)

//NewVerifyPublisher new verify publisher
func NewVerifyPublisher(nc *nats.EncodedConn) model.Event {
	return &verifyPublisher{nc: nc}
}

type verifyPublisher struct {
	nc *nats.EncodedConn
}

func (p verifyPublisher) PublishEvent(subject string, message struct) error {

	bs, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	err = p.nc.Publish(subject, bs)
	defer p.nc.Close()
	if err != nil {
		return err
	}

	return nil
}
