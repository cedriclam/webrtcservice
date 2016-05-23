package protocol

// Action protocol Action type
type Action string

const (
	// SetID Set connection ID
	SetID Action = "SETID"
	// SendMSG send a message
	SendMSG Action = "SENDMSG"
)

// Message represents the protocol message structure
type Message struct {
	Action Action
	From   string `json:",omitempty"`
	To     string `json:",omitempty"`
	Body   string `json:",omitempty"`
}

// NewMessage returns new Message instance
func NewMessage() *Message {
	return &Message{}
}

// NewMessageWithActionAndBody returns new Message instance
func NewMessageWithActionAndBody(to string, action Action, body string) *Message {
	return &Message{
		To:     to,
		Action: action,
		Body:   body,
	}
}
