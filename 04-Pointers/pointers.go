package main

import "fmt"

func main() {
	age := 32 // Just a regular variable
	agePointer := &age

	fmt.Println("Age: ", age)

	adultYears := getAdultYears(age)
	fmt.Println(adultYears)
}

// ______________________________________________________________________
func getAdultYears(age int) int {
	return age - 18
}
