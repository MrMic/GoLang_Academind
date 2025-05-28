package main

import "fmt"

func main() {
	// ! WARN: Slice & nothing is pre-allocated
	// ! WARN: like with `make()`
	// userNames := []string{}
	userNames := make([]string, 2, 5)

	userNames[0] = "Alice"
	userNames[1] = "Bob"
	fmt.Println("🪚 userNames:", userNames)

	userNames = append(userNames, "John")
	userNames = append(userNames, "Jane")
	fmt.Println("🪚 userNames:", userNames)
}
