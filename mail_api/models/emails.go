package models

type SearchRequest struct {
	Term       string `json:"term"`
	Field      string `json:"field"`
	From       int    `json:"from"`
	MaxResults int    `json:"max_results"`
}
