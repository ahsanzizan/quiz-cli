package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
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

func findAnswerByKey(options []Answer, key string) (Answer, bool) {
	for _, a := range options {
		if strings.EqualFold(a.Key, key) {
			return a, true
		}
	}

	return Answer{}, false
}

func main() {
	questions, err := parseQuestionsConfig()
	if err != nil {
		panic("Failed to open the questions config file.")
	}

	fmt.Println("Starting a quiz session...")
	fmt.Println()

	correctAnswers := 0

	for index, question := range questions {
		fmt.Printf("%d.  %s\n", index+1, question.Question)

		// Print the answers option
		for _, answer := range question.Answers {
			fmt.Printf("    %s. %s\n", answer.Key, answer.Value)
		}

		var userAnswer string

		for {
			// Prompt the user to answer with the Answer Key
			fmt.Print("Enter your answer: ")
			fmt.Scan(&userAnswer)

			// If the key does not exist, then print "Your answer is invalid, please answer properly this time." -> Re-prompt (While true loop)
			answer, found := findAnswerByKey(question.Answers, userAnswer)
			if !found {
				// Reloop current question
				fmt.Printf("Your answer (%s) is not recognized as a valid answer to the question.", userAnswer)
				continue
			}

			// If the key exist, then check if the key is a correct answer
			if answer.IsCorrect {
				// If it is a correct answer, then modify a Score value somewhere
				correctAnswers++
			}

			break
		}

		if index != len(questions)-1 {
			fmt.Println("Proceeding to the next question.")
		}

		fmt.Println()
	}

	fmt.Printf("\nCongratulations! You've just finished the quiz.\nYour total score: %d/%d", correctAnswers, len(questions))
}
