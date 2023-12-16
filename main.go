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
	listPools(l)
	listDomains(l)

	if err := l.Disconnect(); err != nil {
		log.Fatalf("failed to disconnect: %v", err)
	}
}

func listNetworks(l *libvirt.Libvirt) {
	// https://libvirt.org/html/libvirt-libvirt-network.html#virConnectListAllNetworks
	nets, _, err := l.ConnectListAllNetworks(1, libvirt.ConnectListNetworksActive|libvirt.ConnectListNetworksInactive)
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

func listPools(l *libvirt.Libvirt) {
	// https://libvirt.org/html/libvirt-libvirt-storage.html#virConnectListAllStoragePools
	pools, _, err := l.ConnectListAllStoragePools(1, libvirt.ConnectListStoragePoolsActive|libvirt.ConnectListStoragePoolsInactive)
	if err != nil {
		log.Fatalf("failed to retrieve storage pools: %v", err)
	}

	fmt.Println("POOLS:")
	fmt.Println("\tName\t\tUUID")
	fmt.Printf("--------------------------------\n")
	for _, n := range pools {
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
