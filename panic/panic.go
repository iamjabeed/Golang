// package main

// func employee(name *string, age int) {
// 	if age > 65 {
// 		panic("Age can't be greater than 65")
// 	}
// }
// func main() {
// 	name := "Alice"
// 	age := 68

// 	employee(&name, age)
// }

// panic with defer
package main

import "fmt"

func main() {
    defer fmt.Println("This line will be printed last, before the program exits")

    fmt.Println("Start of main function")

    panic("Something went wrong!")

    fmt.Println("This line will never be executed")
}
