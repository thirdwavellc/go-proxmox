package proxmox

type ProxmoxClient struct {
	Host       string
	User       string
	Password   string
	Auth       AuthInfo
	Node       string
	Vmid       string
	Cpus       int
	Disk       int
	Hostname   string
	Ip_address string
	Memory     int
	Swap       int
}
