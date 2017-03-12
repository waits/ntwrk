// Command ntwrk is a tool for testing network performance.
package main

import (
	"flag"
	"fmt"
	"os"
)

const unit_base = 1024
const version = "0.1.0-alpha"

func main() {
	var cmd string
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	} else {
		cmd = "help"
	}

	serverFlags := flag.NewFlagSet("server", flag.ExitOnError)
	port := serverFlags.Int("port", 1600, "port to listen on")

	clientFlags := flag.NewFlagSet("client", flag.ExitOnError)
	host := clientFlags.String("host", "ntwrk.waits.io:1600", "server to test against")

	switch cmd {
	case "help":
		help()
	case "server":
		serverFlags.Parse(os.Args[2:])
		startServer(*port)
	case "test":
		clientFlags.Parse(os.Args[2:])
		startClient(*host)
	case "version":
		fmt.Printf("ntwrk version %s\n", version)
	default:
		fmt.Printf("Unknown command '%v'.\n", cmd)
	}
}

func help() {
	cmds := [4]string{"help", "server", "test", "version"}
	descriptions := [4]string{
		"Show this help message",
		"Start a test server",
		"Run performance tests",
		"Print version number"}

	fmt.Print("usage: ntwrk <command> [arguments]\n\n")
	fmt.Print("commands:\n")
	for i, cmd := range cmds {
		fmt.Printf("    %v\t%v\n", cmd, descriptions[i])
	}
}
