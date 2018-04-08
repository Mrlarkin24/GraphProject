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

	if len(nStack) != 1 {
		fmt.Println("Uh oh:", len(nStack), nStack)
	}

	return nStack[0]
}

func addState(l []*state, s *state, a *state) []*state {
	l = append(l, s)

	if s != a && s.symbol == 0 {
		l = addState(l, s.edge1, a)
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}

	return l
}

func pomatch(po string, s string) bool {
	ismatch := false
	ponfa := poregtonfa(po)

	current := []*state{}
	next := []*state{}

	current = addState(current[:], ponfa.initial, ponfa.accept)

	for _, r := range s {
		for _, c := range current {
			if c.symbol == r {
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		current, next = next, []*state{}
	}

	for _, c := range current {
		if c == ponfa.accept {
			ismatch = true
			break
		}
	}

	return ismatch
}

func main() {
	fmt.Println(pomatch("ab.c*|", ""))
}
