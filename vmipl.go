package main

import "github.com/armenbadal/vmipl/assembler"

func main() {
	println("IPL VM\n------")
	assembler.Assemble("examples/ex0.am", "")
}
