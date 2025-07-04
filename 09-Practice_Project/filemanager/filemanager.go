package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

// -----------------------------------------------------------------
// FileManager - * INFO: Struct to manage file operations
func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("failed to open file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err() // Check for errors during scanning
	if err != nil {     // Handle error
		file.Close() // Ensure the file is closed after reading
		return nil, errors.New("failed to read lines in file.")
	}
	file.Close() // Close the file after reading

	return lines, nil
}

// -----------------------------------------------------------------
func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		errors.New("failed to create file")
	}

	// * INFO: Sleep for 3 seconds to simulate a delay before writing
	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close() // Ensure the file is closed after writing
		errors.New("failed to encode data to JSON")
	}

	defer file.Close() // Ensure the file is closed after writing
	return nil
}

// * INFO: CONSTRUCTOR -----------------------------------------------------------------
func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
