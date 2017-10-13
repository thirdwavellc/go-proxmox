package main

import (
	"flag"
	"fmt"
	"os"
)

type Options struct {
	configFile      string
	host            string
	user            string
	password        string
	action          string
	realm           string
	group_id        string
	role_id         string
	node            string
	upid            string
	vmid            string
	os_template     string
	cpus            int
	disk            int
	hostname        string
	ip_address      string
	memory          int
	swap            int
	datastore       string
	net0            string
	storage         string
	root_fs         string
	cores           int
	on_boot         int
	root_password   string
	ssh_public_keys string
	unprivileged    int
}

func getOpts() Options {
	options := Options{}

	flag.StringVar(&options.configFile, "config", "", "Proxmox config file")
	flag.StringVar(&options.host, "host", "", "Proxmox host")
	flag.StringVar(&options.user, "user", "root@pam", "Proxmox user")
	flag.StringVar(&options.password, "password", "", "Proxmox user password")
	flag.StringVar(&options.action, "action", "", "Proxmox api action")
	flag.StringVar(&options.realm, "realm", "pve", "Proxmox realm")
	flag.StringVar(&options.group_id, "group-id", "", "Proxmox group")
	flag.StringVar(&options.role_id, "role-id", "", "Proxmox role")
	flag.StringVar(&options.node, "node", "", "Proxmox node")
	flag.StringVar(&options.upid, "upid", "", "Proxmox task UPID")
	flag.StringVar(&options.vmid, "vmid", "", "OpenVZ container VMID")
	flag.StringVar(&options.os_template, "os-template", "", "OpenVZ container template")
	flag.IntVar(&options.cpus, "cpus", 0, "Number of CPUs")
	flag.IntVar(&options.disk, "disk", 0, "Disk size")
	flag.StringVar(&options.hostname, "hostname", "", "Hostname")
	flag.StringVar(&options.ip_address, "ip-address", "", "IP Address")
	flag.IntVar(&options.memory, "memory", 512, "Memory")
	flag.IntVar(&options.swap, "swap", 512, "Swap")
	flag.StringVar(&options.datastore, "datastore", "", "Datastore identifier")
	flag.StringVar(&options.net0, "net0", "", "Network interface 0 config")
	flag.StringVar(&options.storage, "storage", "", "Storage identifier")
	flag.StringVar(&options.root_fs, "root-fs", "", "Root Filesystem")
	flag.IntVar(&options.cores, "cores", 1, "CPU Cores")
	flag.IntVar(&options.on_boot, "on-boot", 0, "Startup on boot")
	flag.StringVar(&options.root_password, "root-password", "", "Root password")
	flag.StringVar(&options.ssh_public_keys, "ssh-public-keys", "", "SSH Public Keys")
	flag.IntVar(&options.unprivileged, "unprivileged", 0, "Unprivileged user")

	flag.Parse()

	return options
}

func main() {
	options := getOpts()

	config := ReadProxmoxConfig(options.configFile)

	if config.Host != "" && options.host == "" {
		options.host = config.Host
	}
	if config.User != "" && options.user == "" {
		options.user = config.User
	}
	if config.Password != "" && options.password == "" {
		options.password = config.Password
	}
	if config.DefaultNode != "" && options.node == "" {
		options.node = config.DefaultNode
	}

	proxmox := Proxmox{}
	proxmox.host = options.host
	proxmox.user = options.user
	proxmox.password = options.password
	proxmox.auth = proxmox.GetAuth()

	switch options.action {
	case "getDomains":
		domains := proxmox.GetDomains()
		PrintDataSlice(domains)
	case "getRealmConfig":
		domain := Domain{}
		domain.Realm = options.realm
		config := proxmox.GetRealmConfig(domain)
		PrintDataStruct(config)
	case "getGroups":
		groups := proxmox.GetGroups()
		PrintDataSlice(groups)
	case "getGroupConfig":
		var group Group
		group.GroupId = options.group_id
		config := proxmox.GetGroupConfig(group)
		PrintDataStruct(config)
	case "createGroup":
		var group Group
		group.GroupId = options.group_id
		proxmox.CreateGroup(group)
	case "getRoles":
		roles := proxmox.GetRoles()
		PrintDataSlice(roles)
	case "getRoleConfig":
		var role Role
		role.RoleId = options.role_id
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
	case "getNodeTaskStatus":
		request := NodeTaskStatusRequest{}
		request.Node = options.node
		request.UPID = options.upid
		status := proxmox.GetNodeTaskStatus(request)
		PrintDataStruct(status)
	case "getContainers":
		proxmox.node = options.node
		containers := proxmox.GetContainers()
		PrintDataSlice(containers)
	case "getContainerConfig":
		var req = ContainerRequest{}
		req.Node = options.node
		req.VMID = options.vmid
		containerConfig := proxmox.GetContainerConfig(req)
		PrintDataStruct(containerConfig)
	case "createContainer":
		req := &ContainerRequest{}
		req.Node = options.node
		req.VMID = options.vmid
		req.OsTemplate = options.os_template
		req.Net0 = options.net0
		req.Storage = options.storage
		req.RootFs = options.root_fs
		req.Cores = options.cores
		req.Memory = options.memory
		req.Swap = options.swap
		req.Hostname = options.hostname
		req.OnBoot = options.on_boot
		req.Password = options.root_password
		req.SshPublicKeys = options.ssh_public_keys
		req.Unprivileged = options.unprivileged
		upid := proxmox.CreateContainer(req)
		statusRequest := NodeTaskStatusRequest{}
		statusRequest.Node = options.node
		statusRequest.UPID = upid
		task := proxmox.CheckNodeTaskStatus(statusRequest)
		if task.ExitStatus == "OK" {
			fmt.Println("Container successfully created!")
		} else {
			fmt.Println("Exit Status: %s", task.ExitStatus)
		}
	case "deleteContainer":
		request := &ContainerRequest{}
		request.Node = options.node
		request.VMID = options.vmid
		fmt.Printf("Deleting container")
		upid := proxmox.DeleteContainer(request)
		statusRequest := NodeTaskStatusRequest{}
		statusRequest.Node = options.node
		statusRequest.UPID = upid
		task := proxmox.CheckNodeTaskStatus(statusRequest)
		if task.ExitStatus == "OK" {
			fmt.Println("Container successfully deleted!")
		} else {
			fmt.Println("Exit Status: %s", task.ExitStatus)
		}
	case "getNodeDatastores":
		datastores := proxmox.GetNodeDatastores(options.node)
		PrintDataSlice(datastores)
	case "getNodeDatastoreContent":
		content := proxmox.GetNodeDatastoreContent(options.node, options.datastore)
		PrintDataSlice(content)
	default:
		fmt.Printf("Unsupported action: %s", options.action)
		os.Exit(1)
	}
}
