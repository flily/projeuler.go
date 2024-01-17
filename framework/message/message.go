package message

import (
	"fmt"
	"time"
)

type MessageType byte

const (
	// Message types
	MessageType_Unknown MessageType = 0
	MessageType_Invalid MessageType = 1
	MessageType_Ping    MessageType = 2
	MessageType_Pong    MessageType = 3
	MessageType_Run     MessageType = 4
	MessageType_Result  MessageType = 5
)

const (
	// Message flags
	MessageFlag_Finished = (1 << iota)
	MessageFlag_Timeout
	MessageFlag_Error
)

// MessageHeader is common header for all messages.
// Message header is fixed 4 bytes. The first byte is command type, the following 3 bytes is total
// length of message, in big endian, with maxium message length 16M.
// +-----+-----+-----+-----+
// |  0  |  1  |  2  |  3  |
// +-----+-----+-----+-----+
// | CMD |  total length   |
// +-----+-----+-----+-----+
type MessageHeader struct {
	Command     MessageType
	TotalLength int
}

func (h *MessageHeader) MessageLength() int {
	return 4
}

func (h *MessageHeader) SerializeTo(buffer []byte, offset int) (int, error) {
	if offset+4 > len(buffer) {
		return 0, ErrBufferTooSmall
	}

	buffer[offset+0] = byte(h.Command)
	writeUint24(buffer, offset+1, h.TotalLength)
	return 4, nil
}

func (h *MessageHeader) DeserializeFrom(buffer []byte, offset int) (int, error) {
	if offset+4 > len(buffer) {
		return 0, ErrBufferTooSmall
	}

	h.Command = MessageType(buffer[offset+0])
	h.TotalLength, _ = readUint24(buffer, offset+1)
	return 4, nil
}

type MessagePing struct {
	MessageHeader

	Sequence uint32
}

func NewPingMessage(sequence uint32) *MessagePing {
	return &MessagePing{
		MessageHeader: MessageHeader{
			Command:     MessageType_Ping,
			TotalLength: 8,
		},
		Sequence: sequence,
	}
}

func (m *MessagePing) MessageLength() int {
	return m.MessageHeader.MessageLength() + 4
}

func (m *MessagePing) SerializeTo(buffer []byte, offset int) (int, error) {
	length := m.MessageLength()
	if offset+length > len(buffer) {
		return 0, ErrBufferTooSmall
	}

	m.MessageHeader.TotalLength = length
	headerLength, _ := m.MessageHeader.SerializeTo(buffer, offset)
	writeUint32(buffer, offset+headerLength, m.Sequence)

	return length, nil
}

func (m *MessagePing) Serialize() ([]byte, error) {
	length := m.MessageLength()
	buffer := make([]byte, length)
	_, _ = m.SerializeTo(buffer, 0)
	return buffer, nil
}

func (m *MessagePing) DeserializeFrom(buffer []byte, offset int) (int, error) {
	if offset+8 > len(buffer) {
		return 0, ErrBufferTooSmall
	}

	headerLength, _ := m.MessageHeader.DeserializeFrom(buffer, offset)
	m.Sequence, _ = readUint32(buffer, offset+headerLength)

	return 8, nil
}

func (m *MessagePing) MakePong() *MessagePing {
	return &MessagePing{
		MessageHeader: MessageHeader{
			Command: MessageType_Pong,
		},
		Sequence: 0xffffffff ^ m.Sequence,
	}
}

// MessageRun presents a message to run a problem.
// +-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+-----+
// |  Message Header (4B)  |  Problem ID (uint32)  |     Method (VLSS)     |
// +-----------------------+-----------------------+-----------------------+
type MessageRun struct {
	MessageHeader

	Problem int
	Method  string
}

func NewRunMessage(problem int, method string) *MessageRun {
	m := &MessageRun{
		MessageHeader: MessageHeader{
			Command:     MessageType_Run,
			TotalLength: 8 + len(method) + 1,
		},
		Problem: problem,
		Method:  method,
	}

	return m
}

func DeserializeRunMessage(buffer []byte, offset int) (*MessageRun, error) {
	message := &MessageRun{}
	if _, err := message.DeserializeFrom(buffer, offset); err != nil {
		return nil, err
	}

	return message, nil
}

func (m *MessageRun) MessageLength() int {
	length := m.MessageHeader.MessageLength()
	length += 4 + len(m.Method) + 1
	return length
}

func (m *MessageRun) SerializeTo(buffer []byte, offset int) (int, error) {
	length := m.MessageLength()
	if offset+length > len(buffer) {
		return 0, ErrBufferTooSmall
	}

	m.MessageHeader.TotalLength = length
	headerLength, _ := m.MessageHeader.SerializeTo(buffer, offset)
	writeUint32(buffer, offset+headerLength+0, uint32(m.Problem))
	writeShortString(buffer, offset+headerLength+4, m.Method)

	return headerLength + 8 + len(m.Method), nil
}

func (m *MessageRun) Serialize() ([]byte, error) {
	length := m.MessageLength()
	buffer := make([]byte, length)
	_, _ = m.SerializeTo(buffer, 0)
	return buffer, nil
}

func (m *MessageRun) DeserializeFrom(buffer []byte, offset int) (int, error) {
	headerLength, err := m.MessageHeader.DeserializeFrom(buffer, offset)
	if err != nil {
		return 0, err
	}

	if m.MessageHeader.Command != MessageType_Run {
		return 0, fmt.Errorf("message is not RunMessage, got '%d'", m.MessageHeader.Command)
	}

	var length int
	problem, _ := readUint32(buffer, offset+headerLength+0)
	m.Problem = int(problem)
	if m.Method, length = readShortString(buffer, offset+headerLength+4); length < 0 {
		return 0, ErrBufferTooSmall
	}

	return headerLength + 8 + len(m.Method), nil
}

// MessageResultItem presents a message to return result of a method.
// +-----------------------+-----------------------+-----------------------+
// |  Message Header (4B)  |  Problem ID (uint32)  |     Method (VLSS)     |
// +-----------------------+-----------------------+-----------------------+
// |  Flags Mask (uint32)  |     Result (int64)    |    Duration (int64)   |
// +-----------------------+-----------------------+-----------------------+
type MessageResultItem struct {
	ProblemId  int
	Method     string
	Result     int64
	Duration   time.Duration
	IsTimeout  bool
	IsFinished bool
}

// MessageResult presents a message to return result of a run request.
// +-----------------------+-----------------------+-----------------------+
// |  Message Header (4B)  | Result count (uint32) |      Result Item      |
// +-----------------------+-----------------------+-----------------------+
type MessageResult struct {
	MessageHeader

	ResultCount int
	Results     []MessageResultItem
}

func DeserializeHeader(buffer []byte, offset int) (*MessageHeader, error) {
	header := &MessageHeader{}
	if _, err := header.DeserializeFrom(buffer, offset); err != nil {
		return nil, err
	}

	return header, nil
}
