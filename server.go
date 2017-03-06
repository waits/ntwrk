package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

// startServer starts a network test server on `addr`.
func startServer(addr string) {
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

	remote := conn.RemoteAddr()
	log.Printf("New connection from %v", remote)

	buf := bufio.NewReader(conn)
	msg, _ := buf.ReadString('\n')
	action := strings.TrimSpace(msg)
	switch action {
	case "DOWNLOAD":
		bytes, _ := upload(conn, MAX)
		log.Printf("Sent %d bytes to %v", bytes, remote)
	case "UPLOAD":
		bytes, _ := download(conn, MAX)
		log.Printf("Received %d bytes from %v", bytes, remote)
	default:
		log.Fatalf("Unknown action %s", action)
	}
}
