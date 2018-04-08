package main

import (
	"fmt"
)

type state struct {
	symbol rune
	edge1  *state
	edge2  *state
}

type nFrag struct {
	initial *state
	accept  *state
}

func poregtonfa(postfix string) *nFrag {
	nStack := []*nFrag{}

	for _, r := range postfix {
		switch r {
		case '.':
			//Pops 2 Fragments of the stack
			frag2 := nStack[len(nStack)-1]
			nStack = nStack[:len(nStack)-1]
			frag1 := nStack[len(nStack)-1]
			nStack = nStack[:len(nStack)-1]

			//joins the accept state of frag1 to the initial state of frag2
			frag1.accept.edge1 = frag2.initial

			nStack = append(nStack, &nFrag{initial: frag1.initial, accept: frag2.accept})
		case '|':
			//Pops 2 Fragments of the stack
			frag2 := nStack[len(nStack)-1]
			nStack = nStack[:len(nStack)-1]
			frag1 := nStack[len(nStack)-1]
			nStack = nStack[:len(nStack)-1]

			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			accept := state{}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			nStack = append(nStack, &nFrag{initial: &initial, accept: &accept})
		case '*':
			frag := nStack[len(nStack)-1]
			nStack = nStack[:len(nStack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nStack = append(nStack, &nFrag{initial: &initial, accept: &accept})
		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}

			nStack = append(nStack, &nFrag{initial: &initial, accept: &accept})
		}
	}

	return nStack[0]
}

func main() {
	nFrag := poregtonfa("ab.c*|")
	fmt.Println(nFrag)
}
