package main

import (
	"io"
	"io/ioutil"
	"net"
	"strings"
	"time"
)

const sampleData = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

// download reads data from `conn` and returns the number of bytes read.
func download(conn net.Conn, timeout time.Duration) (bytes int64, err error) {
	if timeout > 0 {
		conn.SetDeadline(time.Now().Add(timeout))
	}

	bytes, _ = io.Copy(ioutil.Discard, conn)
	return
}

// upload writes data to `conn` and returns the number of bytes written.
func upload(conn net.Conn, timeout time.Duration) (bytes int64, err error) {
	if timeout > 0 {
		conn.SetDeadline(time.Now().Add(timeout))
	}

	chunk := strings.Repeat(sampleData, 256)
	for {
		n, err := conn.Write([]byte(chunk))
		if err != nil {
			break
		}
		bytes += int64(n)
	}
	return
}
