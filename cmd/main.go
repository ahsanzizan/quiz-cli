package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	data, err := os.Open("../data/questions.json")
	if err != nil {
		panic("Failed to open the questions.json file. Make sure it exist in the /data/questions.json first.")
	}

	fmt.Println("Successfully opened the questions.json file.")
	defer data.Close()

	byteValue, err := io.ReadAll(data)
	if err != nil {
		panic("Failed to read the questions.json file.")
	}

	var questions []Question
	json.Unmarshal([]byte(byteValue), &questions)
}
