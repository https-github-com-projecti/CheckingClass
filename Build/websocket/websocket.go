package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

type clientManager struct {
	clients    map[*client]bool
	Broadcast  chan []byte
	register   chan *client
	unregister chan *client
}

type client struct {
	id     string
	socket *websocket.Conn
	send   chan []byte
}

//Manager is ...
var Manager = clientManager{
	Broadcast:  make(chan []byte),
	register:   make(chan *client),
	unregister: make(chan *client),
	clients:    make(map[*client]bool),
}

func (manager *clientManager) Start() {
	fmt.Println("HELLO FUNC START")
	for {
		select {
		case conn := <-Manager.register:
			// fmt.Println("func (manager *clientManager) start() conn manager.register : ", conn)
			manager.clients[conn] = true
			jsonMessage, _ := json.Marshal(&message{Content: "/A new socket has connected."})
			manager.send(jsonMessage, conn)
			// fmt.Println("func (manager *clientManager) start() conn manager.register manager.send(jsonMessage, conn) : ", conn)
		case conn := <-Manager.unregister:
			// fmt.Println("func (manager *clientManager) start() conn manager.unregister  : ", conn)
			if _, ok := manager.clients[conn]; ok {
				close(conn.send)
				delete(manager.clients, conn)
				jsonMessage, _ := json.Marshal(&message{Content: "/A socket has disconnected."})
				manager.send(jsonMessage, conn)
				// fmt.Println("func (manager *ClientManager) start() conn manager.unregister manager.send(jsonMessage, conn) : ", conn)
			}
		case message := <-Manager.Broadcast:
			// fmt.Println("func (manager *ClientManager) start() message manager.broadcast  : ", message)
			for conn := range manager.clients {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(manager.clients, conn)
				}
			}
		}
	}
}
func (manager *clientManager) send(message []byte, ignore *client) {
	fmt.Println("HELLO FUNC SEND")
	for conn := range manager.clients {
		// fmt.Println("conn := range manager.clients : ", conn)
		if conn != ignore {
			// fmt.Println("message form send func = ", message)
			conn.send <- message
		}
	}
}
func (c *client) read() {
	fmt.Println("HELLO FUNC READ")
	defer func() {
		Manager.unregister <- c
		c.socket.Close()
	}()

	for {
		_, messages, err := c.socket.ReadMessage()
		// fmt.Println("message from func read : ", messages)
		// asInt := string(messages)
		// fmt.Println("asInt : ", asInt)
		// num, _  := strconv.Atoi(asInt)
		if err != nil {
			Manager.unregister <- c
			c.socket.Close()
			break
		}
		jsonMessage, _ := json.Marshal(&message{Sender: c.id, Content: string(messages)})
		// fmt.Println("jsonMessage from func read : ", string(jsonMessage))
		Manager.Broadcast <- jsonMessage
		// fmt.Println("manager.broadcast <- jsonMessage : ", Manager.broadcast)
	}
}
func (c *client) write() {
	fmt.Println("HELLO FUNC WRITE")
	defer func() {
		c.socket.Close()
	}()

	// fmt.Println("test test = ", dataAuthen)
	for {
		// fmt.Println("test test = ", dataAuthen)
		select {
		case message, ok := <-c.send:
			// fmt.Println("c.send : ", c.send)
			// fmt.Println("message from func write : ", message)
			if !ok {
				// fmt.Println("message from func write !OK : ", ok)
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			// fmt.Println("message from func write OK : ", ok)
			// fmt.Println("message from func write message : ", message)
			// for {
			// 	time.Sleep(2 * time.Second)
			// 	jsonMessage, _ := json.Marshal(&Message{Sender: c.id, Content: string(message)})
			// 	c.socket.WriteMessage(websocket.TextMessage, jsonMessage)
			// }
			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

type message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

// type Message2 struct {
// 	Sender    string              `json:"sender,omitempty"`
// 	Recipient string              `json:"recipient,omitempty"`
// 	Content   []models.AuthenData `json:"content,omitempty"`
// }

var getClientID string

//Websockethandler is ...
func Websockethandler(c *gin.Context) {
	fmt.Println("HELLO FUNC WSHANDLER")
	// conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	// fmt.Println("conn : ", conn)
	if err != nil {
		// http.NotFound(res, req)
		return
	}
	u2, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	client := &client{id: u2.String(), socket: conn, send: make(chan []byte)}
	getClientID = client.id
	// cientID = client.id
	// fmt.Println("cientID = ", cientID)
	// fmt.Println("func wshandler client : ", client)
	// fmt.Println("func wshandler client.id : ", client.id)
	// fmt.Println("func wshandler client.socket : ", client.socket)
	// fmt.Println("func wshandler client.send : ", client.send)

	Manager.register <- client

	go client.read()
	go client.write()

	// out, err := json.Marshal("Success")
	// fmt.Printf("Out(type) = %T \n", out)
	// if err != nil {
	// 	panic(err)
	// }
	jsonMessage, _ := json.Marshal(&message{Sender: getClientID, Content: string("Success")})
	Manager.Broadcast <- jsonMessage
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		fmt.Println("HELLO FUNC UPGRADE")
		return true
	},
}
