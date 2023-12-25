package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	Conn   *websocket.Conn
	Egress chan []byte
}

func (c *Client) Read() {
	defer c.Conn.Close()
	c.Conn.SetReadLimit(512)
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Println("ERROR:", err)
			break
		}
		c.Egress <- message
		fmt.Println(string(message))
	}
	fmt.Println("EXITING READ")
}

func (c *Client) Write() {
	defer c.Conn.Close()
	fmt.Println("STARTING WRITE", c.Egress)
	for msg := range c.Egress {
		if err := c.Conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			fmt.Println("ERROR WRITING MESSAGE:", err)
			return
		}
		fmt.Println("SENT MESSAGE:", string(msg))
	}
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		Conn:   conn,
		Egress: make(chan []byte, 1),
	}
}

type ClientManager struct {
	sync.RWMutex
	Clients []*Client
}

func NewClientManger() *ClientManager {
	return &ClientManager{
		Clients: make([]*Client, 0),
	}
}

func (cm *ClientManager) AddClient(client *Client) {
	cm.Lock()
	defer cm.Unlock()
	cm.Clients = append(cm.Clients, client)
}

func (cm *ClientManager) RemoveClient(conn *websocket.Conn) {
	cm.Lock()
	defer cm.Unlock()

	for i, client := range cm.Clients {
		if client.Conn == conn {
			cm.Clients = append(cm.Clients[:i], cm.Clients[i+1:]...)
			break
		}
	}
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	clientManager := NewClientManger()
	var client *Client

	router.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		client = NewClient(conn)
		clientManager.AddClient(client)
		go client.Read()
		go client.Write()

	})
	router.Run(":8080")
}
