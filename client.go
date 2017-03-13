package main

import (
	"fmt"
	"net"
	"time"
)

const timeout = time.Duration(15) * time.Second

var suffixes = [...]string{"b", "kb", "Mb", "Gb", "Tb", "Pb", "Eb"}

// testContext holds a test function, action name, and address to connect to.
type testContext struct {
	Action string
	Fn     func(net.Conn, time.Duration) (int64, error)
	Addr   string
}

// startClient starts the network test suite.
func startClient(host string) {
	perform(testContext{"download", download, host})
	perform(testContext{"upload", upload, host})
}

// perform runs a network test and prints the recorded bandwidth.
func perform(ctx testContext) {
	conn, err := net.Dial("tcp", ctx.Addr)
	if err != nil {
		panic(err)
	}

	t := time.Now()
	conn.Write([]byte(":" + ctx.Action + "\n"))
	bytes, err := ctx.Fn(conn, timeout)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}

	elapsed := time.Since(t).Seconds()
	fmt.Printf("%s bandwidth: %s\n", ctx.Action, formatBytes(bytes, elapsed))
}

// whoami requests the client's external IP address from `host` and prints it.
func whoami(host string) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		panic(err)
	}
	conn.Write([]byte(":whoami\n"))
	resp := make([]byte, 40)
	conn.Read(resp)
	fmt.Print(string(resp))
}
