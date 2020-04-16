package connect

import (
	"bufio"
	"fmt"
	"log"
	nats "github.com/nats-io/nats.go"
	natsProtobuf "github.com/nats-io/nats.go/encoders/protobuf"
	"os"
	"strings"
)

func getConnectionUrl() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter your connection url (local=press enter) :")
	url,err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return ""
	}

	if strings.TrimSpace(url) == "\n" {
		url = "nats://127.0.0.1:4222"
		return url
	}

	return url
}

//NewConnection
func NewConnection() *nats.EncodedConn {
	url := getConnectionUrl()

	nc, err := nats.Connect(url)
	if err != nil {
		log.Println(err)
		return nil
	}

	ec, err := nats.NewEncodedConn(nc, natsProtobuf.PROTOBUF_ENCODER)
	if err != nil {
		log.Println(err)
		return nil
	}

	fmt.Println("Connected successfully: ", url)
	return ec
}