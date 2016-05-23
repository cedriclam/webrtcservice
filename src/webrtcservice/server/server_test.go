package server

import (
	"net"
	"testing"

	"webrtcservice/client"
	"webrtcservice/protocol"
)

func TestConnectionClient(t *testing.T) {
	config := &Config{
		Port:          8090,
		Host:          GetLocalIP(),
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

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
