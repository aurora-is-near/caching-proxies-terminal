package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"io"
	"net/http"
	"time"

	"caching-proxies-terminal/config"
	"caching-proxies-terminal/nats"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	nats2 "github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()
	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	// Echo instance
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// Middleware
	e.Use(middleware.Recover())

	// Routes
	e.POST("/process", process)

	// Start server
	e.Logger.Fatal(e.Start("0.0.0.0:1323"))
}

func process(c echo.Context) error {
	authorizationToken := c.QueryParam("token")
	if authorizationToken == "" {
		return c.String(403, "Forbidden for the provided authorization token")
	}

	bearer := authorizationToken
	jwt, err := Verify(bearer)
	if err != nil {
		logrus.Error(err)
		return c.String(403, "Forbidden for the provided authorization token")
	}

	logrus.Info("Active jwt for the request is: ", jwt)

	receivedAt := time.Now()
	previousHashID := c.QueryParam("previous_hash_id")
	shardID := c.QueryParam("shard_id")
	blob, _ := io.ReadAll(c.Request().Body)

	ns, err := nats.GetForJWT(jwt)
	if err != nil {
		logrus.Error(err)
		return c.String(403, "Forbidden: NATS forbid the connection")
	}
	defer ns.Drain()
	defer ns.Close()

	blockHasher := sha256.New()
	blockHasher.Write(blob)
	blockHash := blockHasher.Sum(nil)
	blockHashAsHex := hex.EncodeToString(blockHash)

	var msgID string
	if *config.FlagUseBlobHash {
		msgID = blockHashAsHex
	} else {
		msgID = previousHashID + ":" + shardID
	}

	err = ns.PublishMsg(&nats2.Msg{
		Subject: *config.FlagShardPrefix + "." + shardID,
		Header: map[string][]string{
			nats2.MsgIdHdr:       {msgID},
			"X-Block-Hash":       {blockHashAsHex},
			"X-Previous-Hash-Id": {previousHashID},
			"X-Shard-Id":         {shardID},
			"X-Received-At":      {time.Now().String()},
		},
		Data: blob,
	})
	if err != nil {
		logrus.Error(err)
		return err
	}

	logrus.WithFields(map[string]interface{}{
		"previous_hash_id": previousHashID,
		"shard_id":         shardID,
		"block_hash":       blockHashAsHex,
		"msg_id":           msgID,
		"received_at":      receivedAt.String(),
		"subject":          *config.FlagShardPrefix + "." + shardID,
	}).Infof(
		"Published a message with %s header of %s. Block hash is %s. Time when request came: %s. Time spent working: %s.",
		nats2.MsgIdHdr, msgID, blockHashAsHex, receivedAt.String(), time.Since(receivedAt).String(),
	)

	return c.String(http.StatusOK, "OK "+msgID)
}
