package client

import "net"

type Client struct {
	conn net.Conn
}

func NewClient(serverAddr string) (*Client, error) {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
	}, nil
}

func (c *Client) Connect() {
	// 建立隧道、转发数据等逻辑
}
