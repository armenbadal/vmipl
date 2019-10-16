package main

import "github.com/armenbadal/iplvm/assembler"

func main() {
	println("IPL VM")
	assembler.Assemble("ex0.am", "")
}
