package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/digitalocean/go-libvirt"
)

func main() {
	c, err := net.DialTimeout("unix", "/var/run/libvirt/libvirt-sock", 2*time.Second)
	if err != nil {
		log.Fatalf("failed to dial libvirt: %v", err)
	}

	l := libvirt.New(c)
	if err := l.Connect(); err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	h, err := l.ConnectGetHostname()
	if err != nil {
		log.Fatalf("failed to retrieve libvirt hostname: %v", err)
	}
	fmt.Println("Connected to:", h)

	v, err := l.ConnectGetLibVersion()
	if err != nil {
		log.Fatalf("failed to retrieve libvirt version: %v", err)
	}
	fmt.Println("Version:", v)

	listNetworks(l)
	listDomains(l)

	if err := l.Disconnect(); err != nil {
		log.Fatalf("failed to disconnect: %v", err)
	}
}

func listNetworks(l *libvirt.Libvirt) {
	nets, _, err := l.ConnectListAllNetworks(0, 0)
	if err != nil {
		log.Fatalf("failed to retrieve networks: %v", err)
	}

	fmt.Println("NETWORKS:")
	fmt.Println("\tName\t\tUUID")
	fmt.Printf("--------------------------------\n")
	for _, n := range nets {
		fmt.Printf("%s\t%x\n", n.Name, n.UUID)
	}
}

func listDomains(l *libvirt.Libvirt) {
	domains, err := l.Domains()
	if err != nil {
		log.Fatalf("failed to retrieve domains: %v", err)
	}

	fmt.Println("DOMAINS:")
	fmt.Println("ID\tName\t\tUUID")
	fmt.Printf("--------------------------------\n")
	for _, d := range domains {
		fmt.Printf("%d\t%s\t%x\n", d.ID, d.Name, d.UUID)
	}
}
