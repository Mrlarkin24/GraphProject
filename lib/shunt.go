package shunt

import (
	"fmt"
)

//Intopost called by main
func Intopost(infix string) string { //function to convert inflix to postfix
	//Keeps track of the special characters and sets their precedent
	specials := map[rune]int{'*': 10, '+': 9, '.': 8, '|': 7}

	//Initialises the runes as empty
	postfix, s := []rune{}, []rune{}

	//Range converts a string into an array of runes
	for _, r := range infix {
		//Shunting yard algorithm
		switch {
		case r == '(':
			s = append(s, r)
		case r == ')':
			for s[len(s)-1] != '(' {
				postfix = append(postfix, s[len(s)-1])
				s = s[:len(s)-1]
			}
			s = s[:len(s)-1]
		case specials[r] > 0:
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				postfix = append(postfix, s[len(s)-1])
				s = s[:len(s)-1]
			}
			s = append(s, r)
		default:
			postfix = append(postfix, r) //Puts 'r' at the end of postfix
		}
	}

	//when the stacks length is greater than 0
	for len(s) > 0 {
		//Appends the top element of the stack to postfix
		postfix = append(postfix, s[len(s)-1])

		//Removes the top element from the stack
		s = s[:len(s)-1]
	}

	fmt.Println("\nInfix:   ", infix)
	fmt.Println("Postfix: ", string(postfix))

	return string(postfix)
}
