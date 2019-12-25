package event

import (
	"log"
	"natsp/config"

	nats "github.com/nats-io/nats.go"
	nat "github.com/nats-io/nats.go/encoders/protobuf"
)

//NewLocalConnection
func NewLocalConnection() *nats.EncodedConn {

	//connect to nats
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	ec, err := nats.NewEncodedConn(nc, nats.PROTOBUF_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	return ec
}

//NewConnection
func NewConnection(configURL string) *nats.EncodedConn {

    //connect to nats
    nc, err := nats.Connect(configURL)
    if err != nil {
        log.Fatal(err)
    }
    ec, err := nats.NewEncodedConn(nc, nat.PROTOBUF_ENCODER)
    if err != nil {
        log.Fatal(err)
    }
    return ec
}