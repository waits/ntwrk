package main

import (
	"fmt"
	"net"
	"time"
)

const pingCount = 10
const protoFmt = "ntwrk%s :%s\r\n"
const timeout = time.Duration(15) * time.Second

// testContext holds a test function, action name, and address to connect to.
type testContext struct {
	Action string
	Fn     func(net.Conn, time.Duration) (int64, error)
	Addr   string
}

// startClient starts the network test suite.
func startClient(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)
	ping(addr)
	perform(testContext{"download", download, addr})
	perform(testContext{"upload", upload, addr})
}

// perform runs a network test and prints the recorded bandwidth.
func perform(ctx testContext) {
	conn, err := openConn(ctx.Addr, ctx.Action)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}

	since := time.Now()
	ticker := time.NewTicker(time.Millisecond * 150)
	go func() {
		for t := range ticker.C {
			elapsed := t.Sub(since)
			progress := formatProgress(elapsed, timeout)
			fmt.Printf("\r %9s: %s", ctx.Action, progress)
		}
	}()
	defer ticker.Stop()

	bytes, err := ctx.Fn(conn, timeout)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}

	elapsed := time.Since(since).Seconds()
	fmt.Printf(" %s\n", formatBytes(bytes, elapsed))
}

// ping performs a network latency test.
func ping(addr string) {
	conn, err := openConn(addr, "echo")
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}
	resp := make([]byte, 6)
	since := time.Now()
	for i := 0; i < pingCount; i++ {
		conn.Write([]byte("echo\r\n"))
		conn.Read(resp)
		if string(resp) != "echo\r\n" {
			fmt.Println("error: invalid echo reply")
			return
		}
	}
	elapsed := time.Since(since) / pingCount
	fmt.Printf("\r %9s: %s\n", "latency", elapsed)
}

// whoami requests the client's external IP address from `host` and prints it.
func whoami(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)
	resp := make([]byte, 40)

	conn, err := openConn(addr, "whoami")
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}

	conn.Read(resp)
	fmt.Print(string(resp))
}

// openConn opens a connection to `host` and writes a formatted message to it.
func openConn(host string, action string) (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", host)
	if err != nil {
		return
	}

	fmt.Fprintf(conn, protoFmt, proto, action)
	return
}
