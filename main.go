package main

import (
	"flag"
	"fmt"
)

func main() {
	var host string
	var user string
	var password string
	var action string
	var realm string
	var node string
	var vmid string
	var cpus int
	var disk int
	var hostname string
	var ip_address string
	var memory int
	var swap int

	flag.StringVar(&host, "host", "", "Proxmox host")
	flag.StringVar(&user, "user", "root@pam", "Proxmox user")
	flag.StringVar(&password, "password", "", "Proxmox user password")
	flag.StringVar(&action, "action", "", "Proxmox api action")
	flag.StringVar(&realm, "realm", "pve", "Proxmox realm")
	flag.StringVar(&node, "node", "", "Proxmox node")
	flag.StringVar(&vmid, "vmid", "", "OpenVZ container VMID")
	flag.IntVar(&cpus, "cpus", 0, "Number of CPUs")
	flag.IntVar(&disk, "disk", 0, "Disk size")
	flag.StringVar(&hostname, "hostname", "", "Hostname")
	flag.StringVar(&ip_address, "ip-address", "", "IP Address")
	flag.IntVar(&memory, "memory", 0, "Memory")
	flag.IntVar(&swap, "swap", 0, "Swap")

	flag.Parse()

	proxmox := Proxmox{}
	proxmox.host = host
	proxmox.user = user
	proxmox.password = password
	proxmox.auth = proxmox.GetAuth()

	switch action {
	case "getDomains":
		domains := proxmox.GetDomains()
		PrintDataSlice(domains)
	case "getRealmConfig":
		domain := Domain{}
		domain.Realm = realm
		config := proxmox.GetRealmConfig(domain)
		PrintDataStruct(config)
	case "getClusterStatus":
		cluster := proxmox.GetClusterStatus()
		PrintDataSlice(cluster)
	case "getClusterTasks":
		clusterTasks := proxmox.GetClusterTasks()
		PrintDataSlice(clusterTasks)
	case "getClusterBackupSchedule":
		clusterBackupSchedule := proxmox.GetClusterBackupSchedule()
		PrintDataSlice(clusterBackupSchedule)
	case "getNodes":
		nodes := proxmox.GetNodes()
		PrintDataSlice(nodes)
	case "getContainers":
		proxmox.node = node
		containers := proxmox.GetContainers()
		PrintDataSlice(containers)
	case "getContainerConfig":
		var req = ContainerRequest{}
		req.node = node
		req.vmid = vmid
		containerConfig := proxmox.GetContainerConfig(req)
		PrintDataStruct(containerConfig)
	case "createContainer":
		req := &ContainerRequest{}
		req.node = node
		req.vmid = vmid
		if cpus > 0 {
			req.ContainerConfig.CPUs = cpus
		}
		if disk > 0 {
			req.ContainerConfig.Disk = disk
		}
		if hostname != "" {
			req.ContainerConfig.Hostname = hostname
		}
		if ip_address != "" {
			req.ContainerConfig.IP_Address = ip_address
		}
		if memory > 0 {
			req.ContainerConfig.Memory = memory
		}
		if swap > 0 {
			req.ContainerConfig.Swap = swap
		}
		proxmox.CreateContainer(req)
	default:
		fmt.Printf("Unsupported action: %s", action)
	}
}
