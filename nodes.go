package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type NodeList struct {
	Data []Node
}

type Node struct {
	Disk    int
	CPU     float64
	MaxDisk int
	MaxMem  int
	Node    string
	MaxCPU  int
	Level   string
	Uptime  int
	Id      string
	Type    string
	Mem     int
}

func (p Proxmox) GetNodes() []Node {
	p.api_endpoint = "/api2/json/nodes"
	body := p.GetContent()
	var nodes NodeList
	json.Unmarshal(body, &nodes)
	return nodes.Data
}

type NodeTaskStatusList struct {
	Data NodeTaskStatus
}

type NodeTaskStatus struct {
	ExitStatus string `json:"exitstatus"`
	Id         string `json:"id"`
	Node       string `json:"node"`
	PID        int    `json:"pid"`
	PStart     int    `json:"pstart"`
	StartTime  int    `json:"starttime"`
	Status     string `json:"status"`
	Type       string `json:"type"`
	UPID       string `json:"upid"`
	User       string `json:"user"`
}

type NodeTaskStatusRequest struct {
	Node string `json:"node"`
	UPID string `json:"upid"`
}

func (p Proxmox) GetNodeTaskStatus(req NodeTaskStatusRequest) NodeTaskStatus {
	p.api_endpoint = "/api2/json/nodes/" + req.Node + "/tasks/" + req.UPID + "/status"
	body := p.GetContent()
	var task NodeTaskStatusList
	json.Unmarshal(body, &task)
	return task.Data
}

func (p Proxmox) CheckNodeTaskStatus(req NodeTaskStatusRequest) NodeTaskStatus {
	var task NodeTaskStatus
	for {
		task = p.GetNodeTaskStatus(req)
		if task.Status == "stopped" {
			fmt.Printf("done.\n")
			return task
		} else {
			fmt.Printf(".")
			time.Sleep(time.Second)
		}
	}
}
