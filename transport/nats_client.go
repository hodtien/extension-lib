package transport

import (
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

// ClientConstructor - ClientConstructor
func ClientConstructor(url string) *nats.Conn {
	connectionID := uuid.New().String()
	opts := []nats.Option{nats.Name("NATS_ExtensionLib_Requestor_" + connectionID)}

	nc, err := nats.Connect(url, opts...)
	if err != nil {
		log.Println("CONNECT TO NATS SERVER FAILED:", err)
		os.Exit(-1)
	}

	return nc
}

// ClientDestructor - ClientDestructor
func ClientDestructor(nc *nats.Conn) {
	nc.Close()
}
