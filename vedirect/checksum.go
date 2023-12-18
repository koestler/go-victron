package vedirect

func computeChecksum(cmd byte, data []byte) (checksum byte) {
	checksum = byte(0x55)
	checksum -= cmd
	for _, v := range data {
		checksum -= v
	}
	return
}
