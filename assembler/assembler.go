package assembler

import (
	"bufio"
	"container/list"
	"os"
)

// Assemble ֆունկցիան src ֆայլում գրված ծրագիրը թարգմանում է
// կատարման համար պատրաստ բինար կոդի և գրում է dest ֆայլում։
func Assemble(src, dst string) {
	srcFile, er := os.Open(src)
	defer srcFile.Close()

	if er == nil {
		pars := new(parser)
		pars.source = bufio.NewReader(srcFile)

		// TEST
		ast, err := pars.parse()
		if err == nil {
			for el := ast.Front(); el != nil; el = el.Next() {
				el.Value.(*instruction).print()
			}
		} else {
			println(err.Error())
		}
	}
}

func assemble(ins *list.List) []byte {
	var code []byte

	return code
}
