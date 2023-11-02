package api

import (
	"net/http"
	"time"
)

const (
	url    = "https://api.legiscan.com/?key="
	apiKey = "fd72601a74ea1325d85b5a89bc11bcae"
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
