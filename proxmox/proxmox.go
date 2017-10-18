package proxmox

type ProxmoxClient struct {
	Host     string
	User     string
	Password string
	Auth     AuthInfo
	Node     string
}
