// Command ntwrk is a tool for testing network performance.
package main

import (
	"flag"
	"fmt"
	"os"
)

const proto = "0.1"
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
	case "ip":
		clientFlags.Parse(os.Args[2:])
		whoami(*host)
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
	cmds := [5]string{"help", "ip\t", "server", "test", "version"}
	descriptions := [5]string{
		"Show this help message",
		"Print external IP address",
		"Start a test server",
		"Run performance tests",
		"Print version number"}

	fmt.Print("usage: ntwrk <command> [arguments]\n\n")
	fmt.Print("commands:\n")
	for i, cmd := range cmds {
		fmt.Printf("    %v\t%v\n", cmd, descriptions[i])
	}
}
