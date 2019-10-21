package machine

// Մեքենայի հրամանների անուները (կոդերը)
const (
	None byte = iota
	Add
	And
	Alloc
	Call
	Div
	Dup
	Enter
	Eq
	Geq
	Gr
	Halt
	Jump
	JumpZ
	JumpI
	Leq
	Le
	Load
	LoadA
	LoadC
	LoadR
	LoadRC
	Malloc
	Mark
	Mod
	Mul
	Neg
	Neq
	New
	Not
	Or
	Pop
	Return
	Slide
	Store
	StoreA
	StoreR
	Sub
)

var operations = []func(m *Machine){
	doNone,
	doAdd,
	doAnd,
	doAlloc,
	doCall,
	doDiv,
	doDup,
	doEnter,
	doEq,
	doGeq,
	doGr,
	doHalt,
	doJump,
	doJumpZ,
	doJumpI,
	doLeq,
	doLe,
	doLoad,
	doLoadA,
	doLoadC,
	doLoadR,
	doLoadRC,
	doMalloc,
	doMark,
	doMod,
	doMul,
	doNeg,
	doNeq,
	doNew,
	doNot,
	doOr,
	doPop,
	doReturn,
	doSlide,
	doStore,
	doStoreA,
	doStoreR,
	doSub,
}

func doNone(m *Machine) {}

func doAdd(m *Machine) {
	m.binary(func(a, b int) int { return a + b })
}

func doAnd(m *Machine) {
	m.binary(func(a, b int) int { return integer((a != 0) && (b != 0)) })
}

func doAlloc(m *Machine) {}

func doCall(m *Machine) {}

func doDiv(m *Machine) {
	m.binary(func(a, b int) int { return a / b })
}

func doDup(m *Machine) {
	n0 := m.pop()
	m.push(n0)
	m.push(n0)
}

func doEnter(m *Machine) {}

func doEq(m *Machine) {
	m.binary(func(a, b int) int { return integer(a == b) })
}

func doGeq(m *Machine) {
	m.binary(func(a, b int) int { return integer(a >= b) })
}

func doGr(m *Machine) {
	m.binary(func(a, b int) int { return integer(a > b) })
}

func doHalt(m *Machine) {}

func doJump(m *Machine) {
	m.pc = getInteger(m.memory[m.pc : m.pc+4])
}

func doJumpI(m *Machine) {
	m.pc = m.pop() + getInteger(m.memory[m.pc:m.pc+4])
}

func doJumpZ(m *Machine) {
	if n0 := m.pop(); n0 == 0 {
		m.pc = getInteger(m.memory[m.pc : m.pc+4])
	}
}

func doLeq(m *Machine) {
	m.binary(func(a, b int) int { return integer(a <= b) })
}

func doLe(m *Machine) {
	m.binary(func(a, b int) int { return integer(a < b) })
}

func doLoad(m *Machine) {
	addr := m.pop()
	vl := getInteger(m.memory[addr : addr+4])
	m.push(vl)
}

func doLoadA(m *Machine) {}

func doLoadC(m *Machine) {
	nv := getInteger(m.memory[m.pc : m.pc+4])
	m.pc += 4
	m.push(nv)
}

func doLoadR(m *Machine) {}

func doLoadRC(m *Machine) {}

func doMalloc(m *Machine) {}

func doMark(m *Machine) {}

func doMod(m *Machine) {}

func doMul(m *Machine) {
	m.binary(func(a, b int) int { return a * b })
}

func doNeg(m *Machine) {
	n0 := m.pop()
	m.push(-n0)
}

func doNeq(m *Machine) {}

func doNew(m *Machine) {}

func doNot(m *Machine) {
	n0 := m.pop()
	m.push(integer(n0 == 0))
}

func doOr(m *Machine) {
	m.binary(func(a, b int) int { return integer((a != 0) || (b != 0)) })
}

func doPop(m *Machine) {}

func doReturn(m *Machine) {}

func doSlide(m *Machine) {
	addr := m.pop()
	vl := m.pop()
	putInteger(vl, m.memory[addr:addr+4])
}

func doStore(m *Machine) {}

func doStoreA(m *Machine) {}

func doStoreR(m *Machine) {}

func doSub(m *Machine) {
	m.binary(func(a, b int) int { return a - b })
}

func integer(b bool) int {
	if b {
		return 1
	}
	return 0
}
