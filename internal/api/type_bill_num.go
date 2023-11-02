package api

type Bill struct {
	Status     string         `json:"status"`
	Masterlist map[string]Num `json:"masterlist"`
}

type Num struct {
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
}
