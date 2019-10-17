package main

import "github.com/armenbadal/vmipl/assembler"

func main() {
	println("IPL VM\n------")
	assembler.New("examples/ex0.am", "")
}
