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
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// * INFO: ══ CONSTRUCTOR ═════════════════════════════════════════════════════
func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

// ______________________________________________________________________
func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	jsonContent, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonContent, 0o644)
}

// ______________________________________________________________________
func (note Note) Display() string {
	fmt.Printf("\nTitle: %v\n", note.Title)
	fmt.Printf("Content: %v\n", note.Content)
	fmt.Printf("Created At: %s\n", note.CreatedAt)
	return ""
}
