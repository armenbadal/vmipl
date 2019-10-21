package machine

import "fmt"

const sizeOfMemory = 1024

// Machine ...
type Machine struct {
	ep int // extreme pointer
	fp int // frame pointer
	hp int // heap pointer
	pc int // program counter
	sp int // stack pointer

	memory []byte
}

// Create ...
func Create() *Machine {
	mc := new(Machine)
	mc.memory = make([]byte, sizeOfMemory)

	return mc
}

// Code ...
func (m *Machine) Code(code []byte) {
	for i := 0; i < len(code); i++ {
		m.memory[i] = code[i]
	}
}

// Execute ...
func (m *Machine) Execute() {
	m.pc = 0
	m.sp = len(m.memory)

	opcode := None
	for opcode != Halt {
		opcode = m.memory[m.pc]
		m.pc++
		operations[opcode](m)
	}

	m.printStack()
}

func (m *Machine) binary(op func(a, b int) int) {
	a0 := m.pop()
	a1 := m.pop()
	r := op(a0, a1)
	m.push(r)
}

func (m *Machine) push(v int) {
	m.sp -= 4
	putInteger(v, m.memory[m.sp:m.sp+4])
}

func (m *Machine) pop() int {
	n := getInteger(m.memory[m.sp : m.sp+4])
	m.sp += 4
	return n
}

func (m *Machine) printStack() {
	fmt.Println("STACK\n-----")
	c := 0
	for i := m.sp; i < len(m.memory); i++ {
		fmt.Printf("%02x ", m.memory[i])
		c++
		if c == 16 {
			fmt.Println()
			c = 0
		}
	}
	fmt.Println()
}
