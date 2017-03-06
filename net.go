package main

import "net"

const DATA = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
const UPLOAD_LIMIT = 1024 * 1024

// testContext holds a test function, action name, and address to connect to.
type testContext struct {
	Fn     func(net.Conn) (int, error)
	Action string
	Addr   string
}

// download reads data from `conn` and returns the number of bytes read.
func download(conn net.Conn) (bytes int, err error) {
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
func upload(conn net.Conn) (bytes int, err error) {
	var n int
	for bytes < UPLOAD_LIMIT {
		n, err = conn.Write([]byte(DATA))
		if err != nil {
			return
		}
		bytes += n
	}
	return
}
