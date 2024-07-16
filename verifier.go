package main

import (
	"caching-proxies-terminal/config"
	"caching-proxies-terminal/storage"
	"encoding/json"
	"errors"
	"github.com/nats-io/jwt/v2"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Verify asks the submissions verifier service to verify the token and return the JWT to connect to NATS
func Verify(storage storage.JwtStorage, bearerToken string) (string, error) {
	// Try to get the JWT from the storage
	token, err := storage.Get(bearerToken)
	if err == nil {
		// Check for its expiration time
		tokenRows := strings.Split(token, "\n")
		if len(tokenRows) < 2 {
			logrus.Warn("Cached JWT has wrong format: ", token)
			return "", errors.New("cached JWT has wrong format")
		}

		tokenString := tokenRows[1]

		claims, err := jwt.DecodeUserClaims(tokenString)
		if err != nil {
			logrus.Warn("Could not decode claims from cached JWT: ", token)
		} else if claims.Expires > time.Now().Add(10*time.Second).Unix() {
			return token, nil
		} else {
			logrus.Warn("Cached JWT has expired: ", token)
		}
	} else {
		logrus.Warn("Could not get JWT from cache: ", err)
	}

	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}

	values := url.Values{}
	values.Set("token", bearerToken)

	u, err := url.Parse(*config.FlagSubmissionsVerifierHost + "?" + values.Encode())
	if err != nil {
		return "", err
	}

	req := http.Request{
		Method: "POST",
		URL:    u,
	}

	do, err := httpClient.Do(&req)
	if err != nil {
		return "", err
	}
	defer do.Body.Close()

	body, err := io.ReadAll(do.Body)
	if err != nil {
		return "", err
	}

	answer := map[string]interface{}{}
	err = json.Unmarshal(body, &answer)

	status, ok := answer["status"].(string)
	if status != "ok" || !ok {
		return "", errors.New("status from submissions verifier is not ok")
	}

	token = answer["jwt"].(string)
	err = storage.Store(bearerToken, token)
	if err != nil {
		logrus.Warning("Could not store JWT in cache: ", err)
		return "", err
	}

	return token, nil
}
