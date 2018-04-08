package main

import (
	"bufio"
	"fmt"
	"os"
	//links the other files to main
	L "../GraphProject/lib"
)

func main() {
	var regular string = "a.b"
	var userString string = "ab"

	scanner := bufio.NewScanner(os.Stdin)

	for regular != "" && userString != "" {
		fmt.Println("\nLEAVE BOTH USER INPUT BLANK TO EXIT PROGRAM!")
		fmt.Println("--------------")
		fmt.Println("| Main Menu |")
		fmt.Println("--------------")

		fmt.Println("Enter Regular Expression:")
		scanner.Scan()
		regular := scanner.Text()

		fmt.Println("\nEnter String:")
		scanner.Scan()
		userString := scanner.Text()

		if regular == "" && userString == "" {
			fmt.Println("Goodbye!")
			os.Exit(0)
		} else {
			//fmt.Println("\nRegular: ", regular)
			fmt.Println("User String: ", userString)

			//Calls the function Intopost in shunt.go when it returns it calls Pomatch in rega.go
			fmt.Println("\nMatch: ", L.Pomatch(L.Intopost(regular), userString))
		}
	}

}
