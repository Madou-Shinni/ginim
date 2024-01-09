package common

import (
	"encoding/json"
	"fmt"
	"github.com/Madou-Shinni/gin-quickstart/constants"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"testing"
	"time"
)

func TestConnUser1(t *testing.T) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8080", Path: "/ws/conn"}
	q := u.Query()
	q.Add("userId", "1")
	u.RawQuery = q.Encode()
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			fmt.Printf("received: %s\n", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	from, _ := strconv.ParseUint("1", 10, 64)
	to, _ := strconv.ParseUint("2", 10, 64)
	msg := domain.Message{
		From:      uint(from),
		To:        uint(to),
		Content:   "hi 2",
		Type:      constants.MessageTypePrivate,
		MediaType: constants.MediaTypeText,
	}

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			marshal, _ := json.Marshal(msg)
			err := c.WriteMessage(websocket.TextMessage, marshal)
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func TestConnUser2(t *testing.T) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8080", Path: "/ws/conn"}
	q := u.Query()
	q.Add("userId", "2")
	u.RawQuery = q.Encode()
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			fmt.Printf("received: %s\n", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	from, _ := strconv.ParseUint("2", 10, 64)
	to, _ := strconv.ParseUint("1", 10, 64)
	msg := domain.Message{
		From:      uint(from),
		To:        uint(to),
		Content:   "hello 1",
		Type:      constants.MessageTypePrivate,
		MediaType: constants.MediaTypeText,
	}

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			marshal, _ := json.Marshal(msg)
			err := c.WriteMessage(websocket.TextMessage, marshal)
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
