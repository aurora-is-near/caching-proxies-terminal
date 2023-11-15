package nats

import (
	"sync"

	"caching-proxies-terminal/config"
	"caching-proxies-terminal/support/connection"

	"github.com/nats-io/nats.go"
)

var natsConnection *nats.Conn
var once = &sync.Once{}

func Get() *nats.Conn {
	once.Do(func() {
		connect()
	})
	return natsConnection
}

func connect() {
	ns, _, err := connection.Establish(*config.FlagNatsContext, *config.FlagServer, *config.FlagCreds)
	if err != nil {
		panic(err)
	}

	natsConnection = ns
}
