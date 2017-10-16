package proxmox

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
