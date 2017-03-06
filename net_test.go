package main

import (
	"net"
	"testing"
)

const TEST_SIZE = 16 * 1024

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
		for bytes := 0; bytes < TEST_SIZE; {
			n, _ := conn.Write([]byte(DATA))
			bytes += n
		}
	})

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}

	bytes, err := download(conn, TEST_SIZE)
	if err != nil {
		t.Fatal(err)
	} else if bytes != TEST_SIZE {
		t.Fatalf("incorrect value: got %d want %d", bytes, TEST_SIZE)
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

	bytes, err := upload(conn, TEST_SIZE)
	if err != nil {
		t.Fatal(err)
	} else if bytes != TEST_SIZE {
		t.Fatalf("incorrect value: got %d want %d", bytes, TEST_SIZE)
	}
}
