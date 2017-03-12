package main

import (
	"fmt"
	"math"
	"net"
	"time"
)

const timeout = time.Duration(15) * time.Second

var suffixes = [...]string{"b", "kb", "Mb", "Gb", "Tb", "Pb", "Eb"}

// testContext holds a test function, action name, and address to connect to.
type testContext struct {
	Action string
	Fn     func(net.Conn, time.Duration) (int, error)
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
	fmt.Printf("%s bandwidth: %s\n", ctx.Action, format(bytes, elapsed))
}

// format returns the humanized bandwidth based on `bytes` and `seconds`.
func format(bytes int, seconds float64) string {
	raw := float64(bytes*8) / seconds
	if raw <= 10 {
		return fmt.Sprintf("%.2f b/s", raw)
	}

	exp := math.Floor(math.Log(raw) / math.Log(unit_base))
	suffix := suffixes[int(exp)]
	bandwidth := raw / math.Pow(unit_base, exp)
	return fmt.Sprintf("%.2f %s/s", bandwidth, suffix)
}
