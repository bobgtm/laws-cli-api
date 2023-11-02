package api

import (
	"net/http"
	"time"
)

const (
	apiKey = "fd72601a74ea1325d85b5a89bc11bcae"
	url    = "https://api.legiscan.com/?key="
)

type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
