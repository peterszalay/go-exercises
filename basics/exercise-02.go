package main

import "fmt"

// create a struct that holds person fields
// create a struct that holds secret agent fields and embeds person type
// attach a method to person: pSpeak
// attach a method to secret agent: saSpeak
// create a variable of type person
// create a variable of type secret agent
// print a field from person
// run pSpeak attached to the variable of type person
// print a field from secret agent
// run saSpeak attached to the variable of type secret agent
// run pSpeak attached to the variable of type secret agent

type person struct {
	firstName string
	lastName  string
}

type secretAgent struct {
	person
	licenceToKill bool
}

func (p person) pSpeak(s string) string {
	return fmt.Sprint(p.firstName, " ", p.lastName, " says: ", s)
}

func (sa secretAgent) saSpeak(s string) string {
	return fmt.Sprint(sa.firstName, " ", sa.lastName, " says: ", s)
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

	fmt.Println(mp.pSpeak("Hi, " + jb.firstName + "!"))
	fmt.Println(jb.saSpeak("Oh hello, " + mp.firstName + " " + mp.lastName + "!"))
	fmt.Println(jb.pSpeak("Is M in the office, " + mp.lastName + "?"))
}
