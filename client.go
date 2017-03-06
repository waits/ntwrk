package main

import (
	"fmt"
	"net"
	"time"
)

// startServer starts a network test client on `addr`.
func startClient(addr string) {
	perform(testContext{download, "DOWNLOAD", addr})
	perform(testContext{upload, "UPLOAD", addr})
}

// perform runs a test function and reports the time taken.
func perform(ctx testContext) {
	conn, err := net.Dial("tcp", ctx.Addr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected to %s\n", ctx.Addr)

	t := time.Now()
	conn.Write([]byte(ctx.Action + "\n"))
	bytes, err := ctx.Fn(conn, MAX)
	if err != nil {
		fmt.Printf("Test failed: %s", err.Error())
		return
	}
	elapsed := time.Since(t)
	fmt.Printf("Processed %d bytes in %v\n", bytes, elapsed)
}
