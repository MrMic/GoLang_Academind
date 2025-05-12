package main

import (
	"fmt"

	"michaelchlon.fr/note/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
}

// ______________________________________________________________________
func getNoteData() (string, string) {
	title := getUserInput("Note Title: ")

	content := getUserInput("Note Content: ")

	return title, content
}

// ______________________________________________________________________
func getUserInput(prompt string) string {
	fmt.Print(prompt)

	var input string
	fmt.Scanln(&input)

	return input
}
