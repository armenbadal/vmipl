package bytecode

import (
	"fmt"

	"github.com/armenbadal/vmipl/machine"
)

// Builder կառուցվածքն ապահովում է բայթ-կոդի գեներացիայի մեթոդներ
type Builder struct {
	addresses  map[string]int
	unresolved map[string][]int

	code []byte
}

// NewBuilder ֆունկցիան ստեղծում է բայթ-կոդի գեներացիայի նոր օբյեկտ
func NewBuilder() *Builder {
	builder := new(Builder)
	builder.addresses = make(map[string]int)
	builder.unresolved = make(map[string][]int)
	builder.code = make([]byte, 0, 1024)
	return builder
}

// Code ֆունկցիան վերադարձնում է կառուցված բայթ-կոդը, կամ ազդարարում է սխալի մասին
func (bi *Builder) Code() ([]byte, error) {
	if len(bi.unresolved) != 0 {
		return nil, fmt.Errorf("Որոշ պիտակներ սահմանված չեն։")
	}

	return bi.code, nil
}

// Add գումարման գործողությունն է
func (bi *Builder) Add() { bi.addByte(machine.Add) }

// And կոնյունկցիայի գործողությունն է
func (bi *Builder) And() { bi.addByte(machine.And) }

// Alloc ֆունկցիայի պարամետրերի համար տեղ ռեզերվացնող հրամանն է
func (bi *Builder) Alloc(n int) {
	bi.addByte(machine.Alloc)
	bi.addInt(n)
}

// Call ֆունկցիայի կանչի հրամանն է
func (bi *Builder) Call() { bi.addByte(machine.Call) }

// Div բաժանման գործողությունն է
func (bi *Builder) Div() { bi.addByte(machine.Div) }

// Dup ստեկի գագաթի տարրը կրկնելու հրամանը
func (bi *Builder) Dup() { bi.addByte(machine.Dup) }

// Enter ...
func (bi *Builder) Enter(n int) {
	bi.addByte(machine.Enter)
	bi.addInt(n)
}

// Eq հավասրության ստուգման գործողությունն է
func (bi *Builder) Eq() { bi.addByte(machine.Eq) }

// Geq մեծ կամ հավասար լինելը ստուգող գործողությունն է
func (bi *Builder) Geq() { bi.addByte(machine.Geq) }

// Gr մեծ լինելը ստուգող գործողությունն է
func (bi *Builder) Gr() { bi.addByte(machine.Gr) }

// Halt աշխատանքի ավարտի հրամանն է
func (bi *Builder) Halt() { bi.addByte(machine.Halt) }

// Jump առանց պայմանի անցման հրամանն է
func (bi *Builder) Jump(label string) {
	bi.addByte(machine.Jump)
	bi.addInt(bi.addressOf(label))
}

// JumpI ...
func (bi *Builder) JumpI(label string) {
	bi.addByte(machine.JumpI)
	bi.addInt(bi.addressOf(label))
}

// JumpZ անցում, եթե ստեկի գագաթին զրո է
func (bi *Builder) JumpZ(label string) {
	bi.addByte(machine.JumpZ)
	bi.addInt(bi.addressOf(label))
}

// Leq ...
func (bi *Builder) Leq() {}

// Le ...
func (bi *Builder) Le() {}

// Load ...
func (bi *Builder) Load() {}

// LoadA ...
func (bi *Builder) LoadA() {}

// LoadC ստեկում ավելացնում է տրված հաստատունը
func (bi *Builder) LoadC(n int) {
	bi.addByte(machine.LoadC)
	bi.addInt(n)
}

// LoadR ...
func (bi *Builder) LoadR() {}

// LoadRC ...
func (bi *Builder) LoadRC() {}

// Malloc ...
func (bi *Builder) Malloc() {}

// Mark ...
func (bi *Builder) Mark() {}

// Mod ...
func (bi *Builder) Mod() {}

// Mul ...
func (bi *Builder) Mul() {}

// Neg ...
func (bi *Builder) Neg() {}

// Neq ...
func (bi *Builder) Neq() {}

// New ...
func (bi *Builder) New() {}

// Not ժխտման գործողությունն է
func (bi *Builder) Not() { bi.addByte(machine.Not) }

// Or դիզյունկցիայի գործողությունն է
func (bi *Builder) Or() { bi.addByte(machine.Or) }

// Pop ...
func (bi *Builder) Pop() {}

// Return ...
func (bi *Builder) Return() {}

// Slide ...
func (bi *Builder) Slide(a, b int) {
	bi.addByte(machine.Slide)
	bi.addInt(a)
	bi.addInt(b)
}

// Store ...
func (bi *Builder) Store() {}

// StoreA ...
func (bi *Builder) StoreA() {}

// StoreR ...
func (bi *Builder) StoreR() {}

// Sub ...
func (bi *Builder) Sub() { bi.addByte(machine.Sub) }

// SetLabel ֆունկցիան կոդի գեներացիայի ընթացիկ հասցեն նշում է տրված պիտակով
func (bi *Builder) SetLabel(label string) error {
	if _, marked := bi.addresses[label]; marked {
		return fmt.Errorf("%s պիտակն արդեն սահմանված է։", label)
	}

	addr := len(bi.code)
	bi.addresses[label] = addr

	if els, found := bi.unresolved[label]; found {
		for _, ps := range els {
			bi.setInt(addr, ps)
		}
		delete(bi.unresolved, label)
	}

	return nil
}

func (bi *Builder) addByte(v byte) {
	bi.code = append(bi.code, v)
}

func (bi *Builder) addInt(v int) {
	bi.code = append(bi.code, byte(v), byte(v>>8), byte(v>>16), byte(v>>24))
}

func (bi *Builder) setInt(v int, i int) {
	bi.code[i+0] = byte(v)
	bi.code[i+1] = byte(v >> 8)
	bi.code[i+2] = byte(v >> 16)
	bi.code[i+3] = byte(v >> 24)
}

func (bi *Builder) addressOf(label string) int {
	if addr, marked := bi.addresses[label]; marked {
		return addr
	}

	if _, exists := bi.unresolved[label]; !exists {
		bi.unresolved[label] = make([]int, 0, 8)
	}
	bi.unresolved[label] = append(bi.unresolved[label], len(bi.code))

	return 0
}
