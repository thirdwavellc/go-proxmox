package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
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

func GetContainers(host string, node string, auth AuthInfo) []Container {
	url := host + "/api2/json/nodes/" + node + "/openvz"
	body := GetContent(url, auth)
	var containers ContainerList
	json.Unmarshal(body, &containers)
	return containers.Data
}

func printContainers(containers []Container) {
	for _, container := range containers {
		printContainer(container)
	}
}

func printContainer(container Container) {
	s := reflect.ValueOf(container)
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%s: %v\n", typeOfT.Field(i).Name, f.Interface())
	}
	fmt.Printf("\n")
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

func GetContainerConfig(host string, node string, vmid string, auth AuthInfo) ContainerConfig {
	url := host + "/api2/json/nodes/" + node + "/openvz/" + vmid + "/config"
	body := GetContent(url, auth)
	var containerConfig ContainerConfigList
	json.Unmarshal(body, &containerConfig)
	return containerConfig.Data
}

func printContainerConfig(containerConfig ContainerConfig) {
	s := reflect.ValueOf(containerConfig)
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%s: %v\n", typeOfT.Field(i).Name, f.Interface())
	}
	fmt.Printf("\n")
}

type ContainerRequest struct {
	ContainerConfig ContainerConfig
	Node            string
	VMID            string
}

func CreateContainer(host string, containerRequest *ContainerRequest, auth AuthInfo) []byte {
	requestUrl := host + "/api2/json/nodes/" + containerRequest.Node + "/openvz/" + containerRequest.VMID + "/config"
	fmt.Println("Fetching:", requestUrl)

	data := url.Values{}
	data.Set("node", containerRequest.Node)
	data.Set("vmid", containerRequest.VMID)
	if containerRequest.ContainerConfig.CPUs > 0 {
		data.Set("cpus", strconv.Itoa(containerRequest.ContainerConfig.CPUs))
	}
	if containerRequest.ContainerConfig.Disk > 0 {
		data.Set("disk", strconv.Itoa(containerRequest.ContainerConfig.Disk))
	}
	if containerRequest.ContainerConfig.Hostname != "" {
		data.Set("hostname", containerRequest.ContainerConfig.Hostname)
	}
	if containerRequest.ContainerConfig.IP_Address != "" {
		data.Set("ip_address", containerRequest.ContainerConfig.IP_Address)
	}
	if containerRequest.ContainerConfig.Memory > 0 {
		data.Set("memory", strconv.Itoa(containerRequest.ContainerConfig.Memory))
	}
	if containerRequest.ContainerConfig.Swap > 0 {
		data.Set("swap", strconv.Itoa(containerRequest.ContainerConfig.Swap))
	}

	req, err := http.NewRequest("PUT", requestUrl, bytes.NewBufferString(data.Encode()))
	if err != nil {
		PrintError(err)
	}

	req.Header.Add("CSRFPreventionToken", auth.CSRFPreventionToken)
	cookie := http.Cookie{Name: "PVEAuthCookie", Value: auth.Ticket, Expires: time.Now().Add(356 * 24 * time.Hour), HttpOnly: true}
	req.AddCookie(&cookie)

	fmt.Printf("Request: %+v", req)

	client := &http.Client{}
	resp, err := client.Do(req)
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
