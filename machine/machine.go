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

execution:
	for {
		opcode := m.memory[m.pc]
		m.pc++
		switch opcode {
		case Add:
			n0 := m.pop()
			n1 := m.pop()
			m.push(n0 + n1)
		case And:
			b0 := 0 != m.pop()
			b1 := 0 != m.pop()
			if b0 && b1 {
				m.push(1)
			} else {
				m.push(0)
			}
		case Halt:
			break execution
		case Load:
			addr := m.pop()
			vl := getInteger(m.memory[addr : addr+4])
			m.push(vl)
		case LoadC:
			nv := getInteger(m.memory[m.pc : m.pc+4])
			m.pc += 4
			m.push(nv)
		case Mul:
			n0 := m.pop()
			n1 := m.pop()
			m.push(n0 * n1)
		case Neg:
			n0 := m.pop()
			m.push(-n0)
		case Not:
			n0 := m.pop()
			if n0 == 0 {
				m.push(1)
			} else {
				m.push(0)
			}
		case Store:
			addr := m.pop()
			vl := m.pop()
			putInteger(vl, m.memory[addr:addr+4])
		}
	}

	m.printStack()
}

func (m *Machine) printStack() {
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

func (m *Machine) push(v int) {
	m.sp -= 4
	putInteger(v, m.memory[m.sp:m.sp+4])
}

func (m *Machine) pop() int {
	n := getInteger(m.memory[m.sp : m.sp+4])
	m.sp += 4
	return n
}
