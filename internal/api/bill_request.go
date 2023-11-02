package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListBills() (Bill, error) {
	URL := url + apiKey + "&op=getMasterList&state=FL"
	res, err := http.Get(URL)
	if err != nil {
		fmt.Printf("Failed to send get request: %v\n", err)
	}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return Bill{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Bill{}, err
	}
	defer res.Body.Close()

	if resp.StatusCode > 399 {
		return Bill{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return Bill{}, err
	}
	billList := Bill{}
	err = json.Unmarshal(dat, &billList)
	if err != nil {
		fmt.Printf("Failed to read response body: %v\n", err)
	}
	return billList, nil
}
