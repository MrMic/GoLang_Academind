package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"michaelchlon.fr/note/note"
	"michaelchlon.fr/note/todo"
)

// ______________________________________________________________________
type saver interface {
	Save() error
}

// * INFO: ══ MAIN ════════════════════════════════════════════════════════════
func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = todo.Save()
	if err != nil {
		fmt.Println("Saving the todo failed.")
		return
	}
	fmt.Println("todo saved successfully.")

	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}
	fmt.Println("Note saved successfully.")
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

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
