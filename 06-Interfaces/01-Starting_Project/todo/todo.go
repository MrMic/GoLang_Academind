package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

// * INFO: ══ CONSTRUCTOR ═════════════════════════════════════════════════════
func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("invalid input")
	}

	return Todo{
		Text: content,
	}, nil
}

// ______________________________________________________________________
func (todo Todo) Save() error {
	fileName := "todo.json"

	jsonContent, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonContent, 0o644)
}

// ______________________________________________________________________
func (todo Todo) Display() string {
	fmt.Printf("Content: %v\n", todo.Text)
	return ""
}
