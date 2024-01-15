package message

func writeInt(buffer []byte, offset int, size_in_byte int, value int64) int {
	for i := 0; i < size_in_byte; i++ {
		o := (size_in_byte - i - 1) * 8
		buffer[offset+i] = byte((value >> o) & 0xff)
	}

	return size_in_byte
}

func readInt(buffer []byte, offset int, size_in_byte int) (int64, int) {
	value := int64(0)
	for i := 0; i < size_in_byte; i++ {
		value = value << 8
		value = value | int64(buffer[offset+i])
	}

	return value, size_in_byte
}

func writeUint(buffer []byte, offset int, size_in_byte int, value uint64) int {
	for i := 0; i < size_in_byte; i++ {
		o := (size_in_byte - i - 1) * 8
		buffer[offset+i] = byte((value >> o) & 0xff)
	}

	return size_in_byte
}

func readUint(buffer []byte, offset int, size_in_byte int) (uint64, int) {
	value := uint64(0)
	for i := 0; i < size_in_byte; i++ {
		value = value << 8
		value = value | uint64(buffer[offset+i])
	}

	return value, size_in_byte
}

func readUint24(buffer []byte, offset int) (int, int) {
	value, shift := readUint(buffer, offset, 3)
	return int(value), shift
}

func writeUint24(buffer []byte, offset int, value int) int {
	return writeUint(buffer, offset, 3, uint64(value))
}

func readUint32(buffer []byte, offset int) (int, int) {
	value, shift := readUint(buffer, offset, 4)
	return int(value), shift
}

func writeUint32(buffer []byte, offset int, value uint32) int {
	return writeUint(buffer, offset, 4, uint64(value))
}

func readInt64(buffer []byte, offset int) (int64, int) {
	value, shift := readInt(buffer, offset, 8)
	return value, shift
}

func writeInt64(buffer []byte, offset int, value int64) int {
	return writeInt(buffer, offset, 8, value)
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
