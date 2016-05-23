package protocol

import "testing"

func TestNewMessage(t *testing.T) {
	msg := NewMessage()
	if msg.To != "" {
		t.Error("Error msg.To != '', current:", msg.To)
	}

	if msg.Action != "" {
		t.Error("Error msg.Action != SetID, current:", msg.Action)
	}

	if msg.Body != "" {
		t.Error("Error msg.Body != '', body:", msg.Body)
	}

}

func TestNewMessageWithActionAndBody(t *testing.T) {
	msg := NewMessageWithActionAndBody("ID1", SetID, "ID2")
	if msg.To != "ID1" {
		t.Error("Error msg.To != ID1, current:", msg.To)
	}

	if msg.Action != SetID {
		t.Error("Error msg.Action != SetID, current:", msg.Action)
	}

	if msg.Body != "ID2" {
		t.Error("Error msg.Body != ID2, body:", msg.Body)
	}

}
