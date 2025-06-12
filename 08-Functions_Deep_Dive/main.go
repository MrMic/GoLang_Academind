package main

import (
	"fmt"
)

func main() {
	numnbers := []int{1, 2, 3, 4, 5}
	sum := sumup(1, 10, 15)
	anotherSum := sumup(1, numnbers...)

	// fmt.Println(sum)
	fmt.Println("The sum is:", sum)
	fmt.Println("The another sum is:", anotherSum)

}

// * INFO: ------------------------------
func sumup(startingValue int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}
	return sum
}
