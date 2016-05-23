package server

import (
	"testing"
	"webrtcservice/protocol"
)

func TestAddConnection(t *testing.T) {
	h := NewConnectionHandler()

	h.AddConnection(&fakeConnection{
		ID:  "connectionID1",
		Err: nil,
	})

	if len(h.conns) != 1 {
		t.Error("h.conns len should be equal to 1, current:", len(h.conns))
	}

	h.AddConnection(&fakeConnection{
		ID:  "connectionID2",
		Err: nil,
	})

	if len(h.conns) != 2 {
		t.Error("h.conns len should be equal to 1, current:", len(h.conns))
	}
}

func TestAddConnectionThatsAlreadyExist(t *testing.T) {
	h := NewConnectionHandler()

	c := &fakeConnection{
		ID:  "connectionID1",
		Err: nil,
	}

	h.AddConnection(c)

	if len(h.conns) != 1 {
		t.Error("h.conns len should be equal to 1, current:", len(h.conns))
	}

	err := h.AddConnection(c)

	if err == nil {
		t.Error("AddConnection should returns an Error")
	}
}

func TestAddDeleteConnection(t *testing.T) {
	h := NewConnectionHandler()

	c := &fakeConnection{
		ID:  "connectionID1",
		Err: nil,
	}

	h.AddConnection(c)
	if len(h.conns) != 1 {
		t.Error("h.conns len should be equal to 1, current:", len(h.conns))
	}

	h.CloseConnection(c)
	if len(h.conns) != 0 {
		t.Error("h.conns len should be equal to 0, current:", len(h.conns))
	}

}

func TestForwardMessage(t *testing.T) {
	h := NewConnectionHandler()

	c1 := &fakeConnection{
		ID:  "connectionID1",
		Err: nil,
	}

	c2 := &fakeConnection{
		ID:  "connectionID2",
		Err: nil,
	}

	h.AddConnection(c1)
	h.AddConnection(c2)

	msg := protocol.NewMessage()
	msg.To = "connectionID2"
	msg.Body = "HelloMsg"

	if err := h.ForwardMessage(msg); err != nil {
		t.Error("Forward should not return an error, err:", err)
	}

	if c2.MessageToSend == nil {
		t.Error("MessageToSend should be nil")
	}

	if c2.MessageToSend.Body != "HelloMsg" {
		t.Error("Message should be equal to 'HelloMsg, received:", c2.MessageToSend)
	}
}

func TestForwardMessageError(t *testing.T) {
	h := NewConnectionHandler()

	msg := protocol.NewMessage()
	msg.To = "connectionID2"
	msg.Body = "HelloMsg"

	if err := h.ForwardMessage(msg); err == nil {
		t.Error("Forward should return an error")
	}

	c1 := &fakeConnection{
		ID:  "connectionID1",
		Err: nil,
	}
	h.AddConnection(c1)

	if err := h.ForwardMessage(msg); err == nil {
		t.Error("Forward should return an error")
	}
}

type fakeConnection struct {
	ID            string
	Err           error
	MessageToSend *protocol.Message
}

// GetID
func (c fakeConnection) GetID() string {
	return c.ID
}

// WriteMessage writes message to the current connection
func (c *fakeConnection) WriteMessage(msg *protocol.Message) error {
	c.MessageToSend = msg
	return c.Err
}

// Close used to store Connection resources
func (c *fakeConnection) Close() {

}
