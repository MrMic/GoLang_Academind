package main

import "fmt"

func main() {
	age := 32 // Just a regular variable
	agePointer := &age

	fmt.Println("Age: ", age)
	fmt.Println("Age Pointer: ", agePointer)
	fmt.Println("Value of Age Pointer: ", *agePointer)

	adultYears := getAdultYears(age)
	fmt.Println(adultYears)
}

// ______________________________________________________________________
func getAdultYears(age int) int {
	return age - 18
}
