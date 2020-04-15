package main

import "fmt"

// Initialize a MAP using a composite literal where the key is a string and the value is an int;
// print out the map;
// range over the map printing out just the key;
// range over the map printing out both the key and the value

func main() {

	xm := map[string]int{"Peter": 43, "Zsofi": 36, "Vince": 4}

	fmt.Println(xm)

	for k := range xm {
		fmt.Println(k)
	}

	for k, v := range xm {
		fmt.Println(k, ":", v)
	}
}
