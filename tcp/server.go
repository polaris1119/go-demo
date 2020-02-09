package main

import (
	"bufio"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":2020")
	if err != nil {
		panic(err)
	}

	go sendMessage2Other()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept error:", err)
			time.Sleep(100)
			continue
		}

		go handleConn(conn)
	}
}

type User struct {
	name         string
	messageQueue chan string

	conn net.Conn
}

type Message struct {
	content string
	owner   *User
}

var (
	EnteringQueue = make(chan *User)
	LeavingQueue  = make(chan *User)
	MessageQueue  = make(chan *Message, 8)
)

func sendMessage2Other() {
	users := make(map[string]*User)
	for {
		select {
		case user := <-EnteringQueue:
			users[user.name] = user
		case user := <-LeavingQueue:
			delete(users, user.name)
			close(user.messageQueue)
		case message := <-MessageQueue:
			slice := strings.SplitN(message.content, " ", 2)
			if !strings.HasPrefix(slice[0], "@") {
				message.owner.messageQueue <- "message must begin with @xxx, try again!"
				break
			}

			if len(slice) < 2 {
				message.owner.messageQueue <- "after @xxx with space, try again!"
				break
			}
			toUserName := slice[0][1:]

			userOnline := false
			for name, user := range users {
				if name == toUserName {
					userOnline = true
					user.messageQueue <- "from @" + message.owner.name + " " + slice[1]
					break
				}
			}

			if !userOnline {
				message.owner.messageQueue <- toUserName + " user is not online!"
			}
		}
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	// 第一条消息，告知自己是谁
	reader := bufio.NewReader(conn)
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Read name error:", err)
		return
	}
	name = strings.TrimSpace(name)

	log.Println("User:", name, "Connect!")

	user := &User{
		name: name,
		messageQueue: make(chan string, 8),
		conn: conn,
	}
	// register user
	EnteringQueue <- user

	// new goroutine for sending message
	go serverSendMessage(user)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := &Message{
			content: scanner.Text(),
			owner: user,
		}
		MessageQueue <- message
	}

	LeavingQueue <- user
	log.Println("User:", name, "Leaving!")
}

func serverSendMessage(user *User) {
	for msg := range user.messageQueue {
		log.Println("send message:", msg, "to:", user.name)
		user.conn.Write([]byte(msg+"\n"))
	}
}