package main

// Import package
import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

// Function to check if string is letter
func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// Function to find middle letter between two letter
func middle(first string, second string) string {
	//Create initial variable
	midArr := []string{}
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	arr := strings.Split(alphabet, "")
	insMid := false
	//Create array between two letter
	for _, r := range arr {
		if r == first || insMid == true {
			insMid = true
			midArr = append(midArr, r)
		}
		if r == second {
			break
		}
	}
	//Find middle letter
	var output string
	length := len(midArr)
	mid := len(midArr) / 2
	if length%2 == 0 {
		twoMid := []string{midArr[mid-1], midArr[mid]}
		output = strings.Join(twoMid, "")
	} else {
		output = midArr[mid]
	}
	// Return middle letter
	return output
}

// Main function
func main() {
	var first string
	var second string
	fmt.Printf("Insert first letter : ")
	fmt.Scanln(&first)
	// Check if first input is letter
	if IsLetter(first) == false {
		fmt.Printf("Error! Insert only letter!")
	} else {
		fmt.Printf("Insert second letter : ")
		fmt.Scanln(&second)
		// Check if second input is letter
		if IsLetter(second) == false {
			fmt.Printf("Error! Insert only letter!")
			// Check if both input is same
		} else if second == first {
			fmt.Printf("Error! Insert different letter!")
		} else {
			// Make both letter into uppercase
			first = strings.ToUpper(first)
			second = strings.ToUpper(second)
			// Sort letter
			sortFS := []string{first, second}
			sort.Sort(sort.StringSlice(sortFS))
			first = sortFS[0]
			second = sortFS[1]
			// Search middle letter with middle function
			output := middle(first, second)
			// Print result
			fmt.Printf("The middle between %s and %s is %s", first, second, output)
		}
	}

}
