package iomanager

type IOMAnager interface {
	ReadLines() ([]string, error) // Read lines from a file
	WriteResult(data any) error   // Write the result to a file
}
