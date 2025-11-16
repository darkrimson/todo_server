package model

type Task struct {
	ID      string `json:"id,primary_key"`
	Title   string `json:"title"`
	Date    string `json:"date"`
	Comment string `json:"comment,omitempty"`
	Repeat  string `json:"repeat,omitempty"`
}
