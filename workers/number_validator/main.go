package main

import (
	"fmt"
	"os"
)

var configuration = getProgramConfiguration()

func main(){
	//fmt.Printf("Hello!\n")
	args := os.Args[1:]
	//fmt.Printf("Args: %v", args)
	phoneNumber := args[0]
	number := ValidateNumber(phoneNumber)
	//fmt.Printf("Number: %v\n",number)
	if number.Number == "" {
		//fmt.Printf("Error: inconcistent_number\n")
		number.Type = "inconcistent_number"
	}
	//fmt.Printf("NumberType: %v",number.Type)
	//fmt.Printf("SET VARIABLE number_type \"%v\"\n", number.Type)
	fmt.Printf("%v\n", number.Type)
}