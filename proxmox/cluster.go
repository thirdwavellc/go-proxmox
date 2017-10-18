package proxmox

import (
	"encoding/json"
)

type Cluster struct {
	Data []ClusterNode
}

type ClusterNode struct {
	Id     string `json:"id"`
	Ip     string `json:"ip"`
	Level  string `json:"level"`
	Local  int    `json:"local"`
	Name   string `json:"name"`
	NodeId int    `json:"nodeid"`
	Online int    `json:"online"`
	Type   string `json:"type"`
}

func (p ProxmoxClient) GetClusterStatus() ([]ClusterNode, error) {
	endpoint_url := "/api2/json/cluster/status"
	body, err := p.GetContent(endpoint_url)

	if err != nil {
		return nil, err
	}

	var cluster Cluster
	json.Unmarshal(body, &cluster)
	return cluster.Data, nil
}

type ClusterTaskList struct {
	Data []ClusterTask
}

type ClusterTask struct {
	EndTime   int    `json:"endtime"`
	Id        string `json:"id"`
	Node      string `json:"node"`
	Saved     string `json:"saved"`
	StartTime string `json:"starttime"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	UPId      string `json:"upid"`
	User      string `json:"user"`
}

func (p ProxmoxClient) GetClusterTasks() ([]ClusterTask, error) {
	endpoint_url := "/api2/json/cluster/tasks"
	body, err := p.GetContent(endpoint_url)

	if err != nil {
		return nil, err
	}

	var clusterTasks ClusterTaskList
	json.Unmarshal(body, &clusterTasks)
	return clusterTasks.Data, nil
}

type ClusterBackupSchedule struct {
	Data []ClusterBackupScheduleItem
}

type ClusterBackupScheduleItem struct {
	Compress         string `json:"compress"`
	DOW              string `json:"dow"`
	Enabled          string `json:"enabled"`
	Id               string `json:"id"`
	MailNotification string `json:"mailnotification"`
	MailTo           string `json:"mailto"`
	Mode             string `json:"mode"`
	Node             string `json:"node"`
	Quiet            int    `json:"quiet"`
	StartTime        string `json:"starttime"`
	Storage          string `json:"storage"`
	VMID             string `json:"vmid"`
}

func (p ProxmoxClient) GetClusterBackupSchedule() ([]ClusterBackupScheduleItem, error) {
	endpoint_url := "/api2/json/cluster/backup"
	body, err := p.GetContent(endpoint_url)

	if err != nil {
		return nil, err
	}

	var clusterBackupSchedule ClusterBackupSchedule
	json.Unmarshal(body, &clusterBackupSchedule)
	return clusterBackupSchedule.Data, nil
}
