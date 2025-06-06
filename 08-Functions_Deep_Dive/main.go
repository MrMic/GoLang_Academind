package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2)
	triple := createTransformer(3)

	transformed := transformNumbers(&numbers, func(numbers int) int {
		return numbers * 2
	})

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println("🪚 doubled:", doubled)
	fmt.Println("🪚 tripled:", tripled)
	fmt.Println("🪚 transformed:", transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	//* INFO: __ Return a CLOSURE!!! _____________________________________________
	return func(number int) int {
		return number * factor
	}
}
