package main

import (
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{}

type severityType int32

const (
	debug        severityType = 0
	info         severityType = 1
	trace        severityType = 2
	warning      severityType = 3
	custom_error severityType = 4
)

func (this severityType) String() string {
	switch this {
	case debug:
		return "debug"
	case info:
		return "info"
	case trace:
		return "trace"
	case warning:
		return "warning"
	default:
		return "error"
	}

}

type messageClient struct {
	Severity severityType      `json:"severity"` //Переопределить маршал и анмаршал для этого типа enum websocket.JSON.Receive
	Text     string            `json:"text"`
	Meta     map[string]string `json:"meta"`
}

type server struct {
	addr *string
}

func (this *server) serve() {
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*this.addr, nil))
}

func con(s string) severityType {
	switch s {
	case "\"debug\"":
		return debug
	case "\"info\"":
		return info
	case "\"trace\"":
		return trace
	case "\"warning\"":
		return warning
	}
	return custom_error
}

func (cf *severityType) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"debug"`, `"info"`, `"trace"`, `"warning"`:
		*cf = con(string(data))
		return nil
	default:
		*cf = custom_error
		return nil
	}
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer c.Close()
	for {
		/*mt, message, err := c.ReadMessage()
		c.ReadJSON(&message)
		if err != nil {
			log.Println("read:", err)
			break
		}*/
		mes := messageClient{Meta: map[string]string{}}
		err = c.ReadJSON(&mes)
		//bytes := []byte(message)
		//err = json.Unmarshal(bytes, &mes)
		if err != nil {
			log.Printf("%s", "Error!!!")
			continue
		}
		t := time.Now()
		s := t.String()
		userId := mes.Meta["user_id"]
		sessionId := mes.Meta["session_id"]
		delete(mes.Meta, "user_id")
		delete(mes.Meta, "session_id")
		log.Printf("%s %s %s %s: %s %v", s[:19], mes.Severity, userId, sessionId, string(mes.Text), mes.Meta)
		//Тут коннект от TCP
		//log.Printf("%d %s %s", mes.Severity, string(mes.Text), string(mes.Meta['0']))

	}
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.textContent = message;
        output.appendChild(d);
        output.scroll(0, output.scrollHeight);
    };
    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
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
        ws.send(input.value);
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
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output" style="max-height: 70vh;overflow-y: scroll;"></div>
</td></tr></table>
</body>
</html>
`))
