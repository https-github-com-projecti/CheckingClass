package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/narawichsaphimarn/backend/api"
	"github.com/narawichsaphimarn/backend/models"

	// "log"
	"encoding/json"
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
	// "strconv"
	// "flag"
	// "time"
	// "strings"
)

type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

type Client struct {
	id     string
	socket *websocket.Conn
	send   chan []byte
}

var manager = ClientManager{
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    make(map[*Client]bool),
}

var cientID string

func (manager *ClientManager) start() {
	fmt.Println("HELLO FUNC START")
	for {
		select {
		case conn := <-manager.register:
			fmt.Println("func (manager *ClientManager) start() conn manager.register : ", conn)
			manager.clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected."})
			manager.send(jsonMessage, conn)
			fmt.Println("func (manager *ClientManager) start() conn manager.register manager.send(jsonMessage, conn) : ", conn)
		case conn := <-manager.unregister:
			fmt.Println("func (manager *ClientManager) start() conn manager.unregister  : ", conn)
			if _, ok := manager.clients[conn]; ok {
				close(conn.send)
				delete(manager.clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected."})
				manager.send(jsonMessage, conn)
				fmt.Println("func (manager *ClientManager) start() conn manager.unregister manager.send(jsonMessage, conn) : ", conn)
			}
		case message := <-manager.broadcast:
			fmt.Println("func (manager *ClientManager) start() message manager.broadcast  : ", message)
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
func (manager *ClientManager) send(message []byte, ignore *Client) {
	fmt.Println("HELLO FUNC SEND")
	for conn := range manager.clients {
		fmt.Println("conn := range manager.clients : ", conn)
		if conn != ignore {
			fmt.Println("message form send func = ", message)
			conn.send <- message
		}
	}
}
func (c *Client) read() {
	fmt.Println("HELLO FUNC READ")
	defer func() {
		manager.unregister <- c
		c.socket.Close()
	}()

	for {
		_, message, err := c.socket.ReadMessage()
		fmt.Println("message from func read : ", message)
		asInt := string(message)
		fmt.Println("asInt : ", asInt)
		// num, _  := strconv.Atoi(asInt)
		if err != nil {
			manager.unregister <- c
			c.socket.Close()
			break
		}
		jsonMessage, _ := json.Marshal(&Message{Sender: c.id, Content: string(message)})
		fmt.Println("jsonMessage from func read : ", string(jsonMessage))
		manager.broadcast <- jsonMessage
		fmt.Println("manager.broadcast <- jsonMessage : ", manager.broadcast)
	}
}
func (c *Client) write() {
	fmt.Println("HELLO FUNC WRITE")
	defer func() {
		c.socket.Close()
	}()

	// fmt.Println("test test = ", dataAuthen)
	for {
		// fmt.Println("test test = ", dataAuthen)
		select {
		case message, ok := <-c.send:
			fmt.Println("c.send : ", c.send)
			fmt.Println("message from func write : ", message)
			if !ok {
				fmt.Println("message from func write !OK : ", ok)
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			fmt.Println("message from func write OK : ", ok)
			fmt.Println("message from func write message : ", message)
			// for {
			// 	time.Sleep(2 * time.Second)
			// 	jsonMessage, _ := json.Marshal(&Message{Sender: c.id, Content: string(message)})
			// 	c.socket.WriteMessage(websocket.TextMessage, jsonMessage)
			// }
			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

type Message2 struct {
	Sender    string              `json:"sender,omitempty"`
	Recipient string              `json:"recipient,omitempty"`
	Content   []models.AuthenData `json:"content,omitempty"`
}

func welcome(c *gin.Context) {
	c.JSON(200, "online")
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080/"},
		AllowMethods:     []string{"POST, GET, OPTIONS, PUT, DELETE, UPDATE"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:4200"
		},
		MaxAge: 12 * time.Hour,
	}))
	c1 := r.Group("/home")
	{
		c1.GET("/ping", welcome)
	}

	c2 := r.Group("/user")
	{
		c2.POST("/Add", api.NewUser)
		c2.GET("/allUsers", api.AllUsers)
		c2.POST("/login", api.UserLogin)
		c2.GET("/GETONE/:user", api.GetId)
		c2.GET("/getMyPic/:id", api.GetPicture)
	}

	c3 := r.Group("/Subject")
	{
		c3.POST("/Add", api.CreatClass)
		c3.GET("/allClass", api.AllClass)
		c3.GET("/GetMySubject/:user", api.MyClass)
		c3.GET("/GetOneSubject/:id", api.OneClass)
	}

	c4 := r.Group("/Attendance")
	{
		c4.POST("/Create", api.CreateBarcode)
		c4.GET("/allQr", api.AllQr)
		c4.GET("/getQRcode/:pass", api.MyQr)
		c4.GET("/getShowQrCode/:passOfcouse", api.GetshowQrCode)
		c4.GET("getClientId", sendClientID)
	}

	c5 := r.Group("authen")
	{
		c5.GET("/authenData", api.GetAuthenData)
	}

	ws := r.Group("/websocket")
	{
		// ws.GET("/ws", func(c *gin.Context) {
		//     wshandler(c.Writer, c.Request)
		// })
		ws.GET("/ws", wshandler)
		ws.POST("/testWebsocket", TestpostWebsocket)
		// ws.GET("/ping", ping)
	}

	fmt.Println("Starting application...")
	go manager.start()
	r.Run(":8080")
}

func wshandler(c *gin.Context) {
	fmt.Println("HELLO FUNC WSHANDLER")
	// conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	// fmt.Println("conn : ", conn)
	if err != nil {
		// http.NotFound(res, req)
		return
	}
	client := &Client{id: uuid.NewV4().String(), socket: conn, send: make(chan []byte)}
	cientID = client.id
	fmt.Println("cientID = ", cientID)
	fmt.Println("func wshandler client : ", client)
	fmt.Println("func wshandler client.id : ", client.id)
	fmt.Println("func wshandler client.socket : ", client.socket)
	fmt.Println("func wshandler client.send : ", client.send)

	manager.register <- client

	go client.read()
	go client.write()
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		fmt.Println("HELLO FUNC UPGRADE")
		return true
	},
}

var testDataAuthen []models.AuthenData

func TestpostWebsocket(c *gin.Context) {
	var sp models.AuthenData
	defer c.Request.Body.Close()
	if err := c.ShouldBindJSON(&sp); err != nil {
		return
	}

	testDataAuthen = append(testDataAuthen, sp)

	fmt.Println("testDataAuthen : ", testDataAuthen)

	// stringSlice := []string{"hello", "bye"}
	// stringByte := "\x00" + strings.Join(stringSlice, "\x20\x00") // x20 = space and x00 = null
	// fmt.Println([]byte(stringByte))
	// fmt.Println(string([]byte(stringByte)))

	out, err := json.Marshal(testDataAuthen)
	fmt.Printf("Out(type) = %T \n", out)
	if err != nil {
		panic(err)
	}
	jsonMessage, _ := json.Marshal(&Message2{Sender: "532220bb-e098-4622-9386-0bccb2462bd8", Content: testDataAuthen})
	manager.broadcast <- jsonMessage
}

func sendClientID(c *gin.Context) {
	defer c.Request.Body.Close()
	c.JSON(http.StatusOK, gin.H{
		"client": cientID,
	})
}
