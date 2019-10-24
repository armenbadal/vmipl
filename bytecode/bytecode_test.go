package bytecode

import (
	"testing"
)

func TestNewBuilder(t *testing.T) {
	bi := NewBuilder()
	bi.LoadC(16)
	bi.LoadC(32)
	bi.Add()
	bi.Halt()

	expected := []byte{
		0x13, 0x10, 0x00, 0x00, 0x00, 0x13, 0x20, 0x00,
		0x00, 0x00, 0x01, 0x0b}
	if len(expected) != len(bi.code) {
		t.Fail()
	}

	for i, b := range bi.code {
		if expected[i] != b {
			t.Fail()
		}
	}
}

func TestLabelAndJump(t *testing.T) {
	bi := NewBuilder()

	bi.LoadC(128)
	bi.SetLabel("begin")
	bi.LoadC(64)
	bi.Jump("begin")
	bi.Jump("end")
	bi.SetLabel("end")
	bi.Halt()

	expected := []byte{
		0x13, 0x80, 0x00, 0x00, 0x00, 0x13, 0x40, 0x00,
		0x00, 0x00, 0x0c, 0x05, 0x00, 0x00, 0x00, 0x0c,
		0x14, 0x00, 0x00, 0x00, 0x0b}
	if len(expected) != len(bi.code) {
		t.Fail()
	}

	for i, b := range bi.code {
		if expected[i] != b {
			t.Fail()
		}
	}
}
