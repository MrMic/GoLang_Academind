package filemanager

import (
	"bufio"
	"errors"
	"os"
)

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
