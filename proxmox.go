package main

type Proxmox struct {
	host         string
	user         string
	password     string
	auth         AuthInfo
	api_endpoint string
	node         string
	vmid         string
	cpus         int
	disk         int
	hostname     string
	ip_address   string
	memory       int
	swap         int
}
