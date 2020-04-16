package connect

import (
	"fmt"
	nats "github.com/nats-io/nats.go"
	"natsp/pkg/model"
	"reflect"
	"strings"
)

func Publish (nc *nats.EncodedConn, sbj string, message model.Model) {
	fmt.Print("Enter Subject of message: ")
	_, err := fmt.Scanf("%d\n", sbj)
	if err != nil {
		fmt.Println("ERR IN READ SUBJECT")
	}

	e := reflect.ValueOf(&message).Elem()
	for i := 0; i < e.NumField(); i++ {
		fmt.Sprintf("Enter value of %s (type=%s)", e.Type().Field(i).Name, e.Type().Field(i).Type)
		t := e.Type().Field(i).Type
		v := make(e.Type().Field(i).Type,0)

		e.Field(i).Set()

		reflect.ValueOf(&n).Elem().FieldByName("N").SetInt(7)

	}

	err = nc.Publish(strings.TrimSpace(sbj), message)
	defer nc.Close()
	if err != nil {
		fmt.Println("ERR Publish: ", err)
		return
	}
}
