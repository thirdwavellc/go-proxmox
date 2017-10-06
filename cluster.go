package main

import (
	"encoding/json"
)

type Cluster struct {
	Data []ClusterNode
}

type ClusterNode struct {
	Id     string
	Level  string
	Local  int
	Name   string
	NodeId int
	Pmxcfs int
	State  int
	Type   string
}

func (p Proxmox) GetClusterStatus() []ClusterNode {
	endpoint_url := "/api2/json/cluster/status"
	body := p.GetContent(endpoint_url)
	var cluster Cluster
	json.Unmarshal(body, &cluster)
	return cluster.Data
}

type ClusterTaskList struct {
	Data []ClusterTask
}

type ClusterTask struct {
	EndTime   int
	Id        string
	Node      string
	Saved     string
	StartTime string
	Status    string
	Type      string
	UPId      string
	User      string
}

func (p Proxmox) GetClusterTasks() []ClusterTask {
	endpoint_url := "/api2/json/cluster/tasks"
	body := p.GetContent(endpoint_url)
	var clusterTasks ClusterTaskList
	json.Unmarshal(body, &clusterTasks)
	return clusterTasks.Data
}

type ClusterBackupSchedule struct {
	Data []ClusterBackupScheduleItem
}

type ClusterBackupScheduleItem struct {
	All       int
	Compress  string
	DOW       string
	Exclude   string
	Id        string
	MailTo    string
	Mode      string
	Quiet     int
	StartTime string
	Storage   string
}

func (p Proxmox) GetClusterBackupSchedule() []ClusterBackupScheduleItem {
	endpoint_url := "/api2/json/cluster/backup"
	body := p.GetContent(endpoint_url)
	var clusterBackupSchedule ClusterBackupSchedule
	json.Unmarshal(body, &clusterBackupSchedule)
	return clusterBackupSchedule.Data
}
