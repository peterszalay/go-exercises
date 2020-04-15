package main

import "fmt"

// Initialize a SLICE of int using a composite literal;
// print out the slice;
// range over the slice printing out just the index;
// range over the slice printing out both the index and the value

func main() {
	xs := []int{11, 22, 33, 44, 55}

	fmt.Println(xs)
	fmt.Println("Just the indexes: ")
	for i := range xs {
		fmt.Print(i, "\t")
	}

	fmt.Println("\n", "Indexes and values: ")
	for i, v := range xs {
		fmt.Println(i, " \t", v)
	}
}
