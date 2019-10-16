package assembler

import (
	"bufio"
	"fmt"
	"sort"
	"unicode"
)

type parser struct {
	source *bufio.Reader
	ch     rune
	look   lexeme
}

type instruction struct {
	label  string
	name   string
	number int
	symbol string
}

func (ns instruction) print() {
	fmt.Printf("'%s'\t'%s'\t'%d'\t'%s'\n", ns.label, ns.name, ns.number, ns.symbol)
}

// Քերկանաություն
//
// Program = { Line EOL{EOL} }.
// Line = Label
//      | Instruction
//      | Label Instruction.
// Label = IDENT ':'.
// Instruction = KEYWORD
//             | KEYWORD INTEGER
//             | KEYWORD IDENT .

func (p *parser) parse() bool {
	p.read()
	p.next()

	for p.look.token == xNewLine {
		p.next()
	}

	for p.look.token == xIdent || p.look.token == xKeyword {
		p.parseLine().print()
	}

	return true
}

func (p *parser) parseLine() instruction {
	var instr instruction

	if p.look.token == xIdent {
		instr.label = p.look.value
		p.next()
		p.next()
	}

	if p.look.token == xKeyword {
		instr.name = p.look.value
		p.next()
		if p.look.token == xNumber {
			instr.number = 0
			p.next()
		} else if p.look.token == xIdent {
			instr.symbol = p.look.value
			p.next()
		}
	}

	for p.look.token == xNewLine {
		p.next()
	}

	return instr
}

// var keywords = map[string]int{
// 	"add":    xAdd,
// 	"and":    xAnd,
// 	"alloc":  xAlloc,
// 	"call":   xCall,
// 	"div":    xDiv,
// 	"dup":    xDup,
// 	"enter":  xEnter,
// 	"eq":     xEq,
// 	"geq":    xGeq,
// 	"gr":     xGr,
// 	"halt":   xHalt,
// 	"jump":   xJump,
// 	"jumpz":  xJumpz,
// 	"jumpi":  xJumpi,
// 	"leq":    xLeq,
// 	"le":     xLe,
// 	"load":   xLoad,
// 	"loada":  xLoada,
// 	"loadc":  xLoadc,
// 	"loadr":  xLoadr,
// 	"loadrc": xLoadrc,
// 	"malloc": xMalloc,
// 	"mark":   xMark,
// 	"mul":    xMul,
// 	"neg":    xNeg,
// 	"neq":    xNeq,
// 	"new":    xNew,
// 	"or":     xOr,
// 	"pop":    xPop,
// 	"return": xReturn,
// 	"slide":  xSlide,
// 	"store":  xStore,
// 	"storea": xStorea,
// 	"storer": xStorer,
// 	"sub":    xSub,
// }

var keywords = []string{
	"add",
	"and",
	"alloc",
	"call",
	"div",
	"dup",
	"enter",
	"eq",
	"geq",
	"gr",
	"halt",
	"jump",
	"jumpz",
	"jumpi",
	"leq",
	"le",
	"load",
	"loada",
	"loadc",
	"loadr",
	"loadrc",
	"malloc",
	"mark",
	"mul",
	"neg",
	"neq",
	"new",
	"or",
	"pop",
	"return",
	"slide",
	"store",
	"storea",
	"storer",
	"sub"}

func init() {
	sort.Strings(keywords)
}

func isKeyword(si string) bool {
	k := sort.SearchStrings(keywords, si)
	return keywords[k] == si
}

func (p *parser) next() {
	for p.ch == ' ' || p.ch == '\t' || p.ch == '\r' {
		p.read()
	}

	switch {
	case p.ch == 0:
		p.look = lexeme{xEos, "EOS"}
	case p.ch == ';':
		for p.ch != '\n' {
			p.read()
		}
		p.next()
	case p.ch == ':':
		p.look = lexeme{xColon, ":"}
		p.read()
	case p.ch == '\n':
		p.look = lexeme{xNewLine, "NL"}
		p.read()

	case unicode.IsDigit(p.ch):
		p.look = lexeme{xNumber, ""}
		for unicode.IsDigit(p.ch) {
			p.look.value += string(p.ch)
			p.read()
		}
	case unicode.IsLetter(p.ch):
		p.look = lexeme{xIdent, ""}
		for unicode.IsLetter(p.ch) || unicode.IsDigit(p.ch) {
			p.look.value += string(p.ch)
			p.read()
		}
		if isKeyword(p.look.value) {
			p.look.token = xKeyword
		}
	}
}

func (p *parser) read() {
	ch, _, er := p.source.ReadRune()
	if er != nil {
		p.ch = 0
	} else {
		p.ch = ch
	}
}
