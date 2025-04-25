package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// ______________________________________________________________________
func GetFloatFromFile(fileName string) (float64, error) {
	fmt.Println("Reading balance from file...")
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("failed to find file")
	}

	value, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 0, errors.New("failed to parse stored value")
	}

	return value, nil
}

// ______________________________________________________________________
func WriteFloatToFile(value float64, fileName string) {
	fmt.Println("Writing balance to file:", value)
	balanceText := fmt.Sprintf("%.2f", value)
	os.WriteFile(fileName, []byte(balanceText), 0o644)
}
