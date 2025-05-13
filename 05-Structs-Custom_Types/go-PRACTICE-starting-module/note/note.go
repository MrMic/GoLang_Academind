package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

// * INFO: ══ CONSTRUCTOR ═════════════════════════════════════════════════════
func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}

// ______________________________________________________________________
func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	jsonContent, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonContent, 0o644)
}

// ______________________________________________________________________
func (note Note) Display() string {
	fmt.Printf("\nTitle: %v\n", note.title)
	fmt.Printf("Content: %v\n", note.content)
	fmt.Printf("Created At: %s\n", note.createdAt)
	return ""
}
