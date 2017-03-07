package main

import (
	"net"
	"testing"
	"time"
)

const expected = 1024

func listen(t *testing.T, addr string, fn func(net.Conn)) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		conn, _ := ln.Accept()
		fn(conn)
		conn.Close()
	}()
}

func TestDownload(t *testing.T) {
	addr := ":1616"
	listen(t, addr, func(conn net.Conn) {
		for {
			conn.Write([]byte(DATA))
		}
	})

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}

	bytes, err := download(conn, time.Millisecond)
	if err != nil {
		t.Fatal(err)
	} else if bytes < expected {
		t.Fatalf("too few bytes read: got %d want %d", bytes, expected)
	}
}

func TestUpload(t *testing.T) {
	addr := ":1617"
	listen(t, addr, func(conn net.Conn) {
		data := make([]byte, 1024)
		for {
			conn.Read(data)
		}
	})

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}

	bytes, err := upload(conn, time.Millisecond)
	if err != nil {
		t.Fatal(err)
	} else if bytes < expected {
		t.Fatalf("too few bytes written: got %d want %d", bytes, expected)
	}
}
