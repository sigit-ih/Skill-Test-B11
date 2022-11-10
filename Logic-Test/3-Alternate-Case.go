package main

// Import package
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Function to alternate case of each inputted letter
func alternateCase(input string) string {
	arr := []string{}
	// Loop each letter of inputted word and alternate case
	for _, r := range input {
		// Check if letter is uppercase or lowercase using letter rune
		if unicode.IsUpper(r) {
			// Change case to lowercase
			r = unicode.ToLower(r)
		} else if unicode.IsLower(r) {
			// Change case to uppercase
			r = unicode.ToUpper(r)
		}
		// Insert each letter into arr
		arr = append(arr, string(r))
	}
	// Merge each letter in arr into one word in output variable
	output := strings.Join(arr, "")
	// Return altered word
	return output
}

// Main function
func main() {
	// Input word
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numInp := scanner.Text()
	// Alternate case of inputted word using alternateCase function
	x := alternateCase(numInp)
	// Print result
	fmt.Println(x)
}
