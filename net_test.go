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
			conn.Write([]byte(sampleData))
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

func TestEcho(t *testing.T) {
	addr := ":1618"
	msg := "echo\r\n"
	listen(t, addr, func(conn net.Conn) {
		resp := make([]byte, 6)
		for {
			conn.Write([]byte(msg))
			conn.Read(resp)
			if string(resp) != msg {
				t.Fatalf("invalid echo: got %v want %v", resp, []byte(msg))
			}
		}
	})

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}

	echo(conn, time.Millisecond)
}
