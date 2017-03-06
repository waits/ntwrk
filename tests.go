package main

import "net"

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
func upload(conn net.Conn) (int, error) {
	bytes := 0
	for i := 0; i < 1024; i++ {
		n, err := conn.Write([]byte(DATA))
		if err != nil {
			return 0, err
		}
		bytes += n
	}
	return bytes, nil
}
