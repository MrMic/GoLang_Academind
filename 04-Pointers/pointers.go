package main

import "fmt"

func main() {
	age := 32 // Just a regular variable
	agePointer := &age

	fmt.Println("Age: ", age)
	fmt.Println("Age Pointer: ", agePointer)
	fmt.Println("Value of Age Pointer: ", *agePointer)

	editAgeToAdultYears(agePointer)
	fmt.Println("Adult Years: ", age)
}

// ______________________________________________________________________
func editAgeToAdultYears(age *int) {
	// return *age - 18
	*age -= 18
}
