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

	conn.Write([]byte("UPLOAD\r\n"))
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

	conn.Write([]byte("DOWNLOAD\r\n"))
	data := make([]byte, 1024)
	read := 0
	for {
		n, err := conn.Read(data)
		read += n
		if err != nil {
			if read != 1024*60 {
				t.Fatalf("read %d bytes; expected %d", read, 1024*60)
			}
			break
		}
	}
}
