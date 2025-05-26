package main

import "fmt"

// ______________________________________________________________________
type Product struct {
	id    string
	title string
	price float64
}

// * INFO:  ╾╼ MAIN ╾──────────────────────────────────────────────────────────╼
func main() {
	// 1.
	hobbies := [3]string{"Sport", "Cooking", "Reading"}
	fmt.Println(hobbies)

	// 2.
	fmt.Println(hobbies[1:3])

	// 3.
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	// 4.
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	// 5.
	courseGoals := []string{"Learn Go!", "Learn Python!"}
	fmt.Println(courseGoals)

	// 6.
	courseGoals[1] = "Learn C++!"
	courseGoals = append(courseGoals, "Learn C++")
	fmt.Println(courseGoals)

	// 7.
	products := []Product{
		{"1", "Laptop", 1000},
		{"2", "Mouse", 50},
	}
	fmt.Println(products)

	products = append(products, Product{"3", "Keyboard", 100})
	fmt.Println(products)
}
