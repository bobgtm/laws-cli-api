package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var key = "fd72601a74ea1325d85b5a89bc11bcae"

type Bill struct {
	Status     string `json:"status"`
	Masterlist struct {
		Num0 struct {
			BillID         int    `json:"bill_id"`
			Number         string `json:"number"`
			ChangeHash     string `json:"change_hash"`
			URL            string `json:"url"`
			StatusDate     string `json:"status_date"`
			Status         int    `json:"status"`
			LastActionDate string `json:"last_action_date"`
			LastAction     string `json:"last_action"`
			Title          string `json:"title"`
			Description    string `json:"description"`
		} `json:"0"`
		Num1 struct {
			BillID         int    `json:"bill_id"`
			Number         string `json:"number"`
			ChangeHash     string `json:"change_hash"`
			URL            string `json:"url"`
			StatusDate     string `json:"status_date"`
			Status         int    `json:"status"`
			LastActionDate string `json:"last_action_date"`
			LastAction     string `json:"last_action"`
			Title          string `json:"title"`
			Description    string `json:"description"`
		} `json:"1"`
		Session struct {
			SessionID    int    `json:"session_id"`
			StateID      int    `json:"state_id"`
			YearStart    int    `json:"year_start"`
			YearEnd      int    `json:"year_end"`
			Prefile      int    `json:"prefile"`
			SineDie      int    `json:"sine_die"`
			Prior        int    `json:"prior"`
			Special      int    `json:"special"`
			SessionTag   string `json:"session_tag"`
			SessionTitle string `json:"session_title"`
			SessionName  string `json:"session_name"`
		} `json:"session"`
	} `json:"masterlist"`
}

var url = "http://api.legiscan.com/?key=" + key

// Per the LegiSCan API User Manual
// Most use cases for the api can be satsified with daily updates
// although recommended frequency varies for different operations
func showBySession() error {
	// &op=getMasterList&id=(session ID to be retrieved from Show Session)

	apiURL := url + "&op=getMasterList&id=1987"
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Printf("Failed to send get request: %v\n", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Failed to read response body: %v\n", err)
	}

	var data Bill
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Fields:", data)
	return nil
}
