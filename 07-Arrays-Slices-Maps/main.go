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

	coursesRatings := make(map[string]float64, 3)
	coursesRatings["Go"] = 4.9
	coursesRatings["Python"] = 4.8
	coursesRatings["JavaScript"] = 4.7
	fmt.Println("🪚 coursesRatings:", coursesRatings)
}
