package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}


func main()  {
	
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName:"Alex", lastName:"Anderson"}
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	naumche := person{
		firstName: "Naumche",
		lastName: "Joshevski",
		contactInfo: contactInfo {
			email: "naumche@gmail.com",
			zipCode: 32545,
		},
	}

	naumche.updateName("Nomsa")
	naumche.print()
	
}



func (pointerToPerson *person) updateName(newFirstName string){
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print()  {
	fmt.Printf("%+v", p)
}