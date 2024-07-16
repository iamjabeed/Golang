package main

import "fmt"

// Define a User interface
type User interface {
	GetRole() string
	GetName() string
}

// Define Admin struct
type Admin struct {
	Name string
}

// Implement the GetRole method for Admin
func (a Admin) GetRole() string {
	return "Admin"
}

// Implement the GetName method for Admin
func (a Admin) GetName() string {
	return a.Name
}

// Define Guest struct
type Guest struct {
	Name string
}

// Implement the GetRole method for Guest
func (g Guest) GetRole() string {
	return "Guest"
}

// Implement the GetName method for Guest
func (g Guest) GetName() string {
	return g.Name
}

func printRole(u User) {
    fmt.Printf("%s is a %s\n", u.GetName(), u.GetRole())
}

func main() {
	admin := Admin{Name: "Jaddu"}
	guest := Guest{Name: "Virat"}

	printRole(admin)
	// Jaddu is an Admin
	printRole(guest)
	// Virat is a Guest
}
