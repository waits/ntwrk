package main

import "net"
import "strings"
import "time"

const sampleData = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

// download reads data from `conn` and returns the number of bytes read.
func download(conn net.Conn, timeout time.Duration) (bytes int, err error) {
	if timeout > 0 {
		conn.SetDeadline(time.Now().Add(timeout))
	}

	data := make([]byte, 1024)
	for {
		n, err := conn.Read(data)
		if err != nil {
			break
		}
		bytes += n
	}
	return
}

// upload writes data to `conn` and returns the number of bytes written.
func upload(conn net.Conn, timeout time.Duration) (bytes int, err error) {
	if timeout > 0 {
		conn.SetDeadline(time.Now().Add(timeout))
	}

	chunk := strings.Repeat(sampleData, 128)
	for {
		n, err := conn.Write([]byte(chunk))
		if err != nil {
			break
		}
		bytes += n
	}
	return
}