package main

import "github.com/armenbadal/vmipl/assembler"

func main() {
	println("IPL VM\n======")
	bc := assembler.Assemble("examples/ex0.am")
	bc.Dump()
	bc.Write("examples/ex0.bc")

	// mc := machine.Create()
	// code := []byte{
	// 	machine.LoadC,
	// 	0x00, 0x00, 0x00, 0x04,
	// 	machine.LoadC,
	// 	0x00, 0x00, 0x00, 0x02,
	// 	machine.Add,
	// 	machine.Halt,
	// }
	// mc.Code(code)
	// mc.Execute()
}
