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

	//For loop so you dont have to restart the code each time
	for regular != "" && userString != "" {
		fmt.Println("\nLEAVE BOTH USER INPUTS BLANK TO EXIT PROGRAM!")
		fmt.Println("--------------")
		fmt.Println("| Main Menu |")
		fmt.Println("--------------")

		fmt.Println("Enter Regular Expression:")
		scanner.Scan()
		regular := scanner.Text()

		fmt.Println("\nEnter String:")
		scanner.Scan()
		userString := scanner.Text()

		//if and else to stop the other files running if the user does enter anything
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
