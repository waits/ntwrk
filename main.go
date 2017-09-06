// Command ntwrk is a tool for testing network performance.
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/waits/update"
)

const proto = "0.1"
const updateUrl = "https://api.github.com/repos/waits/ntwrk/releases/latest"

var version = update.Version{Major: 0, Minor: 0, Patch: 0}

func main() {
	var cmd string
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	} else {
		cmd = "help"
	}

	clientFlags := flag.NewFlagSet("client", flag.ExitOnError)
	host := clientFlags.String("host", "ntwrk.waits.io", "server to test against")
	port := 1600

	switch cmd {
	case "help":
		help()
	case "ip":
		clientFlags.Parse(os.Args[2:])
		whoami(*host, port)
	case "server":
		startServer(port)
	case "run":
		clientFlags.Parse(os.Args[2:])
		startClient(*host, port)
	case "update":
		update.Auto(version, updateUrl, update.CheckGithub)
	case "version":
		fmt.Printf("ntwrk %s %s/%s\n", version, runtime.GOOS, runtime.GOARCH)
	default:
		fmt.Printf("Unknown command '%v'.\n", cmd)
	}
}

func help() {
	cmds := []string{"help", "ip\t", "run\t", "server", "update", "version"}
	descriptions := []string{
		"Show this help message",
		"Print external IP address",
		"Run performance tests",
		"Start a test server",
		"Checks for and downloads an updated binary",
		"Print version number"}

	fmt.Print("usage: ntwrk <command> [arguments]\n\n")
	fmt.Print("commands:\n")
	for i, cmd := range cmds {
		fmt.Printf("    %v\t%v\n", cmd, descriptions[i])
	}
}
