package bytecode

import (
	"fmt"
	"os"
)

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
func (bc *ByteCode) Write(dest string) {
	destFile, err := os.Create(dest)
	if err != nil {
		fmt.Println("Ֆայլը ստեղծելու սխալ։")
		return
	}
	defer destFile.Close()

	destFile.Write(bc.code)
}

// AddByte ...
func (bc *ByteCode) AddByte(v byte) {
	bc.code = append(bc.code, v)
}

// AddInteger ...
func (bc *ByteCode) AddInteger(v int) {
	bc.code = append(bc.code, byte(v), byte(v>>8), byte(v>>16), byte(v>>24))
}

// Size ...
func (bc *ByteCode) Size() int {
	return len(bc.code)
}

// SetInteger ...
func (bc *ByteCode) SetInteger(ix int, v int) {
	bc.code[ix] = byte(v)
	bc.code[ix+1] = byte(v >> 8)
	bc.code[ix+2] = byte(v >> 16)
	bc.code[ix+3] = byte(v >> 24)
}

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
