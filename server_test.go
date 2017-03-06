package main

import (
	"net"
	"testing"
	"time"
)

func TestUpload(t *testing.T) {
	addr := ":1600"
	go startServer(addr)
	time.Sleep(time.Millisecond)

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	conn.Write([]byte("UPLOAD\n"))
	for i := 0; i < 64; i++ {
		conn.Write([]byte(DATA))
	}
}

func TestDownload(t *testing.T) {
	addr := ":1601"
	go startServer(addr)
	time.Sleep(time.Millisecond)

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	conn.Write([]byte("DOWNLOAD\n"))
	data := make([]byte, 1024)
	bytes := 0
	for {
		n, err := conn.Read(data)
		bytes += n
		if err != nil {
			if bytes != UPLOAD_LIMIT {
				t.Fatalf("received %d bytes; expected %d", bytes, UPLOAD_LIMIT)
			}
			break
		}
	}
}
