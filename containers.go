package main

import (
	"encoding/json"
	"fmt"
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
	p.api_endpoint = "/api2/json/nodes/" + p.node + "/openvz"
	body := p.GetContent()
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

func (r ContainerRequest) FormatTemplate() string {
	return fmt.Sprintf("local:vztmpl/%s.tar.gz", r.OsTemplate)
}

func (p Proxmox) GetContainerConfig(req ContainerRequest) ContainerConfig {
	p.api_endpoint = "/api2/json/nodes/" + req.Node + "/openvz/" + req.VMID + "/config"
	body := p.GetContent()
	var containerConfig ContainerConfigList
	json.Unmarshal(body, &containerConfig)
	return containerConfig.Data
}

type ContainerResponse struct {
	Data string
}

func (p Proxmox) CreateContainer(req *ContainerRequest) string {
	p.api_endpoint = "/api2/json/nodes/" + req.Node + "/openvz"

	payload := url.Values{}
	payload.Add("ostemplate", req.FormatTemplate())
	payload.Add("vmid", req.VMID)

	body := p.PostContent(payload)
	var response ContainerResponse
	json.Unmarshal(body, &response)
	return response.Data
}

func (p Proxmox) DeleteContainer(req *ContainerRequest) string {
	p.api_endpoint = "/api2/json/nodes/" + req.Node + "/openvz/" + req.VMID

	payload := url.Values{}

	body := p.DeleteContent(payload)
	var response ContainerResponse
	json.Unmarshal(body, &response)
	return response.Data
}
