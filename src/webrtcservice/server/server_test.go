package server

import (
	"testing"

	"webrtcservice/client"
	"webrtcservice/protocol"
)

func TestConnectionClient(t *testing.T) {
	config := &Config{
		Port:          32000,
		Host:          "0.0.0.0",
		IndexFileName: "../../template/index.html",
	}

	srv := NewServer(config)
	defer srv.Close()

	go srv.Serve()

	client := client.NewClient("ID1")
	err := client.Init(config.getAddr())
	if err != nil {
		t.Error("client.Init should not return an error, err:", err)
	}

	msg := protocol.NewMessageWithActionAndBody("ID1", protocol.SendMSG, "test")
	client.SendMessage(msg)

	client.Close()
}
