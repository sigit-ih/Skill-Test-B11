package main

// Import package
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Function to alternate case
func alternateCase(input string) string {
	// Create initial variable
	arrInit := []string{}
	arr := []string{}
	// Split inputted word by space and insert into arrInit
	arrInit = strings.Split(input, " ")
	// Loop every splitted word in arrInit
	for _, r := range arrInit {
		arrInner := []string{}
		checkCase := []bool{}
		// Loop every letter in splitted word
		for _, s := range r {
			/* Check if letter is uppercase or lowercase (true if upper, false if lower) and insert bool into
			checkCase. The order of checkCase is same as inputted word */
			if unicode.IsUpper(s) {
				checkCase = append(checkCase, true)
				arrInner = append(arrInner, string(s))
			} else if unicode.IsLower(s) {
				checkCase = append(checkCase, false)
				arrInner = append(arrInner, string(s))
			}
		}
		// Reverse word and change case in arrInner and insert into reverse slice
		reverse := []string{}
		for i := 0; i < len(arrInner); i++ {
			j := len(arrInner) - i - 1
			// Change case of reversed letter by checkCase order and insert into reverse slice
			if checkCase[i] == true {
				reverse = append(reverse, strings.ToUpper(arrInner[j]))
			} else {
				reverse = append(reverse, strings.ToLower(arrInner[j]))
			}
		}
		// Join every string in reverse into outInner
		outInner := strings.Join(reverse, "")
		// Join every word in outInner into arr
		arr = append(arr, outInner)
	}
	// Join every word in arr into one string in output with space between
	output := strings.Join(arr, " ")
	// Return reversed word
	return output
}

// Main function
func main() {
	// Insert input word
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Input words 	: ")
	scanner.Scan()
	numInp := scanner.Text()
	// Reverse inputted word using alternateCase function and print result
	fmt.Println("Output words 	:", alternateCase(numInp))
}
