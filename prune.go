package main

import (
	"github.com/digitalocean/go-libvirt"
	"log"
)

func pruneAll(l *libvirt.Libvirt) {
	pruneDomains(l)
	prunePools(l)
	pruneNetworks(l)
}

func pruneDomains(l *libvirt.Libvirt) {
	// https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectListAllDomains
	domains, _, err := l.ConnectListAllDomains(1, libvirt.ConnectListDomainsActive|libvirt.ConnectListDomainsInactive)
	if err != nil {
		log.Fatalf("failed to retrieve domains; %v", err)
	}

	for _, d := range domains {
		log.Printf("removing domain %d %s (%x)\n", d.ID, d.Name, d.UUID)
		// https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainDestroy
		err := l.DomainDestroyFlags(d, libvirt.DomainDestroyGraceful)
		if err != nil {
			log.Fatalf("failed to remove domain; %v", err)
		}
	}
}

func prunePools(l *libvirt.Libvirt) {
	// https://libvirt.org/html/libvirt-libvirt-storage.html#virConnectListAllStoragePools
	pools, _, err := l.ConnectListAllStoragePools(1, libvirt.ConnectListStoragePoolsActive|libvirt.ConnectListStoragePoolsInactive)
	if err != nil {
		log.Fatalf("failed to retrieve pools; %v", err)
	}

	for _, p := range pools {
		log.Printf("removing pool %s (%x)\n", p.Name, p.UUID)
		// https://libvirt.org/html/libvirt-libvirt-storage.html#virStoragePoolDestroy
		err := l.StoragePoolDestroy(p)
		if err != nil {
			log.Fatalf("failed to remove pool; %v", err)
		}
	}
}

func pruneNetworks(l *libvirt.Libvirt) {
	// https://libvirt.org/html/libvirt-libvirt-network.html#virConnectListAllNetworks
	networks, _, err := l.ConnectListAllNetworks(1, libvirt.ConnectListNetworksActive|libvirt.ConnectListNetworksInactive)
	if err != nil {
		log.Fatalf("failed to retrieve networks; %v", err)
	}

	for _, n := range networks {
		log.Printf("removing network %s (%x)\n", n.Name, n.UUID)
		// https://libvirt.org/html/libvirt-libvirt-network.html#virNetworkDestroy
		err := l.NetworkDestroy(n)
		if err != nil {
			log.Fatalf("failed to remove network; %v", err)
		}
	}
}
