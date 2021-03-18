package transport

import (
	"github.com/nats-io/nats.go"
)

var Nc *nats.Conn
var Domain string

func Constructor(url, domain string) {
	Nc = ClientConstructor(url)
	Domain = domain
}
