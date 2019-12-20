package nats

import (
	"log"
	"natsp/config"

	"github.com/nats-io/nats.go"
	nats "github.com/nats-io/nats.go/encoders/protobuf"
)

//NewLocalConnection
func NewLocalConnection(config config.Configuration) *nats.EncodedConn {

	//connect to nats
	nc, err := nats.Connect(config.Nats.URL)
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
    ec, err := nats.NewEncodedConn(nc, nats.PROTOBUF_ENCODER)
    if err != nil {
        log.Fatal(err)
    }
    return ec
}