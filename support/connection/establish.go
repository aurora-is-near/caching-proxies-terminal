package connection

import (
	"fmt"
	"os"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func Establish(overrideServer string, jwtCreds string) (*nats.Conn, jetstream.JetStream, error) {
	tmpFile, err := os.CreateTemp("/tmp/", "nats")
	if err != nil {
		return nil, nil, err
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(jwtCreds)
	if err != nil {
		return nil, nil, err
	}

	ns, err := nats.Connect(overrideServer, nats.UserCredentials(tmpFile.Name()))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to nats: %w", err)
	}

	js, err := jetstream.New(ns)
	if err != nil {
		_ = ns.Drain()
		return nil, nil, fmt.Errorf("failed to connect to jetstream: %w", err)
	}

	return ns, js, err
}
