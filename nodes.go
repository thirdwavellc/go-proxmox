package main

import (
	"encoding/json"
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
