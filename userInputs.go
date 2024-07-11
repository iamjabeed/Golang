package main

import "fmt"

func main() {
	var firstName string
	var lastName string
	var age int
	var email string

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	fmt.Print("Enter your email: ")
	fmt.Scanln(&email)

	fmt.Printf("Hello %s %s youre %d old, This is your email address %s.", firstName, lastName, age, email)
}