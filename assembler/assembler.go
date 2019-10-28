package assembler

import (
	"bufio"
	"os"
	"unicode"

	"github.com/armenbadal/vmipl/bytecode"
)

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

var source *bufio.Reader
var look lexeme

type parseError struct {
	line    int
	message string
}

func (er *parseError) Error() string {
	return er.message
}

// Assemble ֆունկցիան src ֆայլում գրված ծրագիրը թարգմանում է
// կատարման համար պատրաստ բինար կոդի և գրում է dest ֆայլում։
func Assemble(src string) []byte {
	srcFile, _ := os.Open(src)
	defer srcFile.Close()

	bc := bytecode.NewBuilder()
	source = bufio.NewReader(srcFile)
	if err := parse(bc); err != nil {
		println(err.Error())
	}

	code, _ := bc.Bytes()
	return code
}

// Քերականաություն
//
// Program = { Line EOL{EOL} }.
// Line = Label
//      | Instruction
//      | Label Instruction.
// Label = IDENT ':'.
// Instruction = KEYWORD
//             | KEYWORD INTEGER
//             | KEYWORD IDENT .

func parse(bc *bytecode.Builder) error {
	next()

	for look.token == xNewLine {
		next()
	}

	for look.token == xIdent || look.token == xKeyword {
		if err := parseLine(bc); err != nil {
			return err
		}
	}

	return nil
}

func parseLine(bc *bytecode.Builder) error {
	// պիտակը
	if look.token == xIdent {
		label, _ := match(xIdent)
		if _, err := match(xColon); err != nil {
			return err
		}
		bc.SetLabel(label)
	}

	// // հրահանգը և արգումենտները
	// if look.token == xKeyword {
	// 	nm, _ := match(xKeyword)
	// 	bc.AddByte(byte(indexOf(nm)))

	// 	// առաջին արգումենտ
	// 	if look.token == xNumber {
	// 		lex, _ := match(xNumber)
	// 		nv, _ := strconv.Atoi(lex)
	// 		bc.AddInteger(nv)
	// 	} else if look.token == xIdent {
	// 		label, _ := match(xIdent)
	// 		if addr, marked := addresses[label]; marked {
	// 			bc.AddInteger(addr)
	// 		} else {
	// 			if _, elis := unresolved[label]; !elis {
	// 				unresolved[label] = list.New()
	// 			}
	// 			unresolved[label].PushBack(bc.Size())
	// 			bc.AddInteger(0)
	// 		}
	// 	}

	// 	// երկրորդ արգումենտ
	// 	if look.token == xNumber {
	// 		lex, _ := match(xNumber)
	// 		nv, _ := strconv.Atoi(lex)
	// 		bc.AddInteger(nv)
	// 	}
	// }

	// տող ավարտող 'նոր տող' նիշը
	if _, err := match(xNewLine); err != nil {
		return err
	}
	for look.token == xNewLine {
		match(xNewLine)
	}

	return nil
}

func match(ex int) (string, error) {
	if look.token == ex {
		lex := look.value
		next()
		return lex, nil
	}

	return "", &parseError{0, "Վերլուծության սխալ։"}
}

var keywords = []string{
	"none",
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
	"jumpi",
	"jumpz",
	"leq",
	"le",
	"load",
	"loada",
	"loadc",
	"loadr",
	"loadrc",
	"malloc",
	"mark",
	"mod",
	"mul",
	"neg",
	"neq",
	"new",
	"not",
	"or",
	"pop",
	"return",
	"slide",
	"store",
	"storea",
	"storer",
	"sub"}

func indexOf(kw string) int {
	index := len(keywords) - 1
	for index >= 0 && keywords[index] != kw {
		index--
	}
	return index
}

func isKeyword(si string) bool {
	return indexOf(si) != -1
}

func next() {
	ch := read()

	for ch == ' ' || ch == '\t' || ch == '\r' {
		ch = read()
	}

	switch {
	case ch == 0:
		look = lexeme{xEos, "EOS"}
	case ch == ';':
		for ch != '\n' {
			ch = read()
		}
		source.UnreadRune()
		next()
	case ch == ':':
		look = lexeme{xColon, ":"}
	case ch == '\n':
		look = lexeme{xNewLine, "NL"}
	case unicode.IsDigit(ch) || ch == '-' || ch == '+':
		look = lexeme{xNumber, string(ch)}
		for ch = read(); unicode.IsDigit(ch); ch = read() {
			look.value += string(ch)
		}
		source.UnreadRune()
	case unicode.IsLetter(ch):
		look = lexeme{xIdent, ""}
		for unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			look.value += string(ch)
			ch = read()
		}
		source.UnreadRune()
		if isKeyword(look.value) {
			look.token = xKeyword
		}
	}
}

func read() rune {
	ch, _, er := source.ReadRune()
	if er != nil {
		return 0
	}
	return ch
}
