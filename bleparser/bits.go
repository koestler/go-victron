package bleparser

// le24 reads a little endian 24 bit unsigned integer from the first three bytes of b.
func le24(b []byte) uint32 {
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16
}

// signExtend interprets the lowest bits bits of v as a two's complement signed integer.
func signExtend(v uint32, bits uint) int32 {
	shift := 32 - bits
	return int32(v<<shift) >> shift
}
