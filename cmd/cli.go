package main

import (
	"fmt"
	"log"
	"os"

	"hub.lol/libvirtctl/cmd/prune"
)

func init() {
	log.SetFlags(0)
}

func main() {
	var cmd string
	if len(os.Args) >= 2 {
		cmd = os.Args[1]
		os.Args = append(os.Args[:1], os.Args[2:]...)
	}
	switch cmd {

	case "prune":
		prune.Run(os.Args)

	case "help":
		fmt.Fprintln(os.Stdout, "TODO: USAGE HELP [CMD]")
		os.Exit(0)

	default:
		fmt.Fprintln(os.Stderr, "TODO: USAGE HELP")
		os.Exit(1)
	}
}
