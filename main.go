package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
)

func PrintError(err error) {
	fmt.Println("There was an error...")
	fmt.Printf("Error: %v", err)
	os.Exit(1)
}

func PrintDataSlice(data interface{}) {
	d := reflect.ValueOf(data)

	for i := 0; i < d.Len(); i++ {
		dataItem := d.Index(i)
		typeOfT := dataItem.Type()

		for j := 0; j < dataItem.NumField(); j++ {
			f := dataItem.Field(j)
			fmt.Printf("%s: %v\n", typeOfT.Field(j).Name, f.Interface())
		}
		fmt.Printf("\n")
	}
}

func PrintDataStruct(data interface{}) {
	d := reflect.ValueOf(data)
	typeOfT := d.Type()

	for j := 0; j < d.NumField(); j++ {
		f := d.Field(j)
		fmt.Printf("%s: %v\n", typeOfT.Field(j).Name, f.Interface())
	}
	fmt.Printf("\n")
}

func main() {
	var host string
	var user string
	var password string
	var action string
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
	flag.StringVar(&node, "node", "", "Proxmox node")
	flag.StringVar(&vmid, "vmid", "", "OpenVZ container VMID")
	flag.IntVar(&cpus, "cpus", 0, "Number of CPUs")
	flag.IntVar(&disk, "disk", 0, "Disk size")
	flag.StringVar(&hostname, "hostname", "", "Hostname")
	flag.StringVar(&ip_address, "ip-address", "", "IP Address")
	flag.IntVar(&memory, "memory", 0, "Memory")
	flag.IntVar(&swap, "swap", 0, "Swap")

	flag.Parse()

	auth := GetAuth(host, user, password)

	switch action {
	case "getClusterStatus":
		cluster := GetClusterStatus(host, auth)
		PrintDataSlice(cluster)
	case "getClusterTasks":
		clusterTasks := GetClusterTasks(host, auth)
		PrintDataSlice(clusterTasks)
	case "getClusterBackupSchedule":
		clusterBackupSchedule := GetClusterBackupSchedule(host, auth)
		PrintDataSlice(clusterBackupSchedule)
	case "getNodes":
		nodes := GetNodes(host, auth)
		PrintDataSlice(nodes)
	case "getContainers":
		containers := GetContainers(host, node, auth)
		PrintDataSlice(containers)
	case "getContainerConfig":
		containerConfig := GetContainerConfig(host, node, vmid, auth)
		PrintDataStruct(containerConfig)
	case "createContainer":
		containerRequest := &ContainerRequest{}
		containerRequest.Node = node
		containerRequest.VMID = vmid
		if cpus > 0 {
			containerRequest.ContainerConfig.CPUs = cpus
		}
		if disk > 0 {
			containerRequest.ContainerConfig.Disk = disk
		}
		if hostname != "" {
			containerRequest.ContainerConfig.Hostname = hostname
		}
		if ip_address != "" {
			containerRequest.ContainerConfig.IP_Address = ip_address
		}
		if memory > 0 {
			containerRequest.ContainerConfig.Memory = memory
		}
		if swap > 0 {
			containerRequest.ContainerConfig.Swap = swap
		}
		CreateContainer(host, containerRequest, auth)
	}
}
