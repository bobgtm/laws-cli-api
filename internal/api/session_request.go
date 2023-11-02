package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListSessions() (SessionList, error) {
	URL := url + apiKey + "&op=getSessionList&state=FL"
	res, err := http.Get(URL)
	if err != nil {
		fmt.Printf("Failed to send get request: %v\n", err)
	}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return SessionList{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return SessionList{}, err
	}
	defer res.Body.Close()

	if resp.StatusCode > 399 {
		return SessionList{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return SessionList{}, err
	}
	sessionList := SessionList{}
	err = json.Unmarshal(dat, &sessionList)
	if err != nil {
		fmt.Printf("Failed to read response body: %v\n", err)
	}
	return sessionList, nil
}
