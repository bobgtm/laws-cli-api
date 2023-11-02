package api

type SessionList struct {
	Status   string `json:"status"`
	Sessions []struct {
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
		DatasetHash  string `json:"dataset_hash"`
		SessionHash  string `json:"session_hash"`
		Name         string `json:"name"`
	} `json:"sessions"`
}
