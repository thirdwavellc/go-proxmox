package proxmox

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"net/url"
)

type ContainerList struct {
	Data []Container
}

type Container struct {
	CPU       float64 `json:"cpu"`
	CPUs      int     `json:"cpus"`
	Disk      int     `json:"disk"`
	DiskRead  int     `json:"diskread"`
	DiskWrite int     `json:"diskwrite"`
	Lock      string  `json:"lock"`
	MaxDisk   int     `json:"maxdisk"`
	MaxMem    int     `json:"maxmem"`
	MaxSwap   int     `json:"maxswap"`
	Mem       int     `json:"mem"`
	Name      string  `json:"name"`
	NetIn     int     `json:"netin"`
	NetOut    int     `json:"netout"`
	Status    string  `json:"status"`
	Swap      int     `json:"swap"`
	Template  string  `json:"template"`
	Type      string  `json:"type"`
	Uptime    int     `json:"uptime"`
	VMID      string  `json:"vmid"`
}

type ContainerRequest struct {
	Node string `url:"node"`
}

func (p ProxmoxClient) GetContainers(req *ContainerRequest) ([]Container, error) {
	endpoint_url := "/api2/json/nodes/" + req.Node + "/lxc"
	payload := url.Values{}
	body, err := p.GetContent(endpoint_url, payload)

	if err != nil {
		return nil, err
	}

	var containers ContainerList
	json.Unmarshal(body, &containers)
	return containers.Data, nil
}

type ContainerConfigList struct {
	Data ContainerConfig
}

type ContainerConfig struct {
	Arch         string
	Cores        int
	Digest       string
	Hostname     string
	Memory       int
	Net0         string
	OsType       string `json:"ostype"`
	RootFs       string `json:"rootfs"`
	Swap         int
	SearchDomain string `json:"searchdomain"`
	Nameserver   string `json:"nameserver"`
}

type NewContainerRequest struct {
	Node          string `json:"node" url:"-"`
	OsTemplate    string `json:"ostemplate" url:"ostemplate,omitempty"`
	VMID          string `json:"vmid" url:"vmid"`
	Net0          string `json:"net0" url:"net0,omitempty"`
	Storage       string `json:"storage" url:"storage,omitempty"`
	RootFs        string `json:"rootfs" url:"rootfs,omitempty"`
	Cores         int    `json:"cores" url:"cores,omitempty"`
	Memory        int    `json:"memory" url:"memory,omitempty"`
	Swap          int    `json:"swap" url:"swap,omitempty"`
	Hostname      string `json:"hostname" url:"hostname,omitempty"`
	OnBoot        int    `json:"onboot" url:"onboot,omitempty"`
	Password      string `json:"password" url:"password,omitempty"`
	SshPublicKeys string `json:"ssh-public-keys" url:"ssh-public-keys,omitempty"`
	Unprivileged  int    `json:"unprivileged" url:"unprivileged,omitempty"`
	SearchDomain  string `json:"searchdomain" url:"searchdomain,omitempty"`
	Nameserver    string `json:"nameserver" url:"nameserver,omitempty"`
}

type ExistingContainerRequest struct {
	Node          string `json:"node" url:"-"`
	OsTemplate    string `json:"ostemplate" url:"ostemplate,omitempty"`
	VMID          string `json:"vmid" url:"-"`
	Net0          string `json:"net0" url:"net0,omitempty"`
	Storage       string `json:"storage" url:"storage,omitempty"`
	RootFs        string `json:"rootfs" url:"rootfs,omitempty"`
	Cores         int    `json:"cores" url:"cores,omitempty"`
	Memory        int    `json:"memory" url:"memory,omitempty"`
	Swap          int    `json:"swap" url:"swap,omitempty"`
	Hostname      string `json:"hostname" url:"hostname,omitempty"`
	OnBoot        int    `json:"onboot" url:"onboot,omitempty"`
	Password      string `json:"password" url:"password,omitempty"`
	SshPublicKeys string `json:"ssh-public-keys" url:"ssh-public-keys,omitempty"`
	Unprivileged  int    `json:"unprivileged" url:"unprivileged,omitempty"`
	SearchDomain  string `json:"searchdomain" url:"searchdomain,omitempty"`
	Nameserver    string `json: "nameserver" url:"nameserver,omitempty"`
}

type ContainerConfigRequest struct {
	Node string `url:"node"`
	VMID string `url:"vmid"`
}

func (p ProxmoxClient) GetContainerConfig(req *ContainerConfigRequest) (ContainerConfig, error) {
	endpoint_url := "/api2/json/nodes/" + req.Node + "/lxc/" + req.VMID + "/config"
	payload := url.Values{}
	body, err := p.GetContent(endpoint_url, payload)

	if err != nil {
		return ContainerConfig{}, err
	}

	var containerConfig ContainerConfigList
	json.Unmarshal(body, &containerConfig)
	return containerConfig.Data, nil
}

type ContainerStatusRequest struct {
	Node string `url:"node"`
	VMID string `url:"vmid"`
}

type ContainerStatusWrapper struct {
	Data ContainerStatus `json:"data"`
}

type ContainerStatus struct {
	CPU       int    `json:"cpu"`
	CPUs      int    `json:"cpus"`
	Disk      int    `json:"disk"`
	DiskRead  string `json:"diskread"`
	DiskWrite string `json:"diskwrite"`
	// HA
	Lock     string `json:"lock"`
	MaxDisk  int    `json:"maxdisk"`
	MaxMem   int    `json:"maxmem"`
	MaxSwap  int    `json:"maxswap"`
	Mem      int    `json:"mem"`
	Name     string `json:"name"`
	NetIn    int    `json:"netin"`
	NetOut   int    `json:"netout"`
	PID      string `json:"pid"`
	Status   string `json:"status"`
	Swap     int    `json:"swap"`
	Template string `json:"template"`
	Type     string `json:"type"`
	Uptime   int    `json:"uptime"`
}

func (p ProxmoxClient) GetContainerStatus(req *ContainerStatusRequest) (ContainerStatus, error) {
	endpoint_url := "/api2/json/nodes/" + req.Node + "/lxc/" + req.VMID + "/status/current"
	payload := url.Values{}
	body, err := p.GetContent(endpoint_url, payload)

	if err != nil {
		return ContainerStatus{}, err
	}

	var containerStatus ContainerStatusWrapper
	json.Unmarshal(body, &containerStatus)
	return containerStatus.Data, nil
}

type ContainerResponse struct {
	Data string
}

func (p ProxmoxClient) CreateContainer(req *NewContainerRequest) (string, error) {
	endpoint_url := "/api2/json/nodes/" + req.Node + "/lxc"
	payload, _ := query.Values(req)
	body, err := p.PostContent(endpoint_url, payload)

	if err != nil {
		return "", err
	}

	var response ContainerResponse
	json.Unmarshal(body, &response)
	return response.Data, nil
}

func (p ProxmoxClient) UpdateContainer(req *ExistingContainerRequest) (string, error) {
	endpoint_url := "/api2/json/nodes/" + req.Node + "/lxc/" + req.VMID + "/config"
	payload, _ := query.Values(req)
	body, err := p.PutContent(endpoint_url, payload)

	if err != nil {
		return "", err
	}

	var response ContainerResponse
	json.Unmarshal(body, &response)
	return response.Data, nil
}

func (p ProxmoxClient) DeleteContainer(req *ExistingContainerRequest) (string, error) {
	endpoint_url := "/api2/json/nodes/" + req.Node + "/lxc/" + req.VMID
	payload := url.Values{}
	body, err := p.DeleteContent(endpoint_url, payload)

	if err != nil {
		return "", err
	}

	var response ContainerResponse
	json.Unmarshal(body, &response)
	return response.Data, nil
}

func (p ProxmoxClient) StartContainer(req *ExistingContainerRequest) (string, error) {
	endpoint_url := "/api2/json/nodes/" + req.Node + "/lxc/" + req.VMID + "/status/start"
	payload := url.Values{}
	body, err := p.PostContent(endpoint_url, payload)

	if err != nil {
		return "", err
	}

	var response ContainerResponse
	json.Unmarshal(body, &response)
	return response.Data, nil
}

func (p ProxmoxClient) ShutdownContainer(req *ExistingContainerRequest) (string, error) {
	endpoint_url := "/api2/json/nodes/" + req.Node + "/lxc/" + req.VMID + "/status/shutdown"
	payload := url.Values{}
	body, err := p.PostContent(endpoint_url, payload)

	if err != nil {
		return "", err
	}

	var response ContainerResponse
	json.Unmarshal(body, &response)
	return response.Data, nil
}
