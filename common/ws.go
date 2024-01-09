package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

var (
	clientManager = NewClientManager()
)

func init() {
	// 添加处理程序
	go clientManager.processed()
	fmt.Println("WebSocket 启动程序成功")
}

func WsHandler(c *gin.Context) {

	// 升级协议
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		fmt.Println("升级协议", "ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])

		return true
	}}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)

		return
	}

	fmt.Println("webSocket 建立连接:", conn.RemoteAddr().String())
	client := NewClient(conn)

	userIdStr := c.Query("userId")
	userId, _ := strconv.ParseUint(userIdStr, 10, 64)
	client.userId = uint(userId)

	// 连接管理
	clientManager.login <- client
	clientManager.AddClient(client)

	go client.read()
	go client.write()
}
