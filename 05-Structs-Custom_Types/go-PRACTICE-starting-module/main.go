package main

import (
	"errors"
	"fmt"
)

func main() {
	title, content, err := getNoteData()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// ______________________________________________________________________
func getNoteData() (string, string, error) {
	title, err := getUserInput("Note Title: ")
	if err != nil {
		return "", "", err
	}

	content, err := getUserInput("Note Content: ")
	if err != nil {
		return "", "", err
	}

	return title, content, nil
}

// ______________________________________________________________________
func getUserInput(prompt string) (string, error) {
	fmt.Println(prompt)

	var input string
	fmt.Scanln(&input)

	if input == "" {
		return "", errors.New("invalid input")
	}
	return input, nil
}
