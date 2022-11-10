package main

// Import package
import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Function to sort number from lowest to highest
func maxRedigit(input string) string {
	// Convert each inputted number into string and insert into arr
	arr := []string{}
	for i := 0; i < len(input); i++ {
		n := string(input[i])
		arr = append(arr, n)
	}
	// Sort number
	sort.Sort(sort.StringSlice(arr))
	// Join each number in arr into output
	output := strings.Join(arr, "")
	// Return sorted number
	return output
}

// Function to check if string is number
func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// Main function
func main() {
	var numInp string
	fmt.Scanln(&numInp)
	// Check if number count is three and inputted number is number
	if len(numInp) == 3 && isNumeric(numInp) == true {
		// If true, sort number with maxRedigit function and print result
		x := maxRedigit(numInp)
		fmt.Println(x)
		// If not, print null
	} else {
		fmt.Println("null")
	}
}
