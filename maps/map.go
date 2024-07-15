package main

import "fmt"

func main() {
	studentInfo := map[string]int{
		"Jaddu": 24,
		"Sara":  22,
		"Bob":   21,
		"Max":   25,
	}
	// fmt.Println(studentInfo)

	studentInfo["Jaddu"]=22
	fmt.Println(studentInfo)

	fmt.Println(studentInfo["Jaddu"])

	stdentMarks := map[int]int{
		1:89,
		2:78,
		3:92,
		4:94,
	}
	fmt.Println(stdentMarks[2])

	// iterate through map
	for index, value := range stdentMarks{
		
		fmt.Println(index, value)

		// fmt.Println("index: ", index)
		// fmt.Println("value: ", value)
	}
}