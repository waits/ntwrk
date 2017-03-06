package main

import (
	"fmt"
	"math"
	"net"
	"time"
)

var SUFFIXES = [...]string{"B", "kB", "MB", "GB", "TB", "PB", "EB"}

// testContext holds a test function, action name, and address to connect to.
type testContext struct {
	Action string
	Fn     func(net.Conn, int) (int, error)
	Addr   string
}

// startServer starts a network test client on `addr`.
func startClient(addr string) {
	perform(testContext{"download", download, addr})
	perform(testContext{"upload", upload, addr})
}

// perform runs a network test and prints the recorded bandwidth.
func perform(ctx testContext) {
	conn, err := net.Dial("tcp", ctx.Addr)
	if err != nil {
		panic(err)
	}

	t := time.Now()
	conn.Write([]byte(":" + ctx.Action + "\n"))
	bytes, err := ctx.Fn(conn, MAX)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}
	elapsed := time.Since(t).Seconds()
	fmt.Printf("%s bandwidth: %s\n", ctx.Action, format(bytes, elapsed))
}

// format returns the humanized bandwidth based on `bytes` and `seconds`.
func format(bytes int, seconds float64) string {
	raw := float64(bytes) / seconds
	if raw <= 10 {
		return fmt.Sprintf("%.2f B/s", raw)
	}

	exp := math.Floor(math.Log(raw) / math.Log(BASE))
	suffix := SUFFIXES[int(exp)]
	bandwidth := raw / math.Pow(BASE, exp)
	return fmt.Sprintf("%.2f %s/s", bandwidth, suffix)
}
