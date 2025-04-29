package main

import "fmt"

func main() {
	age := 32 // Just a regular variable
	agePointer := &age

	fmt.Println("Age: ", age)
	fmt.Println("Age Pointer: ", agePointer)
	fmt.Println("Value of Age Pointer: ", *agePointer)

	adultYears := getAdultYears(agePointer)
	fmt.Println("Adult Years: ", adultYears)
}

// ______________________________________________________________________
func getAdultYears(age *int) int {
	if age == nil {
		return 0 // Gérer le cas où le pointeur est nil
	}
	if *age < 18 {
		return 0 // Retourner 0 si l'âge est inférieur à 18
	}
	return *age - 18
}
