package main

import "fmt"

func main() {
	age := 22

	if age < 13 {
        fmt.Println("You are a child.")
    } else if age >= 13 && age < 18 {
        fmt.Println("You are a teenager.")
    } else {
        fmt.Println("You are an adult.")
    }

	for i := 0; i < 10; i++ {
		if i==4{
			continue
		}
        fmt.Println(i)
    }

	// For loops as while loop
	i := 0
    for i < 5 {
        fmt.Println(i)
        i++
    }

	day := 5

    switch day {
    case 1:
        fmt.Println("Today is Mon.")
    case 2:
        fmt.Println("Today is Tues.")
    case 3:
        fmt.Println("Today is Wed.")
    case 4:
        fmt.Println("Today is Thu")
    case 5:
        fmt.Println("Today is Fri.")
    case 6:
        fmt.Println("Today is Sat.")
    case 7:
        fmt.Println("Today is Sun.")
    default:
        fmt.Println("It's another day.")
    }
}