package main

import "fmt"

func main() {
	sum := sumup(1, 10, 15)

	fmt.Println(sum)

}

// * INFO: ------------------------------
func sumup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}
	return sum
}
