package main

import "fmt"

// Create a new type called “gator”. The underlying type of “gator” is an int.
// Using var, declare an identifier “g1” as type gator (var g1 gator).
// Assign a value to “g1”. Print out “g1”.
// Print the type of “g1” using fmt.Printf(“%T\n”, g1)
// Using var, declare an identifier “x” as type int (var x int).
// Print out “x”.
// Print the type of “x” using fmt.Printf(“%T\n”, x)
// Now add a method to type gator with this signature ...
// greeting()
// and have it print “Hello, I am a gator”.
// create another type called “flamingo”.
// Make the underlying type of “flamingo” bool.
// Give “flamingo” a method with this signature …
// greeting()
// and have it print “Hello, I am pink and beautiful and wonderful.”
// create a new type “swampCreature”. The underlying type of “swapCreature” is interface.
// The behavior which the “swampCreature” interface defines is such that any type which has this method …
// greeting()
// will implicitly implement the “swampCreature” interface.
// Create a func called “bayou” which takes a value of type “swampCreature” as an argument.
// Have this func print out the greeting for whatever “swampCreature” might be passed in.

type gator int
type flamingo bool

type swampCreature interface {
	greeting()
}

func (g gator) greeting() {
	fmt.Println("Hello, I am gator.")
}

func (f flamingo) greeting() {
	fmt.Println("Hello, I am pink and beautiful and wonderful.")
}

func bayou(s swampCreature) {
	s.greeting()
}

func main() {
	var g1 gator
	var f1 flamingo
	var x int

	g1 = 7
	// cannot assign g1 to x without conversion, because it is a different type
	x = int(g1)
	fmt.Printf("%v\t%T\n", g1, g1)
	fmt.Printf("%v\t%T\n", x, x)

	g1.greeting()
	bayou(f1)
}
