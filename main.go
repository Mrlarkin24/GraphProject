package main

import (
	"bufio"
	"fmt"
	"os"
	//links the other files to main
	L "../GraphProject/lib"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("--------------")
	fmt.Println("| Main Menu |")
	fmt.Println("--------------")

	fmt.Println("Enter Regular Expression:")
	scanner.Scan()
	regular := scanner.Text()

	fmt.Println("\nEnter String:")
	scanner.Scan()
	userString := scanner.Text()

	//fmt.Println("\nRegular: ", regular)
	fmt.Println("User String: ", userString)

	//Calls the function Intopost in shunt.go when it returns it calls Pomatch in rega.go
	fmt.Println("\nMatch: ", L.Pomatch(L.Intopost(regular), userString))

}
