package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type newPerson struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	smit := person{firstName: "Smit", lastName: "Contractor"}
	fmt.Println(smit.firstName)
	fmt.Println(smit)
	smit.print()
	fmt.Println()

	var guddu person
	fmt.Printf("%+v", guddu)
	fmt.Println()
	guddu.firstName = "Guddu"
	guddu.lastName = "Contractor"
	fmt.Printf("%+v", guddu)
	fmt.Println()

	jaanu := person{
		firstName: "Jannu",
		lastName:  "Patel",
		contact: contactInfo{
			email: "j@gmail.com",
		},
	}
	fmt.Printf("%+v", jaanu)
	fmt.Println()

	jenny := person{
		firstName: "Jenny",
		lastName:  "Patel",
		contact: contactInfo{
			email:   "jp@gmail.com",
			zipCode: 123,
		},
	}
	fmt.Printf("%+v", jenny)
	fmt.Println()

	newSmit := newPerson{
		firstName: "Smit",
		lastName:  "Contractor",
		contactInfo: contactInfo{
			email:   "sc@gmail.com",
			zipCode: 123,
		},
	}
	fmt.Printf("%+v", newSmit)
	fmt.Println()

	jenny.updateFirstName("Lalu")
	jenny.print()
}

func (p *person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
