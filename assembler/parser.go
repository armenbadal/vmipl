package assembler

import (
	"bufio"
	"container/list"
	"sort"
	"strconv"
	"unicode"
)

type parser struct {
	source *bufio.Reader
	look   lexeme
}

type parseError struct {
	line    int
	message string
}

func (er *parseError) Error() string {
	return er.message
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

func (p *parser) parse() (*list.List, error) {
	p.next()

	for p.look.token == xNewLine {
		p.next()
	}

	ast := list.New()
	for p.look.token == xIdent || p.look.token == xKeyword {
		el, err := p.parseLine()
		if err != nil {
			return nil, err
		}
		ast.PushBack(el)
	}

	return ast, nil
}

func (p *parser) parseLine() (*instruction, error) {
	instr := new(instruction)

	// պիտակը
	if p.look.token == xIdent {
		instr.label, _ = p.match(xIdent)
		_, err := p.match(xColon)
		if err != nil {
			return nil, err
		}
	}

	// հրահանգը և արգումենտները
	if p.look.token == xKeyword {
		instr.name, _ = p.match(xKeyword)

		// առաջին արգումենտ
		if p.look.token == xNumber {
			nv, _ := p.match(xNumber)
			instr.number0, _ = strconv.Atoi(nv)
		} else if p.look.token == xIdent {
			instr.symbol, _ = p.match(xIdent)
		}

		// երկրորդ արգումենտ
		if p.look.token == xNumber {
			nv, _ := p.match(xNumber)
			instr.number1, _ = strconv.Atoi(nv)
		}
	}

	// տող ավարտող 'նոր տող' նիշը
	_, err := p.match(xNewLine)
	if err != nil {
		return nil, err
	}
	for p.look.token == xNewLine {
		p.match(xNewLine)
	}

	return instr, nil
}

func (p *parser) match(ex int) (string, error) {
	if p.look.token == ex {
		lex := p.look.value
		p.next()
		return lex, nil
	}

	return "", &parseError{0, "Վերլուծության սխալ։"}
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
	ch := p.read()

	for ch == ' ' || ch == '\t' || ch == '\r' {
		ch = p.read()
	}

	switch {
	case ch == 0:
		p.look = lexeme{xEos, "EOS"}
	case ch == ';':
		for ch != '\n' {
			ch = p.read()
		}
		p.source.UnreadRune()
		p.next()
	case ch == ':':
		p.look = lexeme{xColon, ":"}
	case ch == '\n':
		p.look = lexeme{xNewLine, "NL"}
	case unicode.IsDigit(ch) || ch == '-' || ch == '+':
		p.look = lexeme{xNumber, string(ch)}
		for ch = p.read(); unicode.IsDigit(ch); ch = p.read() {
			p.look.value += string(ch)
		}
		p.source.UnreadRune()
	case unicode.IsLetter(ch):
		p.look = lexeme{xIdent, ""}
		for unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			p.look.value += string(ch)
			ch = p.read()
		}
		p.source.UnreadRune()
		if isKeyword(p.look.value) {
			p.look.token = xKeyword
		}
	}
}

func (p *parser) read() rune {
	ch, _, er := p.source.ReadRune()
	if er != nil {
		return 0
	}
	return ch
}
