package main

import (
	"flag"
	"github.com/thirdwavellc/go-proxmox/proxmox"
	"log"
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
	comment         string
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
	flag.StringVar(&options.comment, "comment", "", "Comment")
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

	config, err := proxmox.ReadProxmoxConfig(options.configFile)

	if err != nil {
		proxmox.PrintError(err)
	}

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

	client := proxmox.ProxmoxClient{}
	client.Host = options.host
	client.User = options.user
	client.Password = options.password
	ticketReq := &proxmox.TicketRequest{
		Username: options.user,
		Password: options.password,
	}
	client.Auth, err = client.GetAuth(ticketReq)

	if err != nil {
		proxmox.PrintError(err)
	}

	switch options.action {
	case "getDomains":
		domains, err := client.GetDomains()

		if err != nil {
			proxmox.PrintError(err)
		}

		proxmox.PrintDataSlice(domains)
	case "getRealmConfig":
		request := &proxmox.RealmConfigRequest{
			Realm: options.realm,
		}
		config, err := client.GetRealmConfig(request)

		if err != nil {
			proxmox.PrintError(err)
		}

		proxmox.PrintDataStruct(config)
	case "getGroups":
		groups, err := client.GetGroups()

		if err != nil {
			proxmox.PrintError(err)
		}

		proxmox.PrintDataSlice(groups)
	case "getGroupConfig":
		request := &proxmox.GroupConfigRequest{
			GroupId: options.group_id,
		}
		config, err := client.GetGroupConfig(request)

		if err != nil {
			proxmox.PrintError(err)
		}

		proxmox.PrintDataStruct(config)
	case "createGroup":
		request := &proxmox.NewGroupRequest{
			GroupId: options.group_id,
			Comment: options.comment,
		}
		_, err := client.CreateGroup(request)

		if err != nil {
			proxmox.PrintError(err)
		}
	case "getRoles":
		roles, err := client.GetRoles()

		if err != nil {
			proxmox.PrintError(err)
		}

		proxmox.PrintDataSlice(roles)
	case "getRoleConfig":
		request := &proxmox.RoleConfigRequest{
			RoleId: options.role_id,
		}
		config, err := client.GetRoleConfig(request)

		if err != nil {
			proxmox.PrintError(err)
		}

		proxmox.PrintDataStruct(config)
	case "getClusterStatus":
		cluster, err := client.GetClusterStatus()

		if err != nil {
			proxmox.PrintError(err)
		}

		proxmox.PrintDataSlice(cluster)
	case "getClusterTasks":
		clusterTasks, err := client.GetClusterTasks()

		if err != nil {
			proxmox.PrintError(err)
		}

		proxmox.PrintDataSlice(clusterTasks)
	case "getClusterBackupSchedule":
		clusterBackupSchedule, err := client.GetClusterBackupSchedule()

		if err != nil {
			proxmox.PrintError(err)
		}

		proxmox.PrintDataSlice(clusterBackupSchedule)
	case "getNodes":
		nodes, err := client.GetNodes()

		if err != nil {
			proxmox.PrintError(err)
		}

		proxmox.PrintDataSlice(nodes)
	case "getNodeTaskStatus":
		request := &proxmox.NodeTaskStatusRequest{}
		request.Node = options.node
		request.UPID = options.upid
		status, err := client.GetNodeTaskStatus(request)

		if err != nil {
			proxmox.PrintError(err)
		}

		proxmox.PrintDataStruct(status)
	case "getContainers":
		request := &proxmox.ContainerRequest{
			Node: options.node,
		}
		containers, err := client.GetContainers(request)

		if err != nil {
			proxmox.PrintError(err)
		}

		proxmox.PrintDataSlice(containers)
	case "getContainerConfig":
		request := &proxmox.ContainerConfigRequest{
			Node: options.node,
			VMID: options.vmid,
		}
		containerConfig, err := client.GetContainerConfig(request)

		if err != nil {
			proxmox.PrintError(err)
		}

		proxmox.PrintDataStruct(containerConfig)
	case "getContainerStatus":
		request := &proxmox.ContainerStatusRequest{
			Node: options.node,
			VMID: options.vmid,
		}
		containerStatus, err := client.GetContainerStatus(request)

		if err != nil {
			proxmox.PrintError(err)
		}

		proxmox.PrintDataStruct(containerStatus)
	case "createContainer":
		req := &proxmox.NewContainerRequest{}
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
		upid, err := client.CreateContainer(req)

		if err != nil {
			proxmox.PrintError(err)
		}

		statusRequest := &proxmox.NodeTaskStatusRequest{}
		statusRequest.Node = options.node
		statusRequest.UPID = upid
		task, err := client.CheckNodeTaskStatus(statusRequest)

		if err != nil {
			proxmox.PrintError(err)
		}

		if task.ExitStatus == "OK" {
			log.Println("Container successfully created!")
		} else {
			log.Printf("Exit Status: %s", task.ExitStatus)
		}
	case "updateContainer":
		req := &proxmox.ExistingContainerRequest{}
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
		resp, err := client.UpdateContainer(req)

		if err != nil {
			proxmox.PrintError(err)
		}

		// TODO: handle response
		log.Printf(resp)
	case "deleteContainer":
		request := &proxmox.ExistingContainerRequest{}
		request.Node = options.node
		request.VMID = options.vmid
		log.Printf("Deleting container")
		upid, err := client.DeleteContainer(request)

		if err != nil {
			proxmox.PrintError(err)
		}

		statusRequest := &proxmox.NodeTaskStatusRequest{}
		statusRequest.Node = options.node
		statusRequest.UPID = upid
		task, err := client.CheckNodeTaskStatus(statusRequest)

		if err != nil {
			proxmox.PrintError(err)
		}

		if task.ExitStatus == "OK" {
			log.Println("Container successfully deleted!")
		} else {
			log.Printf("Exit Status: %s", task.ExitStatus)
		}
	case "startContainer":
		request := &proxmox.ExistingContainerRequest{}
		request.Node = options.node
		request.VMID = options.vmid
		log.Printf("Starting container")
		upid, err := client.StartContainer(request)

		if err != nil {
			proxmox.PrintError(err)
		}

		statusRequest := &proxmox.NodeTaskStatusRequest{}
		statusRequest.Node = options.node
		statusRequest.UPID = upid
		task, err := client.CheckNodeTaskStatus(statusRequest)

		if err != nil {
			proxmox.PrintError(err)
		}

		if task.ExitStatus == "OK" {
			log.Println("Container successfully started!")
		} else {
			log.Printf("Exit Status: %s", task.ExitStatus)
		}
	case "shutdownContainer":
		request := &proxmox.ExistingContainerRequest{}
		request.Node = options.node
		request.VMID = options.vmid
		log.Printf("Shutting down container")
		upid, err := client.ShutdownContainer(request)

		if err != nil {
			proxmox.PrintError(err)
		}

		statusRequest := &proxmox.NodeTaskStatusRequest{}
		statusRequest.Node = options.node
		statusRequest.UPID = upid
		task, err := client.CheckNodeTaskStatus(statusRequest)

		if err != nil {
			proxmox.PrintError(err)
		}

		if task.ExitStatus == "OK" {
			log.Println("Container successfully shutdown!")
		} else {
			log.Printf("Exit Status: %s", task.ExitStatus)
		}
	case "getNodeStorage":
		request := &proxmox.NodeStorageRequest{
			Node: options.node,
		}
		storage, err := client.GetNodeStorage(request)

		if err != nil {
			proxmox.PrintError(err)
		}

		proxmox.PrintDataSlice(storage)
	case "getNodeStorageContent":
		request := &proxmox.NodeStorageContentRequest{
			Node:    options.node,
			Storage: options.storage,
		}
		content, err := client.GetNodeStorageContent(request)

		if err != nil {
			proxmox.PrintError(err)
		}

		proxmox.PrintDataSlice(content)
	default:
		log.Printf("Unsupported action: %s", options.action)
		os.Exit(1)
	}
}
