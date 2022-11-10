package main

import (
	"fmt"
	"sort"
	"unicode"
)

func IsUpper(s string) bool {
	for i, r := range s {
		fmt.Println("i : ", i)
		fmt.Println("r : ", r)
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for j, z := range s {
		fmt.Println("j : ", j)
		fmt.Println("z : ", z)
		if !unicode.IsLower(z) && unicode.IsLetter(z) {
			return false
		}
	}
	return true
}

func main() {
	//x, y := IsUpper("VANEEr"), IsLower("roflstomPt")
	x := []rune{}
	x = append(x, rune('a'))
	x = append(x, rune('A'))
	x = append(x, rune('b'))
	x = append(x, rune('B'))
	x = append(x, rune('c'))
	x = append(x, rune('C'))
	fmt.Println(x)

	first := "S"
	second := "H"
	sortFS := []string{first, second}
	sort.Sort(sort.StringSlice(sortFS))
	first = sortFS[0]
	second = sortFS[1]
	fmt.Println(first, second)
	//fmt.Println(y)
}
