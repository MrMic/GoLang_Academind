package functionsarevalues

import "fmt"

// * INFO: ╾╼ TYPE ╾──────────────────────────────────────────────────────────╼
type transformFn func(int) int

// type anotherFn func(int, []string, map[string][]int) ([]int, string)

// * INFO: ╾╼ MAIN ╾──────────────────────────────────────────────────────────╼
func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformerFn1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformerFn2)

	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)
}

// ______________________________________________________________________
func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// ______________________________________________________________________
func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

// ______________________________________________________________________
func double(number int) int {
	return number * 2
}

// ______________________________________________________________________
func triple(number int) int {
	return number * 3
}
