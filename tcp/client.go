package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

var name string

func init() {
	flag.StringVar(&name, "name", "", "who are you")
}

func main() {
	flag.Parse()
	if name == "" {
		flag.PrintDefaults()
		return
	}

	conn, err := net.Dial("tcp", "localhost:2020")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(name+"\n"))
	if err != nil {
		log.Println("write name error：", err)
		return
	}

	go clientSendMessage(conn)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println("receive message：", scanner.Text())
	}
}

func clientSendMessage(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		conn.Write([]byte(scanner.Text()+"\n"))
	}
}