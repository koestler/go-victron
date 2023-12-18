package vedirect

// ComputeChecksum computes the checksum for a given command and data.
// see https://www.victronenergy.com/live/vedirect_protocol:faq
func ComputeChecksum(cmd byte, data []byte) (checksum byte) {
	checksum = byte(0x55)
	checksum -= cmd
	for _, v := range data {
		checksum -= v
	}
	return
}
