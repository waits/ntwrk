package main

import "net"

const DATA = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

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
	for {
		n, err := conn.Write([]byte(DATA))
		if err != nil {
			break
		}
		bytes += n
	}
	return
}
