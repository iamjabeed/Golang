package main

import "fmt"

func main() {
	var arr [2]int
	arr[0] =1
	arr[1] =5

	fmt.Println(arr)

	var arr2 = [2]string{"Jaddu", "Jhon"}
	fmt.Println(arr2)

	arr3:=[2]int{5,6}
    fmt.Println(arr3)
    fmt.Println("The length of array is: ",len(arr3))

	// Iteration

	for i:=0; i<len(arr2); i++{
		fmt.Println(arr2[i])
	}

	// For range

	// for index, value:=range arr2{
	// 	fmt.Printf("Index: %d, Value: %d\n", index, value)
	// }
	arr4 := [5]int{1, 2, 3, 4, 5}
    for index, value := range arr4 {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }

}