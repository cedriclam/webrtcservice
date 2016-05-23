package client

import (
	"net/url"

	"webrtcservice/protocol"

	"github.com/golang/glog"

	"github.com/gorilla/websocket"
)

// Interface Client interface
type Interface interface {
	// Init Client conversation
	Init(addr string) error
	// SendMessage used to send a message to the server
	SendMessage(msg *protocol.Message) error
	// Close used to close Client ressources
	Close() error
}

// Client used to aggregate feeds and method for the client
type Client struct {
	ID   string
	conn *websocket.Conn
}

// NewClient returns new Client instance
func NewClient(id string) *Client {
	return &Client{ID: id}
}

// Close used to close Client ressources
func (c *Client) Close() error {
	err := c.conn.Close()
	return err
}

// Init Client conversation with service
func (c *Client) Init(addr string) error {
	var err error
	u := url.URL{
		Scheme: "ws",
		Host:   addr,
		Path:   "/connect",
	}

	glog.Infof("connecting to %s", u.String())

	c.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		glog.Error("websocket dial err:", err)
		return err
	}

	msg := &protocol.Message{
		Action: protocol.SetID,
		Body:   c.ID,
	}

	return c.SendMessage(msg)
}

// SendMessage used to send a message to the server
func (c *Client) SendMessage(msg *protocol.Message) error {
	if err := c.conn.WriteJSON(msg); err != nil {
		glog.Error("Connection Error, err:", err)
	}

	return nil
}
