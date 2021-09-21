package main

import (
	"vmipl/bytecode"
	"vmipl/machine"
)

func main() {
	println("IPL VM\n======")

	// bc := assembler.Assemble("examples/ex0.am")
	// bc.Dump()
	// bc.Write("examples/ex0.bc")

	bb := bytecode.NewBuilder()
	bb.LoadC(4)
	bb.LoadC(2)
	bb.Add()
	bb.LoadC(0xFF)
	bb.LoadC(0x01)
	bb.Add()
	bb.Halt()

	if code, err := bb.Code(); err == nil {
		mc := machine.Create(code)
		mc.Execute()
	}
}
