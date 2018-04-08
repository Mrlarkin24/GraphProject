package main

import (
	"fmt"
)

func intopost(infix string) string {
	//Keeps track of the special characters and sets their precedent
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}

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

	for len(s) > 0 {
		postfix = append(postfix, s[len(s)-1])
		s = s[:len(s)-1]
	}

	return string(postfix)
}

func main() {
	//Answer: ab.c*.
	fmt.Println("Infix:   ", "a.b.c*")
	fmt.Println("Postfix: ", intopost("a.b.c*"))

	//Answer: abd|.*
	fmt.Println("Infix:   ", "(a.(b|d))*")
	fmt.Println("Postfix: ", intopost("(a.(b|d))*"))

	//Answer: abd|.c*.
	fmt.Println("Infix:   ", "a.(b|d).c*")
	fmt.Println("Postfix: ", intopost("a.(b|d).c*"))

	//Answer: abb.+.c.
	fmt.Println("Infix:   ", "a.(b.b)+.c")
	fmt.Println("Postfix: ", intopost("a.(b.b)+.c"))
}
