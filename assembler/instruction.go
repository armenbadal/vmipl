package assembler

import "fmt"

type instruction struct {
	label   string
	name    string
	number0 int
	number1 int
	symbol  string
}

// for debug
func (ns instruction) print() {
	fmt.Printf("'%s'\t'%s'\t'%d'\t'%s'\n", ns.label, ns.name, ns.number0, ns.symbol)
}
