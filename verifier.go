package main

import (
	"caching-proxies-terminal/config"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Verify asks the submissions verifier service to verify the token and return the JWT to connect to NATS
func Verify(bearerToken string) (string, error) {
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

	jwt := answer["jwt"].(string)

	return jwt, nil
}
