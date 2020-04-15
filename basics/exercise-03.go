package main

import "fmt"

// create an interface type that both person and secretAgent implement
// declare a func with a parameter of the interface’s type
// call that func in main and pass in a value of type person
// call that func in main and pass in a value of type secretAgent

type person struct {
	firstName string
	lastName  string
}

type secretAgent struct {
	person
	licenceToKill bool
}

type human interface {
	speak(s string) string
}

func (p person) speak(s string) string {
	return fmt.Sprint(p.firstName, " ", p.lastName, " says: ", s)
}

func (sa secretAgent) speak(s string) string {
	return fmt.Sprint("Secret Agent ", sa.firstName, " ", sa.lastName, " says: ", s)
}

func main() {
	mp := person{
		firstName: "Miss",
		lastName:  "MoneyPenny",
	}

	jb := secretAgent{
		person:        person{firstName: "James", lastName: "Bond"},
		licenceToKill: true,
	}

	fmt.Println("person's first name is", mp.firstName)
	fmt.Println("secret agent has a licence to kill:", jb.licenceToKill)

	fmt.Println(mp.speak("Hi, " + jb.firstName + "!"))
	fmt.Println(jb.speak("Oh hello, " + mp.firstName + " " + mp.lastName + "!"))
	fmt.Println(jb.speak("Is M in the office, " + mp.lastName + "?"))
}
