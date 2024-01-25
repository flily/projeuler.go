package framework

import (
	"time"

	"github.com/flily/projeuler.go/framework/connection"
	"github.com/flily/projeuler.go/framework/message"
)

type Client struct {
	client         *connection.Client
	ProblemTimeout time.Duration
	MethodTimeout  time.Duration
}

func NewClient(host string, port int) (*Client, error) {
	client, err := connection.NewClient(host, port)
	if err != nil {
		return nil, err
	}

	c := &Client{
		client: client,
	}

	return c, nil
}

func (c *Client) Close() {
	c.client.Close()
}

func (c *Client) SetTimeout(problemTimeout, methodTimeout time.Duration) {
	c.ProblemTimeout = problemTimeout
	c.MethodTimeout = methodTimeout
}

func (c *Client) Run(problemId int, method string) (*Result, error) {
	request := message.NewRunMessage(problemId, method)
	request.SetTimeout(c.ProblemTimeout, c.MethodTimeout)

	resultMessage, err := c.client.Run(request)
	if err != nil {
		return nil, err
	}

	result := NewResult()
	result.FromMessage(resultMessage)
	return result, nil
}
