package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("welcome to file handling in GO:")

	// Create a file named "go.txt" 
	file, err:= os.Create("./files/go2.txt")
	// content :="This is dummy text to write in the file created by go"

	// if err!=nil{
	// 	fmt.Println("error while creating file")
	// 	panic(err)
	// }

	checkErrorNil(err)

	// defer file.Close()

	// Write some text to the file

	length, err := file.WriteString("Hello, this is a sample text file!")

	// length, err := io.WriteString(file,content)

	checkErrorNil(err)

	fmt.Println("File created and text written successfully")

	// fmt.Println("this is the length of file:", n)
	fmt.Println("this is the length of file:", length)

	defer file.Close()

	readFile("./files/go2.txt")

}

// utility func to check error
func checkErrorNil(err error) {
	if err != nil {
		fmt.Println("Error writing to file:")
		panic(err)
	}
}


// function to read the file 

func readFile(file string){
	n, err:= os.ReadFile(file)
	checkErrorNil(err)

	fmt.Println(string(n))
}
