package main

import (
	"encoding/json"
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

func (p Proxmox) GetContainers() []Container {
	endpoint_url := "/api2/json/nodes/" + p.node + "/lxc"
	body := p.GetContent(endpoint_url)
	var containers ContainerList
	json.Unmarshal(body, &containers)
	return containers.Data
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
	IP_Address     string
	Memory         int
	NameServer     string
	OnBoot         int
	OSTemplate     string
	QuotaTime      int
	QuotaUGIDLimit int
	SearchDomain   string
	Storage        string
	Swap           int
}

type ContainerRequest struct {
	Node       string `json:"node"`
	OsTemplate string `json:"ostemplate"`
	VMID       string `json:"vmid"`
}

func (p Proxmox) GetContainerConfig(req ContainerRequest) ContainerConfig {
	endpoint_url := "/api2/json/nodes/" + req.Node + "/lxc/" + req.VMID + "/config"
	body := p.GetContent(endpoint_url)
	var containerConfig ContainerConfigList
	json.Unmarshal(body, &containerConfig)
	return containerConfig.Data
}

type ContainerResponse struct {
	Data string
}

func (p Proxmox) CreateContainer(req *ContainerRequest) string {
	endpoint_url := "/api2/json/nodes/" + req.Node + "/lxc"

	payload := url.Values{}
	payload.Add("ostemplate", req.OsTemplate)
	payload.Add("vmid", req.VMID)

	body := p.PostContent(endpoint_url, payload)
	var response ContainerResponse
	json.Unmarshal(body, &response)
	return response.Data
}

func (p Proxmox) DeleteContainer(req *ContainerRequest) string {
	endpoint_url := "/api2/json/nodes/" + req.Node + "/lxc/" + req.VMID
	payload := url.Values{}

	body := p.DeleteContent(endpoint_url, payload)
	var response ContainerResponse
	json.Unmarshal(body, &response)
	return response.Data
}
