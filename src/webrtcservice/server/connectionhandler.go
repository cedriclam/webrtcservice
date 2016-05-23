package server

import (
	"sync"
	"webrtcservice/protocol"

	"github.com/golang/glog"
)

// ConnectionHandler agregate feeds and method for handling connection
type ConnectionHandler struct {
	mutex sync.RWMutex
	conns map[string]ConnectionInterface
}

// NewConnectionHandler return new ConnectionHandler instance
func NewConnectionHandler() *ConnectionHandler {
	return &ConnectionHandler{
		conns: make(map[string]ConnectionInterface),
	}
}

// AddConnection add connection to the handler
func (c *ConnectionHandler) AddConnection(conn ConnectionInterface) error {
	defer c.mutex.Unlock()
	c.mutex.Lock()
	if _, ok := c.conns[conn.GetID()]; ok {
		glog.Info("Connection Already exist")
		return ErrConnectionAlreadyExist
	}

	c.conns[conn.GetID()] = conn

	return nil
}

// CloseConnection close a specific connection
func (c *ConnectionHandler) CloseConnection(conn ConnectionInterface) error {
	defer c.mutex.Unlock()
	c.mutex.Lock()
	if obj, ok := c.conns[conn.GetID()]; ok {
		delete(c.conns, obj.GetID())
	} else {
		return ErrConnectionDidntExist
	}

	return nil
}

// Close used to close all connections
func (c *ConnectionHandler) Close() {
	defer c.mutex.Unlock()
	c.mutex.Lock()
	for key, conn := range c.conns {
		conn.Close()
		delete(c.conns, key)
	}
}

// ForwardMessage used to forward a message to another connection
func (c *ConnectionHandler) ForwardMessage(msg *protocol.Message) error {
	defer c.mutex.RUnlock()
	c.mutex.RLock()
	var err error
	if conn, ok := c.conns[msg.To]; ok {
		err = conn.WriteMessage(msg)
	} else {
		err = ErrUnknowClientID
	}

	return err
}
