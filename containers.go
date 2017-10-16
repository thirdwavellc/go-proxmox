package main

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"net/url"
)

type ContainerList struct {
	Data []Container
}

type Container struct {
	MaxSwap   int
	Disk      int
	IP        string
	Status    string
	Netout    int
	MaxDisk   int
	MaxMem    int
	Uptime    int
	Swap      int
	VMID      string
	NProc     string
	DiskRead  int
	CPU       float64
	NetIn     int
	Name      string
	FailCnt   int
	DiskWrite int
	Mem       int
	Type      string
	CPUs      int
}

func (p Proxmox) GetContainers() ([]Container, error) {
	endpoint_url := "/api2/json/nodes/" + p.node + "/lxc"
	body, err := p.GetContent(endpoint_url)

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
	CPUs           int
	CPUUnits       int
	Digest         string
	Disk           int
	Hostname       string
	Net0           string
	IP_Address     string
	Memory         int
	NameServer     string
	OnBoot         int
	OSTemplate     string `json:"ostemplate"`
	QuotaTime      int
	QuotaUGIDLimit int
	SearchDomain   string
	Storage        string
	Swap           int
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
}

func (p Proxmox) GetContainerConfig(req *ExistingContainerRequest) (ContainerConfig, error) {
	endpoint_url := "/api2/json/nodes/" + req.Node + "/lxc/" + req.VMID + "/config"
	body, err := p.GetContent(endpoint_url)

	if err != nil {
		return ContainerConfig{}, err
	}

	var containerConfig ContainerConfigList
	json.Unmarshal(body, &containerConfig)
	return containerConfig.Data, nil
}

type ContainerResponse struct {
	Data string
}

func (p Proxmox) CreateContainer(req *NewContainerRequest) (string, error) {
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

func (p Proxmox) UpdateContainer(req *ExistingContainerRequest) (string, error) {
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

func (p Proxmox) DeleteContainer(req *ExistingContainerRequest) (string, error) {
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
