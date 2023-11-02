package main

import "github.com/bobgtm/laws-cli-api/internal/api"

type config struct {
	lawsClient api.Client
}

func main() {
	cfg := config{
		lawsClient: api.NewClient(),
	}
	SartRepl(&cfg)
}
