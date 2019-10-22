package assembler

const (
	xNone = iota

	xNumber
	xIdent

	xKeyword

	xNewLine
	xColon

	xEos
)

type lexeme struct {
	token int
	value string
}
