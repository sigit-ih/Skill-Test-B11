package main

// Import package
import (
	"fmt"
	"strconv"
)

// Function to calculate difference between sum of inputted number and nearest fibonacci number
func fibonacci(arr []int) int {
	// Calculate sum of inputted number
	var sumArr, output int
	for _, x := range arr {
		sumArr += x
	}
	// Create initial variabel to search fibonacci number
	var sum, temp, prev, next, i int
	a := 0
	b := 1
	sum = a + b
	temp = b
	b = sum
	a = temp
	// Loop fibonacci until sum of input are between previous and next number of fibonacci
	for true {
		sum = a + b
		prev = b
		next = sum
		// Check if sumArr number between next and prev fibonacci number
		if prev < sumArr && sumArr < next {
			// Calculate difference
			m := sumArr - prev
			n := next - sumArr
			/* See which fibonacci number between previous and next have smaller difference from sum of
			inputted number, then break loop */
			if m > n {
				output = n
				break
			} else if n > m {
				output = m
				break
			}
			// If not, resume fibonacci
		} else {
			temp = b
			b = sum
			a = temp
		}
		i++
	}
	// Return calculation result
	return output
}

// Main function
func main() {
	var inpCount, result int
	// Input how many number inputted
	fmt.Printf("Number of Input : ")
	fmt.Scanln(&inpCount)
	arr := []int{}
	var insArr int
	// Loop input number
	for i := 0; i < inpCount; i++ {
		j := strconv.Itoa(i + 1)
		fmt.Printf("%s	: ", j)
		fmt.Scanln(&insArr)
		arr = append(arr, insArr)
	}
	// Search difference of nearest fibonacci using fibonacci function
	result = fibonacci(arr)
	// Print result
	fmt.Println("Output : ", result)
}
