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
	Egress chan string
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
		if string(message) == "hello" {
			fmt.Println("SENDING MESSAGE")
			c.Egress <- "world"
		}
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
		Egress: make(chan string, 1),
	}
}

type ClientManager struct {
	sync.RWMutex
	Clients map[string][]*Client
}

func NewClientManger() *ClientManager {
	return &ClientManager{
		Clients: make(map[string][]*Client),
	}
}

func (cm *ClientManager) AddClient(id string, client *Client) {
	cm.Lock()
	defer cm.Unlock()
	cm.Clients[id] = append(cm.Clients[id], client)
}

func (cm *ClientManager) RemoveClient(id string, conn *websocket.Conn) {
	cm.Lock()
	defer cm.Unlock()
	clients, ok := cm.Clients[id]
	if ok {
		defer conn.Close()
		for i, c := range clients {
			if c.Conn == conn {
				cm.Clients[id] = append(cm.Clients[id][:i], cm.Clients[id][i+1:]...)
				break
			}
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
		id := c.Query("id")
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		client = NewClient(conn)
		clientManager.AddClient(id, client)
		go client.Read()
		go client.Write()

	})
	router.Run(":8080")
}
