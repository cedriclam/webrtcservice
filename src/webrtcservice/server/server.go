package server

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"webrtcservice/protocol"

	"github.com/golang/glog"
	"github.com/gorilla/websocket"
)

// Server aggreate data and method use by the signalisation server
type Server struct {
	config   *Config
	indexTpl *template.Template
	upgrader *websocket.Upgrader
	handler  *ConnectionHandler
}

// Interface Server interface
type Interface interface {
	// start Server
	Serve() error
	// Close used to close Server resources
	Close()
}

// NewServer return new instance of Server
func NewServer(c *Config) Interface {
	return &Server{
		config:   c,
		upgrader: &websocket.Upgrader{},
		handler:  NewConnectionHandler(),
		indexTpl: loadIndexTemplate(c.IndexFileName),
	}
}

// Serve start Server serve service
func (s *Server) Serve() error {
	log.SetFlags(0)
	http.HandleFunc("/connect", s.connect)
	http.HandleFunc("/", s.home)
	fmt.Println("[Server] start on:", s.config.getAddr())
	glog.Fatal(http.ListenAndServe(s.config.getAddr(), nil))
	return nil
}

// Close used to close Server resources
func (s *Server) Close() {
	s.handler.Close()
}

func (s *Server) connect(w http.ResponseWriter, r *http.Request) {
	c, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		glog.Info("upgrade:", err)
		return
	}
	conn := NewConnection(c)

	defer closeConnection(conn, s.handler)

	for {
		message := protocol.NewMessage()
		if err = c.ReadJSON(message); err != nil {
			glog.Info("read err:", err)
			break
		}
		if err = handleAction(message, s.handler, conn); err != nil {
			glog.Info("handle Action err:", err)
			break
		}
	}
}

func handleAction(msg *protocol.Message, h *ConnectionHandler, conn *Connection) error {
	switch msg.Action {
	case protocol.SetID:
		conn.SetID(msg.Body)
		h.AddConnection(conn)
	case protocol.SendMSG:
		if err := h.ForwardMessage(msg); err != nil {
			return err
		}
	default:
		glog.Info("Unknow action:", msg.Action)
	}

	return nil
}

func (s *Server) home(w http.ResponseWriter, r *http.Request) {
	s.indexTpl.ExecuteTemplate(w, "index.html", "ws://"+r.Host+"/connect")
}

func closeConnection(c *Connection, h *ConnectionHandler) {
	if err := h.CloseConnection(c); err != nil {
		c.Close()
	}
}

func loadIndexTemplate(filename string) *template.Template {
	tpl, err := template.ParseFiles(filename)
	return template.Must(tpl, err)
}

var homeTemplate = template.Must(template.New("").Parse(``))
