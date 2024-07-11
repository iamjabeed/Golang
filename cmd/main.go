package main

import (
	"fmt"
	"os"
)

func main() {
	//  go run ./cmd/main.go Jaddu
    // os.Args[0] = project file name
    // os.Args[1] = First argument from cmd line  
	fmt.Println("Hello",os.Args[1])
}