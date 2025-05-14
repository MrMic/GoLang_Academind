package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving the todo failed.")
		return
	}

	fmt.Println("Note saved successfully.")
}

//

// ______________________________________________________________________
func getNoteData() (string, string) {
	title := getUserInput("Note Title: ")

	content := getUserInput("Note Content: ")

	return title, content
}

// ______________________________________________________________________
func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
