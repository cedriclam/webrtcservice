package client

import (
	"net/http"
	"testing"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func TestClientInit(t *testing.T) {
	srv := &fakeServer{T: t}
	http.HandleFunc("/connect", srv.connect)
	go http.ListenAndServe("0.0.0.0:8060", nil)

	client := NewClient("id1")
	if err := client.Init("0.0.0.0:8060"); err != nil {
		t.Error("client init error, err:", err)
	}

	client.Close()
}

type fakeServer struct {
	T *testing.T
}

func (f fakeServer) connect(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		f.T.Error("Connect Erro:", err)
		return
	}
	defer c.Close()

	for {
		if _, _, err = c.ReadMessage(); err != nil {
			f.T.Error("Connect Error:", err)
		}
		return
	}
}
