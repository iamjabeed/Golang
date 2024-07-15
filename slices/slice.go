package main

import "fmt"

func main() {
	//From an array
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	slice1 :=arr[1:3]
	fmt.Println(slice1)

	// using make 
	sliceUsingMake := make([]int, 5, 10)
	fmt.Println(sliceUsingMake)

	marks := []int{22,18, 24, 20, 19}
    fmt.Println(marks)

	array := [5]int{1, 2, 3, 4, 5}
    slice := []int{1, 2, 3, 4, 5}
	fmt.Println("this is a slice", slice)
	fmt.Println("this is an array", array)
}