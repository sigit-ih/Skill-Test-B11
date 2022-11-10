package main

// Import package
import (
	"fmt"
	"strconv"
)

// Function to create array and sum of number with multiple of 3 and 5 below inputted number
func solution(input int) (int, string) {
	// Search number with multiplication of 3 and 5 below input number and insert into numbers array
	numbers := []int{}
	for i := 0; i < input; i++ {
		if i%3 == 0 || i%5 == 0 {
			numbers = append(numbers, i)
		}
	}
	// Sum numbers array and join number to create output string
	var numStr string
	output := 0
	for j := 0; j < len(numbers); j++ {
		if j > 0 {
			output += numbers[j]
			numStr += strconv.Itoa(numbers[j])
			if (j + 1) < len(numbers) {
				numStr += " + "
			}
		}
	}
	// Check if input number is zero
	if output == 0 {
		numStr = "0"
	}
	// Return sum number and output string
	return output, numStr
}

// Main function
func main() {
	// Insert number
	var numInp int
	fmt.Scanln(&numInp)
	/* Search number below inputted number with multiplication of 3 and 5 using solution function and print
	result */
	x, y := solution(numInp)
	fmt.Println(x, " = ", y)
}
