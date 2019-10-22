package bytecode

import "fmt"

// ByteCode ...
type ByteCode struct {
	code []byte
}

// New ...
func New(cp int) *ByteCode {
	bc := new(ByteCode)
	bc.code = make([]byte, 0, cp)
	return bc
}

// Read ...
func (bc *ByteCode) Read(src string) {}

// Write ...
func (bc *ByteCode) Write(dest string) {}

// AddByte ...
func (bc *ByteCode) AddByte(v byte) {
	bc.code = append(bc.code, v)
}

// AddInteger ...
func (bc *ByteCode) AddInteger(v int) {}

// Size ...
func (bc *ByteCode) Size() int {
	return len(bc.code)
}

// SetInteger ...
func (bc *ByteCode) SetInteger(ix int) {}

// Dump ...
func (bc *ByteCode) Dump() {
	col := 0
	for i := 0; i < bc.Size(); i++ {
		if col == 0 {
			fmt.Printf("0x%04x ", i)
		}
		fmt.Printf("%02x ", bc.code[i])
		col++
		if col == 16 {
			fmt.Println("")
			col = 0
		}
	}
	fmt.Println("")
}
