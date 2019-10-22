package assembler

import (
	"os"

	"github.com/armenbadal/vmipl/bytecode"
)

// Assemble ֆունկցիան src ֆայլում գրված ծրագիրը թարգմանում է
// կատարման համար պատրաստ բինար կոդի և գրում է dest ֆայլում։
func Assemble(src string) *bytecode.ByteCode {
	srcFile, _ := os.Open(src)
	defer srcFile.Close()

	return nil
}
