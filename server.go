package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const protoErr = "Unknown protocol, expected ntwrk%s\r\n"
const actionErr = "Unknown action\r\n"

// startServer starts a network test server on `port`.
func startServer(port int) {
	addr := fmt.Sprintf(":%d", port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on %s\n", addr)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}

// handle starts an upload or download test on the provided TCP connection.
func handle(conn net.Conn) {
	defer conn.Close()

	remote := formatIP(conn.RemoteAddr())
	log.Printf("New connection from %s", remote)

	buf := bufio.NewReader(conn)
	msg, _ := buf.ReadString('\n')

	var clientProto, action string
	fmt.Sscanf(msg, protoFmt, &clientProto, &action)
	if clientProto != proto {
		msg := fmt.Sprintf(protoErr, proto)
		conn.Write([]byte(msg))
		return
	}

	switch action {
	case "download":
		bytes, _ := upload(conn, 0)
		log.Printf("Sent %d bytes to %s", bytes, remote)
	case "upload":
		bytes, _ := download(conn, 0)
		log.Printf("Received %d bytes from %s", bytes, remote)
	case "whoami":
		conn.Write([]byte(remote + "\r\n"))
		log.Printf("Replied to whoami from %s", remote)
	default:
		conn.Write([]byte(actionErr))
	}
}
