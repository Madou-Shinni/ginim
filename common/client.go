package common

import (
	"encoding/json"
	"fmt"
	"github.com/Madou-Shinni/gin-quickstart/constants"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn     *websocket.Conn // 连接
	sendChan chan []byte     // 消息通道
	userId   uint            // 用户id
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		conn:     conn,
		sendChan: make(chan []byte, 1000),
	}
}

func (c *Client) read() {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			// 读取数据错误 关闭连接
			close(c.sendChan)
			return
		}
		// 处理消息
		var msg domain.Message
		json.Unmarshal(message, &msg)
		fmt.Printf("message: %s; msg: %v\n", string(message), msg)
		c.processed(msg)
	}
}

func (c *Client) write() {
	for {
		select {
		case message, ok := <-c.sendChan:
			if !ok {
				// 发送数据错误 关闭连接
				c.conn.Close()
				clientManager.closeConn <- c
				return
			}
			c.conn.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func (c *Client) Send(msg []byte) {
	c.sendChan <- msg
}

func (c *Client) Login(userId uint) {
	c.userId = userId
}

func (c *Client) processed(message domain.Message) {
	switch message.Type {
	case constants.MessageTypePrivate:
		marshal, _ := json.Marshal(message)
		if clientManager.clientMap[message.To] == nil {
			fmt.Printf("用户不在线: %d\n", message.To)
			return
		}
		clientManager.clientMap[message.To].Send(marshal)
	}
}
