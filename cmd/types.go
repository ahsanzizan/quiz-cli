package main

type Question struct {
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
}

type Answer struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	IsCorrect bool   `json:"is_correct"`
}
