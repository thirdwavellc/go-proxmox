package proxmox

import (
	"encoding/json"
	"log"
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

func (p ProxmoxClient) GetNodes() ([]Node, error) {
	endpoint_url := "/api2/json/nodes"
	body, err := p.GetContent(endpoint_url)

	if err != nil {
		return nil, err
	}

	var nodes NodeList
	json.Unmarshal(body, &nodes)
	return nodes.Data, nil
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

func (p ProxmoxClient) GetNodeTaskStatus(req NodeTaskStatusRequest) (NodeTaskStatus, error) {
	endpoint_url := "/api2/json/nodes/" + req.Node + "/tasks/" + req.UPID + "/status"
	body, err := p.GetContent(endpoint_url)

	if err != nil {
		return NodeTaskStatus{}, err
	}

	var task NodeTaskStatusList
	json.Unmarshal(body, &task)
	return task.Data, nil
}

func (p ProxmoxClient) CheckNodeTaskStatus(req NodeTaskStatusRequest) (NodeTaskStatus, error) {
	for {
		task, err := p.GetNodeTaskStatus(req)

		if err != nil {
			return NodeTaskStatus{}, err
		}

		if task.Status == "stopped" {
			log.Printf("done.\n")
			return task, nil
		} else {
			log.Printf(".")
			time.Sleep(time.Second)
		}
	}
}
