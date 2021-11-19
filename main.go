package main

import (
	"fmt"
)

func main() {

	var addArray [5]string
	addArray[0] = "NAME:"
	addArray[1] = "CITY:"
	addArray[2] = "DISTRICT:"
	addArray[3] = "STATE:"
	addArray[4] = "COUNTRY:"

	fmt.Println(addArray[0])
	fmt.Scanln(&addArray[0])
        fmt.Println(addArray[1])
        fmt.Scanln(&addArray[1])
        fmt.Println(addArray[2])
        fmt.Scanln(&addArray[2])
        fmt.Println(addArray[3])
        fmt.Scanln(&addArray[3])
        fmt.Println(addArray[4])
        fmt.Scanln(&addArray[4])

}

