package machine

import "fmt"

// MemorySize հաստատունով տրվում է վիրտուալ մեքենայի
// հիշասարքի չափը բայթերով։
const MemorySize = 1024

// Machine ...
type Machine struct {
	ep int // extreme pointer
	fp int // frame pointer
	hp int // heap pointer
	pc int // program counter
	sp int // stack pointer

	program []byte
	memory  []int
}

// Create ...
func Create(code []byte) *Machine {
	mc := new(Machine)
	mc.program = make([]byte, MemorySize)
	mc.memory = make([]int, MemorySize)

	copy(mc.program, code)

	return mc
}

// Execute ...
func (m *Machine) Execute() {
	m.pc = 0
	m.sp = -1

	opcode := None
	for opcode != Halt {
		opcode = m.program[m.pc]
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
	m.sp++
	m.memory[m.sp] = v
}

func (m *Machine) pop() int {
	n := m.memory[m.sp]
	m.sp--
	return n
}

func (m *Machine) printStack() {
	fmt.Println("STACK\n-----")
	c := 0
	for i := 0; i <= m.sp; i++ {
		fmt.Printf("%08x ", m.memory[i])
		c++
		if c == 16 {
			fmt.Println()
			c = 0
		}
	}
	fmt.Println()
}
