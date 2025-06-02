package main

import "fmt"

// ______________________________________________________________________
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println("🪚 MAP:", m)
}

// * NOTE: ╾╼ MAIN ╾──────────────────────────────────────────────────────────╼
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

	coursesRatings := make(floatMap, 3)
	coursesRatings["Go"] = 4.9
	coursesRatings["Python"] = 4.8
	coursesRatings["JavaScript"] = 4.7

	// fmt.Println("🪚 coursesRatings:", coursesRatings)
	coursesRatings.output()

	println("-------------------------------")

	// * INFO: for Arrays
	for index, value := range userNames {
		fmt.Println("🪚 index:", index, " => ", "value:", value)
	}

	println("-------------------------------")

	// * INFO: for Maps
	for key, value := range coursesRatings {
		fmt.Println("🪚 key:", key, " => ", "value:", value)
	}
}
