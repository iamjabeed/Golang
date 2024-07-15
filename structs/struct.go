package main

import "fmt"

type Marks struct {
	english string
	math string
	science string
}

type Person struct {
	firstName string
	lastName  string
	age       int
	country   string
	marks     Marks
}

func main() {
	p := Person{
		firstName: "Jaddu",
		lastName:  "Syed",
		age:       22,
		country:   "India",
		marks: Marks{
			english: "94",
			math: "98",
			science: "89",
		},
	}

	fmt.Println(p)
	// fmt.Println(p.firstName)
	// fmt.Println(p.lastName)
	// fmt.Println(p.age)
	// fmt.Println(p.country)
	fmt.Println(p.marks)

}