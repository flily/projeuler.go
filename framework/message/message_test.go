package message

import (
	"testing"
	"time"

	"bytes"
	"errors"
)

func TestMessageHeaderSerialize(t *testing.T) {
	header := &MessageHeader{
		Command:     MessageType_Run,
		TotalLength: 0x1a2b3c,
	}

	got := make([]byte, 4)
	if _, err := header.SerializeTo(got, 0); err != nil {
		t.Errorf("serialize failed: %v", err)
	}

	expected := []byte{0x04, 0x1a, 0x2b, 0x3c}
	if !bytes.Equal(got, expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}

	newHeader, err := DeserializeHeader(got, 0)
	if err != nil {
		t.Errorf("deserialize failed: %v", err)
	}

	if *newHeader != *header {
		t.Errorf("expected %v, got %v", header, newHeader)
	}
}

func TestMessageHeaderSerializeSmallBuffer(t *testing.T) {
	header := &MessageHeader{
		Command:     MessageType_Run,
		TotalLength: 0x1a2b3c,
	}

	got := make([]byte, 3)
	if _, err := header.SerializeTo(got, 0); err == nil {
		t.Errorf("serialize should fail")
	}
}

func TestMessageHeaderDeserializeSmallBuffer(t *testing.T) {
	buffer := []byte{0x01, 0x1a, 0x2b}

	if _, err := DeserializeHeader(buffer, 0); err == nil {
		t.Errorf("deserialize should fail")
	}
}

func TestMessagePingSerialize(t *testing.T) {
	message := NewPingMessage(0x1a2b3c4d)
	if message.Command != MessageType_Ping {
		t.Errorf("expected MessageType_Ping, got %v", message.Command)
	}

	expected := []byte{
		byte(MessageType_Ping), 0x00, 0x00, 0x08,
		0x1a, 0x2b, 0x3c, 0x4d,
	}

	got, err := message.Serialize()
	if err != nil {
		t.Errorf("serialize failed: %v", err)
	}

	if !bytes.Equal(got, expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}

	newMessage := &MessagePing{}
	if _, err := newMessage.DeserializeFrom(got, 0); err != nil {
		t.Errorf("deserialize failed: %v", err)
	}

	if *newMessage != *message {
		t.Errorf("expected %v, got %v", message, newMessage)
	}
}

func TestMessagePingSerializeToSmallBuffer(t *testing.T) {
	message := NewPingMessage(0x1a2b3c4d)
	got := make([]byte, 3)
	if _, err := message.SerializeTo(got, 0); err == nil {
		t.Errorf("serialize should fail")
	}

	newMessage := &MessagePing{}
	if _, err := newMessage.DeserializeFrom(got, 0); err == nil {
		t.Errorf("deserialize should fail")
	}
}

func TestMessagePingAndPong(t *testing.T) {
	message := NewPingMessage(0xcacacaca)
	pong := message.MakePong()
	if pong.Command != MessageType_Pong {
		t.Errorf("expected MessageType_Pong, got %v", pong.Command)
	}

	expected := uint32(0x35353535)
	if pong.Sequence != expected {
		t.Errorf("expected 0x%x, got 0x%x", expected, pong.Sequence)
	}
}

func TestMessageRunSerialize(t *testing.T) {
	message := &MessageRun{
		MessageHeader: MessageHeader{
			Command: MessageType_Run,
		},
		ProblemTimeout: 5 * time.Second,
		MethodTimeout:  3 * time.Second,
		Problem:        0x1a2b3c4d,
		Method:         "lorem",
	}

	expected := []byte{
		0x04, 0x00, 0x00, 0x1e, // header
		0x00, 0x00, 0x00, 0x01, 0x2a, 0x05, 0xf2, 0x00, // problem timeout
		0x00, 0x00, 0x00, 0x00, 0xb2, 0xd0, 0x5e, 0x00, // method timeout
		0x1a, 0x2b, 0x3c, 0x4d, // problem
		0x05, 0x6c, 0x6f, 0x72, 0x65, 0x6d, // method
	}
	got, err := message.Serialize()
	if err != nil {
		t.Errorf("serialize failed: %v", err)
	}

	if !bytes.Equal(got, expected) {
		t.Errorf("serialize result error.\nexpected %v\n     got %v", expected, got)
	}

	// TotalLength is automatically calculated when Serialize() is called.
	createdMessage := NewRunMessage(0x1a2b3c4d, "lorem")
	createdMessage.SetTimeout(5*time.Second, 3*time.Second)
	if *message != *createdMessage {
		t.Errorf("created wrong message struct: %+v", createdMessage)
	}

	newMessage, err := DeserializeRunMessage(got, 0)
	if err != nil {
		t.Errorf("deserialize failed: %v", err)
	}

	if *newMessage != *message {
		t.Errorf("expected %v, got %v", message, newMessage)
	}
}

func TestMessageRunSerializeToSmallBuffer(t *testing.T) {
	message := &MessageRun{
		MessageHeader: MessageHeader{
			Command: MessageType_Run,
		},
		Problem: 0x1a2b3c4d,
		Method:  "lorem",
	}

	buf2 := make([]byte, 2)
	buf6 := make([]byte, 6)
	_, err := message.SerializeTo(buf2, 0)
	if err == nil {
		t.Errorf("serialize should fail")
	}

	if !errors.Is(err, ErrBufferTooSmall) {
		t.Errorf("expected ErrBufferTooSmall, got %v", err)
	}

	_, err = DeserializeRunMessage(buf2, 0)
	if err == nil {
		t.Errorf("deserialize should fail")
	}

	_, err = DeserializeRunMessage(buf6, 0)
	if err == nil {
		t.Errorf("deserialize should fail")
	}
}
func TestMessageRunDeserializeWithInsuffientBufferForTotalMessage(t *testing.T) {
	data := []byte{
		0x04, 0x00, 0x00, 0x1e, // header
		0x00, 0x00, 0x00, 0x01, 0x2a, 0x05, 0xf2, 0x00, // problem timeout
		0x00, 0x00, 0x00, 0x00, 0xb2, 0xd0, 0x5e, 0x00, // method timeout
		0x1a, 0x2b, 0x3c, 0x4d, // problem
		0x08, 0x6c, 0x6f, 0x72, 0x65, 0x6d, // method
	}

	message, err := DeserializeRunMessage(data, 0)
	if err == nil {
		t.Errorf("deserialize should fail")
	}

	if message != nil {
		t.Errorf("expected nil, got %v", message)
	}

	if !errors.Is(err, ErrBufferTooSmall) {
		t.Errorf("expected ErrBufferTooSmall, got %v", err)
	}
}

func TestMessageRunDeserializeWithInsuffientBufferForMethod(t *testing.T) {
	data := []byte{
		0x04, 0x00, 0x00, 0x0e, // header
		0x1a, 0x2b, 0x3c, 0x4d, // problem
		0x05, 0x6c, 0x6f, // method
	}

	message, err := DeserializeRunMessage(data, 0)
	if err == nil {
		t.Errorf("deserialize should fail")
	}

	if message != nil {
		t.Errorf("expected nil, got %v", message)
	}

	if !errors.Is(err, ErrBufferTooSmall) {
		t.Errorf("expected ErrBufferTooSmall, got %v", err)
	}
}

func TestMessageRunDeserializeWithWrongMessage(t *testing.T) {
	data := []byte{
		0x01, 0x00, 0x00, 0x0e, // header
		0x1a, 0x2b, 0x3c, 0x4d, // problem
		0x05, 0x6c, 0x6f, // method
	}

	message, err := DeserializeRunMessage(data, 0)
	if err == nil {
		t.Errorf("deserialize should fail")
	}

	if message != nil {
		t.Errorf("expected nil, got %v", message)
	}
}
