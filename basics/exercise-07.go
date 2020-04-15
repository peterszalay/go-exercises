package main

import "fmt"

// Create a new type: vehicle. The underlying type is a struct. The fields: doors, color.
// Create two new types: truck & sedan. The underlying type of each of these new types is a struct.
// Embed the “vehicle” type in both truck & sedan.
// Give truck the field “fourWheel” which will be set to bool.
// Give sedan the field “luxury” which will be set to bool.
// Using the vehicle, truck, and sedan structs:
// using a composite literal, create a value of type truck and assign values to the fields;
// using a composite literal, create a value of type sedan and assign values to the fields.
// Print out each of these values.
// Print out a single field from each of these values.
// Give a method to both the “truck” and “sedan” types with the following signature transportationDevice() string
// Have each func return a string saying what they do.
// Call the method for each value.
// Create a new type called “transportation”. The underlying type of this new type is interface.
// Create a func called “report” that takes a value of type “transportation” as an argument.
// The func should print the string returned by “transportationDevice()” being called for any type that implements the “transportation” interface.
// Call “report” passing in a value of type truck.
// Call “report” passing in a value of type sedan.

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

type transportation interface {
	transportationDevice() string
}

func (t truck) transportationDevice() string {
	return fmt.Sprint("The truck is delivering goods...")
}

func (s sedan) transportationDevice() string {
	return fmt.Sprint("The sedan is carrying people to their destination...")
}

func report(t transportation) {
	fmt.Println(t.transportationDevice())
}

func main() {
	t := truck{
		vehicle:   vehicle{2, "blue"},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{5, "silver"},
		luxury:  false,
	}

	fmt.Println(t)
	fmt.Println(s)
	fmt.Println("Sedan color:", s.color)
	fmt.Println("Truck number of doors:", t.doors)
	fmt.Println(t.transportationDevice())
	fmt.Println(s.transportationDevice())
	fmt.Println("Calling report:")
	report(s)
	report(t)

}
