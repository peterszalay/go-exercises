package main

import "fmt"

func main() {
	// Using the short declaration operator, create a variable with the identifier “s”
	// and assign “i'm sorry dave i can't do that” to “s”
	s := "i'm sorry dave i can't do that"
	fmt.Println(s)

	// Print “s” converted to a slice of byte.
	b := []byte(s)
	fmt.Println(b)

	// Print “s” converted to a slice of byte and then converted back to a string.
	fmt.Println(string([]byte(s)))

	// Using slicing, print just “i’m sorry dave”
	fmt.Println(s[0:14])

	// Using slicing, print just “dave i can’t”
	fmt.Println(s[10:22])

	// Using slicing, print just “can’t do that”
	fmt.Println(s[17:])

	// print every letter of “s” with one rune (character) on each line
	for _, r := range s[0:] {
		fmt.Println(string(r))
	}
}
