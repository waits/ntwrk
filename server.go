package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

const DATA = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

// startServer starts a network test server on port 8080.
func startServer() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listening on :8080")
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

	addr := conn.RemoteAddr()
	log.Printf("New connection from %v", addr)

	buf := bufio.NewReader(conn)
	msg, _ := buf.ReadString('\n')
	switch msg {
	case "UPLOAD\r\n":
		perform(upload, conn)
	case "DOWNLOAD\r\n":
		perform(download, conn)
	default:
		return
	}
}

// perform runs `test` and reports the time taken.
func perform(test func(net.Conn) int, conn net.Conn) {
	t := time.Now()
	bytes := test(conn)
	elapsed := time.Since(t)
	log.Printf("Processed %d bytes in %v", bytes, elapsed)
}

// upload reads data from `conn` and returns the number of bytes read.
func upload(conn net.Conn) (bytes int) {
	buf := bufio.NewReader(conn)
	for {
		data, err := buf.ReadString('\n')
		if err != nil {
			break
		}
		bytes += len(data)
	}
	return
}

// download writes data to `conn` and returns the number of bytes written.
func download(conn net.Conn) (bytes int) {
	buf := bufio.NewWriter(conn)
	for i := 0; i < 1024; i++ {
		buf.WriteString(DATA)
		bytes += len(DATA)
	}
	return
}
