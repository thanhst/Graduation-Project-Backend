package websocket

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	UserID string
	Role   string
	Conn   *websocket.Conn
}

func (c *Client) SendJSON(data interface{}) error {
	return c.Conn.WriteJSON(data)
}
