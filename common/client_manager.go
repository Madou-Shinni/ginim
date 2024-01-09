package common

import (
	"sync"
)

type ClientManager struct {
	clientExistMap        map[*Client]bool // 连接是否存在
	clientExistMapRWMutex sync.RWMutex     // 读写锁
	clientMap             map[uint]*Client // 用户id对应的连接
	clientMapRWMutex      sync.RWMutex     // 读写锁
	closeConn             chan *Client     // 连接关闭通道
	broadcast             chan broadcast   // 广播通道 向全部成员发送数据
	login                 chan *Client     // 用户登录通道
}

type broadcast struct {
	data []byte // 消息内容
}

func NewClientManager() (clientManager *ClientManager) {
	clientManager = &ClientManager{
		clientExistMap: make(map[*Client]bool),
		clientMap:      make(map[uint]*Client),
		closeConn:      make(chan *Client, 1000),
		login:          make(chan *Client, 1000),
	}
	return
}

// 事件处理
func (c *ClientManager) processed() {
	for {
		select {
		case data := <-c.broadcast:
			c.Broadcasting(data.data)
		case conn := <-c.closeConn:
			c.DelClient(conn)
			c.DelUserClient(conn.userId)
		case client := <-c.login:
			c.clientMap[client.userId] = client
		}
	}
}

func (c *ClientManager) AddClient(client *Client) {
	c.clientExistMapRWMutex.Lock()
	c.clientExistMap[client] = true
	c.clientExistMapRWMutex.Unlock()
}

func (c *ClientManager) DelUserClient(userId uint) {
	c.clientMapRWMutex.Lock()
	delete(c.clientMap, userId)
	c.clientMapRWMutex.Unlock()
}

func (c *ClientManager) DelClient(client *Client) {
	c.clientExistMapRWMutex.Lock()
	delete(c.clientExistMap, client)
	c.clientExistMapRWMutex.Unlock()
}

func (c *ClientManager) GetClients() []*Client {
	var clients []*Client
	for _, client := range c.clientMap {
		clients = append(clients, client)
	}
	return clients
}

func (c *ClientManager) Broadcasting(msg []byte) {
	var clients []*Client
	// 全员
	clients = append(clients, c.GetClients()...)

	for _, client := range clients {
		client.sendChan <- msg
	}
}
