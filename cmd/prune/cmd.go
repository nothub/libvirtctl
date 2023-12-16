package prune

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/digitalocean/go-libvirt"
)

func Run(args []string) {
	// systemSocket := "/var/run/libvirt/libvirt-sock"
	userSocket := "/run/user/1000/libvirt/libvirt-sock"
	c, err := net.DialTimeout("unix", userSocket, 3*time.Second)
	if err != nil {
		log.Fatalf("failed to dial libvirt socket; %v", err)
	}

	l := libvirt.New(c)
	if err := l.ConnectToURI(libvirt.QEMUSession); err != nil {
		log.Fatalf("failed to connect libvirt rpc; %v", err)
	}

	h, err := l.ConnectGetHostname()
	if err != nil {
		log.Fatalf("failed to retrieve libvirt hostname; %v", err)
	}
	fmt.Println("Connected to:", h)

	v, err := l.ConnectGetLibVersion()
	if err != nil {
		log.Fatalf("failed to retrieve libvirt version; %v", err)
	}
	fmt.Println("Version:", v)

	pruneAll(l)

	if err := l.Disconnect(); err != nil {
		log.Fatalf("failed to disconnect; %v", err)
	}
}
