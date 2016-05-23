package server

import (
	"webrtcservice/protocol"

	"github.com/gorilla/websocket"
)

// Connection aggregates feed and method for manager connection
type Connection struct {
	id   string
	conn *websocket.Conn
}

// ConnectionInterface interface for the connection
type ConnectionInterface interface {
	// GetID
	GetID() string
	// WriteMessage writes message to the current connection
	WriteMessage(msg *protocol.Message) error
	// Close used to store Connection resources
	Close()
}

// NewConnection retuns new Connection instance
func NewConnection(c *websocket.Conn) *Connection {
	return &Connection{
		conn: c,
	}
}

// SetID set connection ID
func (c *Connection) SetID(id string) {
	c.id = id
}

// Close used to store Connection resources
func (c *Connection) Close() {
	c.Close()
}

// GetID returns connection Id string
func (c Connection) GetID() string {
	return c.id
}

// WriteMessage writes message to the current connection
func (c *Connection) WriteMessage(msg *protocol.Message) error {
	return c.conn.WriteJSON(msg)
}
