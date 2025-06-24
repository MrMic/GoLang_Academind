package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

// -----------------------------------------------------------------
// FileManager - * INFO: Struct to manage file operations
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("failed to open file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil { // Handle error
		file.Close() // Ensure the file is closed after reading
		return nil, errors.New("failed to read lines in file.")
	}
	file.Close() // Close the file after reading

	return lines, nil
}

// -----------------------------------------------------------------
func WriteJSON(path string, data any) error {
	file, err := os.Create(path)
	if err != nil {
		errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close() // Ensure the file is closed after writing
		errors.New("failed to encode data to JSON")
	}

	defer file.Close() // Ensure the file is closed after writing
	return nil
}
