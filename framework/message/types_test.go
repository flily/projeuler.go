package message

import (
	"testing"

	"bytes"
)

func TestUint23Operations(t *testing.T) {
	writeBuffer := make([]byte, 4)
	gotLength := writeUint24(writeBuffer, 0, 0x1a2b3c)
	if gotLength != 3 {
		t.Errorf("expected 3, got %d", gotLength)
	}

	expected := []byte{0x1a, 0x2b, 0x3c}
	if !bytes.Equal(writeBuffer[:gotLength], expected) {
		t.Errorf("expected %v, got %v", expected, writeBuffer[:gotLength])
	}

	readValue, readLength := readUint24(writeBuffer, 0)
	if readLength != 3 {
		t.Errorf("expected 3, got %d", readLength)
	}

	if readValue != 0x1a2b3c {
		t.Errorf("expected 0x1a2b3c, got 0x%x", readValue)
	}
}

func TestUint32Operations(t *testing.T) {
	writeBuffer := make([]byte, 4)
	gotLength := writeUint32(writeBuffer, 0, 0x1a2b3c4d)
	if gotLength != 4 {
		t.Errorf("expected 4, got %d", gotLength)
	}

	expected := []byte{0x1a, 0x2b, 0x3c, 0x4d}
	if !bytes.Equal(writeBuffer[:gotLength], expected) {
		t.Errorf("expected %v, got %v", expected, writeBuffer[:gotLength])
	}

	readValue, readLength := readUint32(writeBuffer, 0)
	if readLength != 4 {
		t.Errorf("expected 4, got %d", readLength)
	}

	if readValue != 0x1a2b3c4d {
		t.Errorf("expected 0x1a2b3c4d, got 0x%x", readValue)
	}
}

func TestInt64OperationsOnPositiveNumber(t *testing.T) {
	writeBuffer := make([]byte, 8)
	gotLength := writeInt64(writeBuffer, 0, 0x1a2b3c4d5e6f7a8b)
	if gotLength != 8 {
		t.Errorf("expected 8, got %d", gotLength)
	}

	expected := []byte{0x1a, 0x2b, 0x3c, 0x4d, 0x5e, 0x6f, 0x7a, 0x8b}
	if !bytes.Equal(writeBuffer[:gotLength], expected) {
		t.Errorf("expected 0x%x, got 0x%x", expected, writeBuffer[:gotLength])
	}

	readValue, readLength := readInt64(writeBuffer, 0)
	if readLength != 8 {
		t.Errorf("expected 8, got %d", readLength)
	}

	if readValue != 0x1a2b3c4d5e6f7a8b {
		t.Errorf("expected 0x1a2b3c4d5e6f7a8b, got 0x%x", readValue)
	}
}

func TestInt64OperationsOnNegativeNumber(t *testing.T) {
	writeBuffer := make([]byte, 8)
	n := -int64(0x1a2b3c4d5e6f7a8b)
	gotLength := writeInt64(writeBuffer, 0, n)
	if gotLength != 8 {
		t.Errorf("expected 8, got %d", gotLength)
	}

	expectedInt := uint64(0xffffffffffffffff) - uint64(-n) + 1
	expected := []byte{0xe5, 0xd4, 0xc3, 0xb2, 0xa1, 0x90, 0x85, 0x75}
	if !bytes.Equal(writeBuffer[:gotLength], expected) {
		t.Errorf("expected %x, got %x", expected, writeBuffer[:gotLength])
	}

	readValue, readLength := readInt64(writeBuffer, 0)
	if readLength != 8 {
		t.Errorf("expected 8, got %d", readLength)
	}

	if readValue != n {
		t.Errorf("expected -0x1a2b3c4d5e6f7a8b, got 0x%x", readValue)
	}

	if readValue != int64(expectedInt) {
		t.Errorf("expected 0x%x, got 0x%x", expectedInt, readValue)
	}
}

func TestShortStringOperations(t *testing.T) {
	s := "hello world"

	writtenStr := make([]byte, 16)
	gotLength := writeShortString(writtenStr, 0, s)
	if gotLength != len(s)+1 {
		t.Errorf("expected %d, got %d", len(s)+1, gotLength)
	}

	expected := []byte{0x0b, 'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'}
	if !bytes.Equal(writtenStr[:gotLength], expected) {
		t.Errorf("expected %v, got %v", expected, writtenStr[:gotLength])
	}

	readString, readLength := readShortString(writtenStr, 0)
	if readLength != len(s)+1 {
		t.Errorf("expected %d, got %d", len(s)+1, readLength)
	}

	if readString != s {
		t.Errorf("expected %s, got %s", s, readString)
	}
}

func TestShortStringOperationsSmallBuffer(t *testing.T) {
	s := "hello world"

	writeBuffer := make([]byte, 16)
	if length := writeShortString(writeBuffer, 16, s); length != -1 {
		t.Errorf("expected -1, got %d", length)
	}

	writeShortString(writeBuffer, 0, s)
	if result, length := readShortString(writeBuffer, 16); result != "" || length != -1 {
		t.Errorf("expected ('', -1), got (%s, %d)", result, length)
	}

	if result, length := readShortString(writeBuffer[:12], 10); result != "" || length != -1 {
		t.Errorf("expected ('', -1), got (%s, %d)", result, length)
	}
}
