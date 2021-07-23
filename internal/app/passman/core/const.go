package core

type Item struct {
	Id          string   `json:"id"`
	Type        string   `json:"type"`
	Service     string   `json:"service"`
	Description string   `json:"description"`
	Content     []string `json:"content"`
}

var DataDir = ""
