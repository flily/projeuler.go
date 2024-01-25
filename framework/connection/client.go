package connection

import (
	"fmt"
	"log"
	"net"

	"github.com/flily/projeuler.go/framework/message"
)

type Client struct {
	conn net.Conn
}

func NewClient(host string, port int) (*Client, error) {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	c := &Client{
		conn: conn,
	}

	return c, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) Run(request *message.MessageRun) (*message.MessageResult, error) {
	packet, err := request.Serialize()
	if err != nil {
		return nil, err
	}

	_, err = c.conn.Write(packet)
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, 16*1024)
	n, err := c.conn.Read(buffer)
	if err != nil {
		return nil, err
	}

	log.Printf("read %d bytes", n)
	result, err := message.DeserializeResult(buffer[:n], 0)
	if err != nil {
		return nil, err
	}

	return result, nil
}
