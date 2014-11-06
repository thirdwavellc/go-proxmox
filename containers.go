package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
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
	ContainerConfig ContainerConfig
	node            string
	vmid            string
}

func (p Proxmox) GetContainerConfig(req ContainerRequest) ContainerConfig {
	p.api_endpoint = "/api2/json/nodes/" + req.node + "/openvz/" + req.vmid + "/config"
	body := p.GetContent()
	var containerConfig ContainerConfigList
	json.Unmarshal(body, &containerConfig)
	return containerConfig.Data
}

func (p Proxmox) CreateContainer(req *ContainerRequest) []byte {
	p.api_endpoint = "/api2/json/nodes/" + req.node + "/openvz/" + req.vmid + "/config"
	fmt.Println("Fetching:", p.Url())

	data := url.Values{}
	data.Set("node", req.node)
	data.Set("vmid", req.vmid)
	if req.ContainerConfig.CPUs > 0 {
		data.Set("cpus", strconv.Itoa(req.ContainerConfig.CPUs))
	}
	if req.ContainerConfig.Disk > 0 {
		data.Set("disk", strconv.Itoa(req.ContainerConfig.Disk))
	}
	if req.ContainerConfig.Hostname != "" {
		data.Set("hostname", req.ContainerConfig.Hostname)
	}
	if req.ContainerConfig.IP_Address != "" {
		data.Set("ip_address", req.ContainerConfig.IP_Address)
	}
	if req.ContainerConfig.Memory > 0 {
		data.Set("memory", strconv.Itoa(req.ContainerConfig.Memory))
	}
	if req.ContainerConfig.Swap > 0 {
		data.Set("swap", strconv.Itoa(req.ContainerConfig.Swap))
	}

	request, err := http.NewRequest("PUT", p.Url(), bytes.NewBufferString(data.Encode()))
	if err != nil {
		PrintError(err)
	}

	request.Header.Add("CSRFPreventionToken", p.auth.CSRFPreventionToken)

	cookie := http.Cookie{Name: "PVEAuthCookie", Value: p.auth.Ticket,
		Expires: time.Now().Add(356 * 24 * time.Hour), HttpOnly: true}
	request.AddCookie(&cookie)

	fmt.Printf("Request: %+v", request)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		PrintError(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		PrintError(err)
	}

	fmt.Println(string(body[:]))
	return body
}
