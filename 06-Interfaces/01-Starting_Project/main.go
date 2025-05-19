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

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver
	Display() string
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

	outputData(todo)

	outputData(userNote)
}

// ______________________________________________________________________
// ! WARN: "interface{}" is the same as "any" in JS
func printSomething(value interface{}) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer: ", intVal+1)
		return
	}

	// switch value.(type) {
	// case string:
	// 	fmt.Println("This is a string: ", value)
	// case int:
	// 	fmt.Println("This is an int: ", value)
	// default:
	// 	fmt.Println("Unknown type")
	// }
	// fmt.Println(value)
}

// ______________________________________________________________________
func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

// ______________________________________________________________________
func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}
	fmt.Println("Note saved successfully.")

	return nil
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
