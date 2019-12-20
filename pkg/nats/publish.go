package event

import (
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
)

//NewVerifyPublisher new verify publisher
func NewVerifyPublisher(nc *nats.EncodedConn) model.RelationEvent {
	return &verifyPublisher{nc: nc}
}

type verifyPublisher struct {
	nc *nats.EncodedConn
}

func (p verifyPublisher) PublishEvent(isDelete bool, relation string, from, to int64) error {

	message := &relationv1.Relation{
		From:     from,
		To:       to,
		Relation: utills.DecodeRelationTypeFromString(relation),
	}

	bs, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	err = p.nc.Publish("relation."+relation, bs)
	defer p.nc.Close()
	if err != nil {
		return err
	}

	return nil
}
