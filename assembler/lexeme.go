package assembler

const (
	xNone = iota

	xNumber
	xIdent

	xKeyword
	// xAdd
	// xAnd
	// xAlloc
	// xCall
	// xDiv
	// xDup
	// xEnter
	// xEq
	// xGeq
	// xGr
	// xHalt
	// xJump
	// xJumpz
	// xJumpi
	// xLeq
	// xLe
	// xLoad
	// xLoada
	// xLoadc
	// xLoadr
	// xLoadrc
	// xMalloc
	// xMark
	// xMul
	// xNeg
	// xNeq
	// xNew
	// xOr
	// xPop
	// xReturn
	// xSlide
	// xStore
	// xStorea
	// xStorer
	// xSub

	xNewLine
	xColon

	xEos
)

type lexeme struct {
	token int
	value string
}
