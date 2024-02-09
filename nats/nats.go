package nats

import (
	"caching-proxies-terminal/config"
	"caching-proxies-terminal/support/connection"

	"github.com/nats-io/nats.go"
)

func GetForJWT(jwtToken string) (*nats.Conn, error) {
	ns, _, err := connection.Establish(*config.FlagServer, jwtToken)
	if err != nil {
		return nil, err
	}
	return ns, nil
}
