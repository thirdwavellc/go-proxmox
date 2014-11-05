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

func GetClusterStatus(host string, auth AuthInfo) []ClusterNode {
	url := host + "/api2/json/cluster/status"
	body := GetContent(url, auth)
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

func GetClusterTasks(host string, auth AuthInfo) []ClusterTask {
	url := host + "/api2/json/cluster/tasks"
	body := GetContent(url, auth)
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

func GetClusterBackupSchedule(host string, auth AuthInfo) []ClusterBackupScheduleItem {
	url := host + "/api2/json/cluster/backup"
	body := GetContent(url, auth)
	var clusterBackupSchedule ClusterBackupSchedule
	json.Unmarshal(body, &clusterBackupSchedule)
	return clusterBackupSchedule.Data
}
