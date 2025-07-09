package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Question struct {
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
}

type Answer struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	IsCorrect bool   `json:"is_correct"`
}

func parseQuestionsConfig() ([]Question, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get the current working directory")
	}

	dataPath := filepath.Join(wd, "data", "questions.json")

	data, err := os.Open(dataPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open %v. Make sure it exist first", dataPath)
	}

	fmt.Printf("Successfully opened the %v.\n", dataPath)

	defer data.Close()

	byteValue, err := io.ReadAll(data)
	if err != nil {
		return nil, fmt.Errorf("failed to read %v: %w", dataPath, err)
	}

	var questions []Question
	err = json.Unmarshal([]byte(byteValue), &questions)
	if err != nil {
		fmt.Printf("")
		return nil, err
	}

	fmt.Printf("Successfully parsed %v\n", dataPath)

	return questions, nil
}

func main() {
	questions, err := parseQuestionsConfig()
	if err != nil {
		panic("Failed to open the questions config file.")
	}

	for index, question := range questions {
		fmt.Printf("%v. %s", index, question.Question)
		// Print the answers
		// Prompt the user to answer with the Answer Key
		// If the key does not exist, then print "Your answer is invalid, please answer properly this time." -> Re-prompt (While true loop)
		// If the key exist, then check if the key is a correct answer
		// If it is a correct answer, then modify a Score value somewhere
		// Continue to next question
	}
}
