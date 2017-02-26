// Command ntwrk is a tool for testing network performance.
package main

import "fmt"
import "os"

const VERSION = "0.1.0-alpha"

func main() {
	var cmd string
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	} else {
		cmd = "help"
	}

	switch cmd {
	case "help":
		help()
	case "server":
		fmt.Println("Not implemented yet :(")
	case "test":
		fmt.Println("Not implemented yet :(")
	case "version":
		fmt.Printf("ntwrk version %s\n", VERSION)
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
