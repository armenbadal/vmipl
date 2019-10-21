package machine

func putInteger(n int, p []byte) {
	_ = p[3]
	p[0] = byte(n)
	p[1] = byte(n >> 8)
	p[2] = byte(n >> 16)
	p[3] = byte(n >> 24)
}

func getInteger(p []byte) int {
	_ = p[3]
	return int(p[0]) | int(p[1])<<8 | int(p[2])<<16 | int(p[3])<<24
}
