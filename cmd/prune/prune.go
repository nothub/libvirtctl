package prune

import (
	"log"

	"github.com/digitalocean/go-libvirt"
)

func pruneAll(l *libvirt.Libvirt) {
	pruneDomains(l)
	prunePools(l)
	pruneNetworks(l)
}

func pruneDomains(l *libvirt.Libvirt) {
	disableDomains(l)
	removeDomains(l)
}

func prunePools(l *libvirt.Libvirt) {
	disablePools(l)
	removePools(l)
}

func pruneNetworks(l *libvirt.Libvirt) {
	disableNetworks(l)
	removeNetworks(l)
}

func disableDomains(l *libvirt.Libvirt) {
	// https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectListAllDomains
	domains, _, err := l.ConnectListAllDomains(1, libvirt.ConnectListDomainsActive)
	if err != nil {
		log.Fatalf("failed to retrieve active domains; %v", err)
	}

	for _, d := range domains {
		log.Printf("disabling domain %d %s (%x)\n", d.ID, d.Name, d.UUID)
		// https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainDestroy
		err := l.DomainDestroyFlags(d, libvirt.DomainDestroyGraceful)
		if err != nil {
			log.Fatalf("failed to disable domain; %v", err)
		}
	}
}

func removeDomains(l *libvirt.Libvirt) {
	// https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectListAllDomains
	domains, _, err := l.ConnectListAllDomains(1, libvirt.ConnectListDomainsInactive)
	if err != nil {
		log.Fatalf("failed to retrieve inactive domains; %v", err)
	}

	for _, d := range domains {
		log.Printf("removing domain %d %s (%x)\n", d.ID, d.Name, d.UUID)
		// https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainDestroy
		err := l.DomainUndefineFlags(d, libvirt.DomainUndefineManagedSave|
			libvirt.DomainUndefineSnapshotsMetadata|
			libvirt.DomainUndefineCheckpointsMetadata|
			libvirt.DomainUndefineNvram)
		if err != nil {
			log.Fatalf("failed to remove domain; %v", err)
		}
	}
}

func disablePools(l *libvirt.Libvirt) {
	// https://libvirt.org/html/libvirt-libvirt-storage.html#virConnectListAllStoragePools
	pools, _, err := l.ConnectListAllStoragePools(1, libvirt.ConnectListStoragePoolsActive)
	if err != nil {
		log.Fatalf("failed to retrieve active pools; %v", err)
	}

	for _, p := range pools {
		log.Printf("disabling pool %s (%x)\n", p.Name, p.UUID)
		// https://libvirt.org/html/libvirt-libvirt-storage.html#virStoragePoolDestroy
		err := l.StoragePoolDestroy(p)
		if err != nil {
			log.Fatalf("failed to disable pool; %v", err)
		}
	}
}

func removePools(l *libvirt.Libvirt) {
	// https://libvirt.org/html/libvirt-libvirt-storage.html#virConnectListAllStoragePools
	pools, _, err := l.ConnectListAllStoragePools(1, libvirt.ConnectListStoragePoolsInactive)
	if err != nil {
		log.Fatalf("failed to retrieve inactive pools; %v", err)
	}

	for _, p := range pools {
		log.Printf("removing pool %s (%x)\n", p.Name, p.UUID)
		// https://libvirt.org/html/libvirt-libvirt-storage.html#virStoragePoolDestroy
		err := l.StoragePoolUndefine(p)
		if err != nil {
			log.Fatalf("failed to remove pool; %v", err)
		}
	}
}

func disableNetworks(l *libvirt.Libvirt) {
	// https://libvirt.org/html/libvirt-libvirt-network.html#virConnectListAllNetworks
	networks, _, err := l.ConnectListAllNetworks(1, libvirt.ConnectListNetworksActive)
	if err != nil {
		log.Fatalf("failed to retrieve active networks; %v", err)
	}

	for _, n := range networks {
		log.Printf("disabling network %s (%x)\n", n.Name, n.UUID)
		// https://libvirt.org/html/libvirt-libvirt-network.html#virNetworkDestroy
		err := l.NetworkDestroy(n)
		if err != nil {
			log.Fatalf("failed to disable network; %v", err)
		}
	}
}

func removeNetworks(l *libvirt.Libvirt) {
	// https://libvirt.org/html/libvirt-libvirt-network.html#virConnectListAllNetworks
	networks, _, err := l.ConnectListAllNetworks(1, libvirt.ConnectListNetworksInactive)
	if err != nil {
		log.Fatalf("failed to retrieve inactive networks; %v", err)
	}

	for _, n := range networks {
		log.Printf("removing network %s (%x)\n", n.Name, n.UUID)
		// https://libvirt.org/html/libvirt-libvirt-network.html#virNetworkDestroy
		err := l.NetworkUndefine(n)
		if err != nil {
			log.Fatalf("failed to remove network; %v", err)
		}
	}
}
