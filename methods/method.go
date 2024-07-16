package main

import "fmt"

// User struct
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// create an instance of Person
	user := Person{
		FirstName: "Jaddu",
		LastName:  "Syed",
		Age:       22,
	}

	// fmt.Println(user)
	// fmt.Println(user.FirstName)
	// fmt.Println(user.LastName)
	// fmt.Println(user.Age)

    // Calling the FullName method
	user.GetFullName()
	// Calling the IsAdult method
	fmt.Println(user.IsAdult())
}

// methods 
  // Method to get the full name of the person
  func (p Person) GetFullName(){
	fmt.Println(p.FirstName+" "+p.LastName)
  }
  // Method to check if the person is an adult

  func (p Person) IsAdult () bool{
	// fmt.Println(p.Age>=18)
	return p.Age>=18
  }