package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

var _ TelnetClient = (*client)(nil)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &client{address: address, timeout: timeout, in: in, out: out}
}

type client struct {
	address    string
	timeout    time.Duration
	in         io.ReadCloser
	out        io.Writer
	connection net.Conn
}

func (c *client) Connect() (err error) {
	c.connection, err = net.DialTimeout("tcp", c.address, c.timeout)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) Close() error {
	if c.connection == nil {
		return fmt.Errorf("connection error")
	}
	return c.connection.Close()
}

func (c *client) Send() (err error) {
	if c.connection == nil {
		return fmt.Errorf("connection error")
	}
	_, err = io.Copy(c.connection, c.in)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) Receive() (err error) {
	if c.connection == nil {
		return fmt.Errorf("connection error")
	}
	_, err = io.Copy(c.out, c.connection)
	if err != nil {
		return err
	}
	return nil
}
