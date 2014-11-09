package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var host string
	var user string
	var password string
	var action string
	var realm string
	var group_id string
	var role_id string
	var node string
	var vmid string
	var ostemplate string
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
	flag.StringVar(&group_id, "group-id", "", "Proxmox group")
	flag.StringVar(&role_id, "role-id", "", "Proxmox role")
	flag.StringVar(&node, "node", "", "Proxmox node")
	flag.StringVar(&vmid, "vmid", "", "OpenVZ container VMID")
	flag.StringVar(&ostemplate, "ostemplate", "", "OpenVZ container template")
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
	case "getGroups":
		groups := proxmox.GetGroups()
		PrintDataSlice(groups)
	case "getGroupConfig":
		var group Group
		group.GroupId = group_id
		config := proxmox.GetGroupConfig(group)
		PrintDataStruct(config)
	case "createGroup":
		var group Group
		group.GroupId = group_id
		proxmox.CreateGroup(group)
	case "getRoles":
		roles := proxmox.GetRoles()
		PrintDataSlice(roles)
	case "getRoleConfig":
		var role Role
		role.RoleId = role_id
		config := proxmox.GetRoleConfig(role)
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
		req.Node = node
		req.VMID = vmid
		containerConfig := proxmox.GetContainerConfig(req)
		PrintDataStruct(containerConfig)
	case "createContainer":
		req := &ContainerRequest{}
		req.Node = node
		req.VMID = vmid
		req.OsTemplate = ostemplate
		proxmox.CreateContainer(req)
	default:
		fmt.Printf("Unsupported action: %s", action)
		os.Exit(1)
	}
}
