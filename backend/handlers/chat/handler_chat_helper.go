package chat

import (
	"fmt"
	"io"
	"slices"
	"social-network/backend/models"
	"social-network/backend/services/chat"

	"github.com/gorilla/websocket"
)

type ClientList map[string][]*Client

type OnlineUsers struct {
	Type string        `json:"type"`
	Data []models.User `json:"data"`
}

type Client struct {
	session    *models.Session
	userId     string
	service    *chat.ChatService
	connection *websocket.Conn
	chatServer *ChatServer
	Message    chan *models.Message
	ErrorJson  chan *models.ErrorJson
	Online     chan *OnlineUsers
	FirstName  string
	LastName   string
}

func NewClient(conn *websocket.Conn, server *ChatServer, session *models.Session) *Client {
	return &Client{
		session:    session,
		service:    server.service,
		connection: conn,
		chatServer: server,
		Message:    make(chan *models.Message),
		ErrorJson:  make(chan *models.ErrorJson),
		Online:     make(chan *OnlineUsers),
		userId:     session.UserId,
		FirstName:  session.FirstName,
		LastName:   session.LastName,
	}
}

func (server *ChatServer) AddClient(client *Client) {
	server.Lock()
	server.clients[client.userId] = append(server.clients[client.userId], client)
	defer server.Unlock()
}

func (server *ChatServer) RemoveClient(client *Client, logged_out bool) {
	server.Lock()
	defer server.Unlock()
	switch logged_out {
	case true:
		if connections, ok := server.clients[client.userId]; ok {
			for _, conn := range connections {
				conn.connection.Close()
			}
			deleteConnection(server.clients, client.userId, client)
			go server.BroadCastOnlineStatus()
		}
		delete(server.clients, client.userId)
		go server.BroadCastOnlineStatus()
	case false:
		if _, ok := server.clients[client.userId]; ok {
			client.connection.Close()
			deleteConnection(server.clients, client.userId, client)
			go server.BroadCastOnlineStatus()
		}
	}
}

// first time working with channels and they seem great :!
func (client *Client) ReadMessages() {
	logged_out := false
	for {
		message := &models.Message{}
		err := client.connection.ReadJSON(&message)
		if err != nil {
			if err == io.ErrUnexpectedEOF {
				client.ErrorJson <- &models.ErrorJson{
					Message: models.MessageErr{
						Message:    " empty message field",
						ReceiverID: " empty receiver_id field",
						Type:       " empty type field",
						CreatedAt:  " empty createdAt field",
					},
				}
				continue
			}
			if isLogoutError(err) {
				logged_out = true
				fmt.Println("after", logged_out)
				// delete(client.chatServer.clients, client.userId)
				break
			}
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) || err == io.EOF {
				break
			}
		}

		message.SenderID = client.userId

		message_validated, errJson := client.chatServer.service.ValidateMessage(message)
		if errJson != nil {
			client.ErrorJson <- errJson
			continue
		}

		client.Message <- message_validated
		client.BroadCastTheMessage(message_validated)
	}
	defer client.chatServer.RemoveClient(client, logged_out)
}

// i used the channels buy not sure if this is the correct way to handle this

func (client *Client) WriteMessages() {
	defer client.chatServer.RemoveClient(client, false)

	for {
		select {
		case errJson := <-client.ErrorJson:
			err := client.connection.WriteJSON(errJson)
			if err != nil {
				return
			}
		case message := <-client.Message:
			err := client.connection.WriteJSON(message)
			if err != nil {
				return
			}
		case online_users := <-client.Online:
			err := client.connection.WriteJSON(online_users)
			if err != nil {
				return
			}
		}
	}
}

func (sender *Client) BroadCastTheMessage(message *models.Message) {
	// braodcast to the connections dyal sender
	sender.chatServer.Lock()
	defer sender.chatServer.Unlock()

	switch message.Type {
	case "message":
		for _, conn := range sender.chatServer.clients[sender.userId] {
			if conn.connection != sender.connection {
				conn.Message <- message
			}
		}
		// dyal receiver
		for _, value := range sender.chatServer.clients[message.TargetID] {
			value.Message <- message
		}
	case "read":
		for _, conn := range sender.chatServer.clients[sender.userId] {
			if conn.connection != sender.connection {
				conn.Message <- message
			}
		}
	case "typing":

	}
}

// dummy way to delete a connection but i'm done
func deleteConnection(clientList map[string][]*Client, userId string, client_to_be_deleted *Client) {
	index := -1
	for i, value := range clientList[userId] {
		if value == client_to_be_deleted {
			index = i
			break
		}
	}
	if index != -1 {
		clientList[userId] = slices.Delete(clientList[userId], index, index+1)
	}
}

// let's do it inside another function and make it specific to the client
func (server *ChatServer) BroadCastOnlineStatus() {
	server.Lock()
	defer server.Unlock()
	online_users := []models.User{}
	for _, connections := range server.clients {
		if len(connections) != 0 {
			online_users = append(online_users, models.User{ID: connections[0].userId, FirstName: connections[0].FirstName, LastName: connections[0].LastName})
		}
	}

	for _, connections := range server.clients {
		for _, conn := range connections {
			conn.Online <- &OnlineUsers{
				Type: "online",
				Data: online_users,
			}
		}
	}
}
