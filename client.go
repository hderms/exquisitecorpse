package main

import (
	"bufio"
	"net"
)

type Client struct {
	reader *bufio.Reader
	client net.Conn
	writer *bufio.Writer
}

func NewClient(reader *bufio.Reader, writer *bufio.Writer, c net.Conn) *Client {

	return &Client{reader: reader, client: c, writer: writer}

}

func (x *Client) ReadString(delim byte) (line string, err error) {

	return x.reader.ReadString(delim)
}

func (c *Client) SendMessage(text string) {

	c.writer.WriteString(text)
	c.writer.Flush()
}
