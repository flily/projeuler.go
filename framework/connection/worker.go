package connection

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/flily/projeuler.go/framework/message"
)

type WorkerConn struct {
	port       int
	listener   net.Listener
	sendQueue  chan *message.MessageResult
	recvQueue  chan *message.MessageRun
	stopSignal chan struct{}
}

func NewWorkerConn(host string, port int) (*WorkerConn, error) {
	address := fmt.Sprintf("%s:%d", host, port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	w := &WorkerConn{
		port:       port,
		listener:   listener,
		sendQueue:  make(chan *message.MessageResult),
		recvQueue:  make(chan *message.MessageRun),
		stopSignal: make(chan struct{}),
	}

	return w, nil
}

func (w *WorkerConn) Close() {
	w.listener.Close()
	close(w.sendQueue)
	close(w.recvQueue)
}

func (w *WorkerConn) RunLoop() error {
	buffer := make([]byte, 16*1024)
	for {
		conn, err := w.listener.Accept()
		if err != nil {
			return err
		}

		for {
			readLength, err := conn.Read(buffer)
			if err != nil {
				if !errors.Is(err, io.EOF) {
					log.Printf("ERROR on read: %s", err)
				}

				break
			}

			request, err := message.DeserializeRunMessage(buffer[:readLength], 0)
			if err != nil {
				log.Printf("ERROR on deserialize: %s", err)
				break
			}

			w.recvQueue <- request

			result := <-w.sendQueue
			packet, _ := result.Serialize()

			_, err = conn.Write(packet)
			if err != nil {
				log.Printf("ERROR on write: %s", err)
				break
			}
		}
	}
}

func (w *WorkerConn) RecvRun() <-chan *message.MessageRun {
	return w.recvQueue
}

func (w *WorkerConn) SendResult(result *message.MessageResult) {
	w.sendQueue <- result
}
