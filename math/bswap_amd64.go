package math

func bswap32(x uint32) uint32 {
	return ((x << 24) & 0xff000000) |
		((x << 8) & 0x00ff0000) |
		((x >> 8) & 0x0000ff00) |
		((x >> 24) & 0x000000ff)
}

//TODO write in asm or move to separate file
func bswap64(x uint64) uint64 {
	return ((x << 56) & 0xff00000000000000) |
		((x << 40) & 0x00ff000000000000) |
		((x << 24) & 0x0000ff0000000000) |
		((x << 8) & 0x000000ff00000000) |
		((x >> 8) & 0x00000000ff000000) |
		((x >> 24) & 0x0000000000ff0000) |
		((x >> 40) & 0x000000000000ff00) |
		((x >> 56) & 0x00000000000000ff)
}
