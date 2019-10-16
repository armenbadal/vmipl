package assembler

import (
	"bufio"
	"os"
)

// Assemble ֆունկցիան src ֆայլում գրված ծրագիրը թարգմանում է
// կատարման համար պատրաստ բինար կոդի և գրում է dest ֆայլում։
func Assemble(src, dst string) {
	srcFile, er := os.Open(src)
	if er == nil {
		pars := new(parser)
		pars.source = bufio.NewReader(srcFile)
		pars.parse()
		srcFile.Close()
	}
}
