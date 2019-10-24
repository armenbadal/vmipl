package main

import "github.com/armenbadal/vmipl/machine"

func main() {
	println("IPL VM\n======")

	// bc := assembler.Assemble("examples/ex0.am")
	// bc.Dump()
	// bc.Write("examples/ex0.bc")

	mc := machine.Create()
	code := []byte{
		machine.LoadC,
		0x04, 0x00, 0x00, 0x00,
		machine.LoadC,
		0x02, 0x00, 0x00, 0x00,
		machine.Add,

		machine.LoadC,
		0xFF, 0x00, 0x00, 0x00,
		machine.LoadC,
		0x01, 0x00, 0x00, 0x00,
		machine.Add,

		machine.Halt,
	}
	mc.Code(code)
	mc.Execute()
}
