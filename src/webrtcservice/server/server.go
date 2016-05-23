package server

import (
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
	}
}

// Serve start Server serve service
func (s *Server) Serve() error {
	log.SetFlags(0)
	http.HandleFunc("/connect", s.connect)
	http.HandleFunc("/", s.home)
	log.Println("Server: start on:", s.config.getAddr())
	log.Fatal(http.ListenAndServe(s.config.getAddr(), nil))
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
	homeTemplate.Execute(w, "ws://"+r.Host+"/connect")
}

func closeConnection(c *Connection, h *ConnectionHandler) {
	if err := h.CloseConnection(c); err != nil {
		c.Close()
	}
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };
    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("SEND ID");
            var msg = {
                action: 'SETID', 
                body: client_id.value
                }
            ws.send(JSON.stringify(msg, null, '\t'));
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RECEIVED: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };
    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        var msg = {
                action: 'SENDMSG', 
                from: client_id.value,
                to: client_to.value,
                body: input.value
                }
        ws.send(JSON.stringify(msg, null, '\t'));
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<label>Your Client ID</label><input id="client_id" type="text" value="">
<button id="open">Open</button>
<button id="close">Close</button>
<p><label>Client ID to connect</label><input id="client_to" type="text" value=""></p>
<p>Message to send<input id="input" type="text" value="Hello world!"></p>
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
