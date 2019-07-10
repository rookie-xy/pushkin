package pushkin

import "net"

type Conn struct {
	net.Conn
	proto string
}

func NewConn(config *ConnConfig) (*Conn, error) {
	var err error

	c := &Conn{
		proto: "tcp",
	}

	c.Conn, err = net.Dial(c.proto, config.Host)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Conn) Read() error {
	c.Conn.Read()
	return nil
}

func (c *Conn) Write() error {
	return nil
}