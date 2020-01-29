package main

import "fmt"

// Person : a person
type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

// ContactInfo : a persons contact info
type ContactInfo struct {
	email   string
	zipCode int
}

func main() {
	trevor := Person{
		firstName: "Trevor",
		lastName:  "Johnson",
		contact: ContactInfo{
			email:   "trevor@johnson.com",
			zipCode: 85260,
		},
	}

	// trevorPointer := &trevor
	trevor.updateName("Eric")
	trevor.print()
}

func (pointerToPerson *Person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}
