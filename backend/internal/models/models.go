package models

type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"`
	Timestamp string `json:"timestamp"`
}

type Room struct {
	ID string `json:"id"`
	Users []string `json:"users"`
}