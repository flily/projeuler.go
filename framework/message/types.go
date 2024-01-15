package message

func readUint24(buffer []byte, offset int) (int, int) {
	value := (int(buffer[offset+0]) << 16) |
		(int(buffer[offset+1]) << 8) |
		(int(buffer[offset+2]) << 0)
	return value, 3
}

func writeUint24(buffer []byte, offset int, value int) int {
	buffer[offset+0] = byte((value >> 16) & 0xff)
	buffer[offset+1] = byte((value >> 8) & 0xff)
	buffer[offset+2] = byte((value >> 0) & 0xff)

	return 3
}

func readUint32(buffer []byte, offset int) (uint32, int) {
	value := (uint32(buffer[offset+0]) << 24) |
		(uint32(buffer[offset+1]) << 16) |
		(uint32(buffer[offset+2]) << 8) |
		(uint32(buffer[offset+3]) << 0)

	return value, 4
}

func writeUint32(buffer []byte, offset int, value uint32) int {
	buffer[offset+0] = byte((value >> 24) & 0xff)
	buffer[offset+1] = byte((value >> 16) & 0xff)
	buffer[offset+2] = byte((value >> 8) & 0xff)
	buffer[offset+3] = byte((value >> 0) & 0xff)

	return 4
}

func readInt64(buffer []byte, offset int) (int64, int) {
	value := (int64(buffer[offset+0]) << 56) |
		(int64(buffer[offset+1]) << 48) |
		(int64(buffer[offset+2]) << 40) |
		(int64(buffer[offset+3]) << 32) |
		(int64(buffer[offset+4]) << 24) |
		(int64(buffer[offset+5]) << 16) |
		(int64(buffer[offset+6]) << 8) |
		(int64(buffer[offset+7]) << 0)

	return value, 8
}

func writeInt64(buffer []byte, offset int, value int64) int {
	buffer[offset+0] = byte((value >> 56) & 0xff)
	buffer[offset+1] = byte((value >> 48) & 0xff)
	buffer[offset+2] = byte((value >> 40) & 0xff)
	buffer[offset+3] = byte((value >> 32) & 0xff)
	buffer[offset+4] = byte((value >> 24) & 0xff)
	buffer[offset+5] = byte((value >> 16) & 0xff)
	buffer[offset+6] = byte((value >> 8) & 0xff)
	buffer[offset+7] = byte((value >> 0) & 0xff)

	return 8
}

func readShortString(buffer []byte, offset int) (string, int) {
	if offset >= len(buffer) {
		return "", -1
	}

	length := int(buffer[offset])
	if offset+1+length > len(buffer) {
		return "", -1
	}

	result := string(buffer[offset+1 : offset+1+length])
	return result, length + 1
}

func writeShortString(buffer []byte, offset int, value string) int {
	length := len(value)
	if length > 255 || offset+1+length > len(buffer) {
		return -1
	}

	buffer[offset] = byte(length)
	copy(buffer[offset+1:], value)
	return length + 1
}
