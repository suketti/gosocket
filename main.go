package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	IP := "127.0.0.1"
	fmt.Printf("IPアドレス: %s", IP)
	port := ":8097"
	tcpAddr, err := net.ResolveTCPAddr("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	listner, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("接続を待つ")
	for {
		conn, err := listner.Accept()
		if err != nil {
			continue
		} else {
			fmt.Println("接続された", conn.RemoteAddr().String())
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	messageBuf := make([]byte, 1024)
	messageLen, err := conn.Read(messageBuf)
	if err != nil {
		log.Fatal(err)
	}
	message := string(messageBuf[:messageLen])
	fmt.Println(message)

	// この書き込み処理がないと、連続してclientからメッセージを送信すると以下のエラーが発生する
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.Write([]byte(message))
}
